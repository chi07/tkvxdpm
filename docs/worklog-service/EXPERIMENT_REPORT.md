# Báo cáo môi trường thí nghiệm: `worklog-service`

## 1. Mục đích

`worklog-service` là REST API Go nhỏ, được tạo để cung cấp cùng một context
repository cho thí nghiệm năng suất AI. Service không tạo ra dữ liệu năng suất;
nó là môi trường để người tham gia thực hiện các nhiệm vụ phát triển có thể kiểm
chứng bằng mã nguồn, commit và unit test.

## 2. Baseline có thể kiểm chứng

- Ngôn ngữ: Go 1.26.
- Không dùng thư viện ngoài; chạy bằng Go standard library.
- Baseline API: `GET /healthz`, `GET/POST /api/v1/worklogs`,
  `GET /api/v1/worklogs/{id}`.
- Kiểm tra baseline: `go test ./...` và `go vet ./...`.
- Dữ liệu seed nằm ở `internal/store/store.go` chỉ là fixture phục vụ test.
  Các số giờ trong fixture không phải số liệu thí nghiệm và không dùng để kết
  luận AI nhanh hay chậm.

## 3. Bốn nhiệm vụ thí nghiệm

Mô tả chi tiết và acceptance criteria nằm trong
[`task-cards.md`](task-cards.md).

| Thứ tự | Điều kiện | Nhiệm vụ |
|---|---|---|
| T1 | AI | Thêm báo cáo tổng số giờ theo project. |
| T2 | AI | Thêm báo cáo tổng số giờ theo nhãn công việc. |
| T3 | NoAI | Thêm báo cáo số giờ đã duyệt theo thành viên. |
| T4 | NoAI | Thêm báo cáo số giờ theo ngày làm việc. |

Mỗi nhiệm vụ phải thực hiện trên nhánh mới từ cùng baseline commit, có unit test
mới, commit riêng, log `go test ./...`, timestamp/video và nhật ký thời gian.
Không đọc task card tiếp theo trước khi hoàn thành task đang làm.

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

> **Cảnh báo:** Bảng dưới đây là số liệu minh họa để kiểm tra biểu mẫu, công
> thức và cách vẽ biểu đồ. Đây không phải dữ liệu do thành viên thực hiện; không
> được dùng để kết luận AI làm nhanh/chậm hơn trong bài luận hoặc slide.

Các mức thời gian được chọn từ phạm vi hợp lý 35–60 phút cho một task: người làm
phải đọc context service, thêm route, tổng hợp dữ liệu, viết unit test và chạy
`go test ./...`. Cột “ước lượng” là con số người làm dự kiến trước khi bắt đầu;
cột “thực tế” chỉ thay bằng thời gian đo được khi chạy thí nghiệm.

| Task | Điều kiện | Ước lượng trước (phút) | Thực tế minh họa (phút) | Độ lệch ước lượng | Ghi chú minh họa |
|---|---|---:|---:|---:|---|
| T1 | AI | 40 | 52 | -23,1% | Lần đầu đọc route, store và test hiện có. |
| T2 | AI | 35 | 43 | -18,6% | Đã quen kiến trúc sau T1 nhưng vẫn phải review code AI. |
| T3 | NoAI | 40 | 48 | -16,7% | Có điều kiện `approved=true` và nhóm theo thành viên. |
| T4 | NoAI | 40 | 50 | -20,0% | Cần chuẩn hóa/nghiệm thu output theo ngày. |

`Độ lệch ước lượng (%) = 100 × (ước lượng trước − thời gian thực tế) / thời gian thực tế`.

Trong bảng minh họa, thời gian trung bình AI là 47,5 phút, NoAI là 49,0 phút;
chênh lệch AI so với NoAI là -3,1%. Đây chỉ là một phép tính kiểm tra form, không
phải hiệu ứng quan sát được. Thứ tự AI → AI → NoAI → NoAI còn có thể tạo hiệu ứng
quen service hoặc mệt mỏi; vì vậy khi đo thật cần ghi rõ thứ tự, video/timestamp,
hash commit, log test và mọi gián đoạn.
