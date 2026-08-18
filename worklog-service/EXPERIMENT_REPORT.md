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
[`experiments/task-cards.md`](experiments/task-cards.md).

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
