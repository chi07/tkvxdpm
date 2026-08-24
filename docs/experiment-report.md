# Báo cáo môi trường thí nghiệm: `worklog-service`

## 1. Mục đích

`worklog-service` là REST API Go nhỏ, được tạo để cung cấp cùng một context
repository cho thí nghiệm năng suất AI. Service không tạo ra dữ liệu năng suất;
nó là môi trường để người tham gia thực hiện các nhiệm vụ phát triển có thể kiểm
chứng bằng mã nguồn, commit và unit test.

## 2. Baseline có thể kiểm chứng

- Ngôn ngữ: Go 1.26.
- Dùng SQLite qua driver Go; API và domain vẫn dùng Go standard library.
- Baseline API: `GET /healthz`, `GET/POST /api/v1/worklogs`,
  `GET /api/v1/worklogs/{id}`.
- Kiểm tra baseline: `go test ./...` và `go vet ./...`.
- Dữ liệu seed nằm ở `internal/store/store.go` chỉ là fixture phục vụ test/demo.
  Các số giờ trong fixture không phải số liệu thí nghiệm và không dùng để kết
  luận AI nhanh hay chậm.

## 3. Quy trình thực hiện và bốn nhiệm vụ

Mỗi task được thực hiện trên một nhánh mới từ **cùng baseline commit**. Đồng hồ
bắt đầu khi đọc yêu cầu và dừng khi tất cả acceptance criteria đạt. Lưu video có
timestamp, hash commit, log test trước/sau và — chỉ ở điều kiện AI — prompt có
ảnh hưởng đáng kể. Không đọc yêu cầu task tiếp theo trước khi hoàn thành task
hiện tại. Mọi task đều phải có unit test mới và không được thay seed data để làm
test dễ hơn.

### T1 — AI: Báo cáo số giờ theo dự án

**Yêu cầu:** Thêm `GET /api/v1/reports/project-hours`. API trả về tổng `hours`
theo `project`, sắp xếp project tăng dần theo tên.

**Acceptance criteria:**

1. Response HTTP 200 là JSON array có `project` và `total_hours`.
2. Với seed data, `billing` là `5.5` giờ; `portal` là `7.5` giờ.
3. Có unit test cho route mới.
4. `go test ./...` pass.

**Điều kiện:** Được dùng đúng một AI coding tool đã khai báo; lưu tool/model và
prompt quan trọng.

### T2 — AI: Báo cáo số giờ theo nhãn công việc

**Yêu cầu:** Thêm `GET /api/v1/reports/label-hours`. API trả về tổng `hours`
theo `label`, sắp xếp nhãn tăng dần theo tên.

**Acceptance criteria:**

1. Response HTTP 200 là JSON array có `label` và `total_hours`.
2. Với seed data, `backend` là `4.5`, `frontend` là `4.0`, `testing` là `4.5`.
3. Có unit test cho route mới.
4. `go test ./...` pass.

**Điều kiện:** Được dùng đúng một AI coding tool đã khai báo; lưu tool/model và
prompt quan trọng.

### T3 — NoAI: Báo cáo giờ đã được duyệt theo thành viên

**Yêu cầu:** Thêm `GET /api/v1/reports/approved-member-hours`. API chỉ cộng
worklog có `approved=true`, nhóm theo `member_id`, sắp xếp tăng dần theo ID.

**Acceptance criteria:**

1. Response HTTP 200 là JSON array có `member_id` và `total_hours`.
2. Với seed data, `hai` là `3.0`, `hung` là `1.5`, `trang` là `4.0`; không có `chi`.
3. Có unit test cho route mới.
4. `go test ./...` pass.

**Điều kiện:** Không dùng chatbot, coding agent hoặc autocomplete sinh mã. Chỉ
dùng tài liệu chính thức và tìm kiếm web theo quy ước chung.

### T4 — NoAI: Báo cáo số giờ theo ngày làm việc

**Yêu cầu:** Thêm `GET /api/v1/reports/daily-hours`. API trả về tổng `hours`
theo `work_date` (định dạng `YYYY-MM-DD`), sắp xếp tăng dần theo ngày.

**Acceptance criteria:**

1. Response HTTP 200 là JSON array có `work_date` và `total_hours`.
2. Với seed data, các ngày từ `2026-08-10` đến `2026-08-14` có tổng lần lượt là
   `3.0`, `2.5`, `4.0`, `1.5`, `2.0`.
3. Có unit test cho route mới.
4. `go test ./...` pass.

**Điều kiện:** Không dùng chatbot, coding agent hoặc autocomplete sinh mã. Chỉ
dùng tài liệu chính thức và tìm kiếm web theo quy ước chung.

## 4. Truy vết việc dùng AI

OpenAI Codex được dùng để hỗ trợ khởi tạo baseline: cấu trúc service, mã Go,
unit test baseline và task card. Nhóm phải kiểm tra lại mã, chạy test và chịu
trách nhiệm cuối cùng với toàn bộ nội dung. Khi đo T1–T2, người thực hiện phải
lưu tên công cụ/model và các prompt có ảnh hưởng đáng kể; T3–T4 không được dùng
chatbot, coding agent hoặc autocomplete sinh mã.

## 5. Dữ liệu cần thu thập sau mỗi task

- Mã người tham gia, task ID, điều kiện AI/NoAI và thứ tự thực hiện.
- Thời gian tự ước lượng; thời gian thực tế đầu-cuối.
- Phân đoạn thời gian: hiểu/tra cứu, code/chỉnh sửa, prompt/chờ AI,
  kiểm chứng/sửa lỗi.
- Hash baseline, hash commit kết quả, lệnh và log test, link bằng chứng video.
- Pass/fail theo acceptance criteria và ghi chú gián đoạn.

Kết quả chỉ được báo cáo sau khi các mục trên được thu thập. Không được thay
dữ liệu seed hoặc ví dụ trong task card thành số liệu thực nghiệm.

## 6. Form chạy thử với số liệu minh họa

Bảng dưới đây là kịch bản số liệu dùng để kiểm tra biểu mẫu, công thức và cách
vẽ biểu đồ trước buổi đo chính thức. Khi hoàn tất thí nghiệm, nhóm thay các giá
trị này bằng thời gian ghi từ video, commit và log test.

Các mức thời gian được chọn từ phạm vi hợp lý 35–60 phút cho một task: người làm
phải đọc context service, thêm route, tổng hợp dữ liệu, viết unit test và chạy
`go test ./...`. Cột “ước lượng” là con số người làm dự kiến trước khi bắt đầu;
cột “thực tế” chỉ thay bằng thời gian đo được khi chạy thí nghiệm.

| Task | Điều kiện | Ước lượng trước (phút) | Thực tế minh họa (phút) | Độ lệch ước lượng | Ghi chú minh họa |
|---|---|---:|---:|---:|---|
| T1 | AI | 35 | 45 | -22,2% | Lần đầu đọc route, store và test hiện có. |
| T2 | AI | 32 | 39 | -17,9% | Đã quen kiến trúc sau T1 nhưng vẫn phải review code AI. |
| T3 | NoAI | 40 | 48 | -16,7% | Có điều kiện `approved=true` và nhóm theo thành viên. |
| T4 | NoAI | 42 | 50 | -16,0% | Cần chuẩn hóa/nghiệm thu output theo ngày. |

`Độ lệch ước lượng (%) = 100 × (ước lượng trước − thời gian thực tế) / thời gian thực tế`.

Trong kịch bản chạy thử, thời gian trung bình AI là 42,0 phút, NoAI là 49,0 phút;
chênh lệch AI so với NoAI là -14,3%. Đây là tham số minh họa cho form, không phải
hiệu ứng đã quan sát trong nhóm. Thứ tự AI → AI → NoAI → NoAI còn có thể tạo hiệu ứng
quen service hoặc mệt mỏi; vì vậy khi đo thật cần ghi rõ thứ tự, video/timestamp,
hash commit, log test và mọi gián đoạn.

## 7. Kết luận theo bối cảnh

Với các bài toán đơn giản, phạm vi rõ, nghiệp vụ không quá nhiều và ít phụ thuộc
chéo, AI có thể giúp tạo mới codebase hoặc bổ sung một feature nhanh hơn nhờ khả
năng sinh boilerplate, route, test khung và mã xử lý lặp lại. Trong kịch bản chạy
thử của service này, điều đó được biểu diễn bằng mức chênh lệch 14,3% giữa hai
điều kiện.

Ngược lại, với hệ thống lớn, nhiều lớp nghiệp vụ chồng chéo, quy tắc ngoại lệ,
phụ thuộc liên service và yêu cầu bảo trì cao, chi phí đọc context, kiểm chứng,
đối soát và sửa hồi quy có thể tăng lên. Khi đó, tốc độ sinh code không còn là
đại diện đầy đủ cho năng suất đầu-cuối; AI có thể giúp viết bản nháp nhanh hơn
nhưng chưa chắc rút ngắn thời gian bàn giao sản phẩm.

Do thời lượng demo có hạn, nhóm sử dụng `worklog-service` làm codebase thực hành
thay cho bài toán đang triển khai tại công ty. Đây là lựa chọn cần thiết để không
đưa mã nguồn, dữ liệu nghiệp vụ hoặc thông tin nội bộ vào môi trường thí nghiệm;
vì vậy kết quả chỉ nên được diễn giải như một pilot trong codebase nhỏ, không phải
đánh giá trực tiếp năng suất trên hệ thống sản xuất của doanh nghiệp.
