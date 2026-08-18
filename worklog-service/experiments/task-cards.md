# Bốn task thí nghiệm cho `worklog-service`

## Quy tắc chung

- Tạo một nhánh và một commit riêng cho từng task, từ cùng baseline commit đã ghi lại.
- Đồng hồ bắt đầu khi đọc task card và dừng khi toàn bộ acceptance criteria đạt.
- Quay màn hình có timestamp; lưu hash commit, lệnh test, log test trước/sau và prompt đáng kể (chỉ ở điều kiện AI).
- Không dùng AI dưới mọi hình thức trong NoAI: chatbot, IDE completion sinh mã hoặc coding agent.
- Tất cả task đều yêu cầu thêm unit test cho hành vi mới. Không sửa seed data theo hướng làm test dễ hơn.
- Không được đọc task card tiếp theo trước khi hoàn thành task hiện tại.

## T1 — AI: Báo cáo số giờ theo dự án

**Yêu cầu:** Thêm `GET /api/v1/reports/project-hours`. API trả về tổng `hours` theo `project`, với các project sắp xếp tăng dần theo tên.

**Acceptance criteria:**

1. Response HTTP 200 là JSON array có các trường `project` và `total_hours`.
2. Với seed data, `billing` có `5.5` giờ; `portal` có `7.5` giờ.
3. Có unit test cho route mới.
4. `go test ./...` pass.

**Điều kiện:** Được dùng đúng một AI coding tool đã khai báo. Lưu prompt và model/tool.

## T2 — AI: Báo cáo số giờ theo nhãn công việc

**Yêu cầu:** Thêm `GET /api/v1/reports/label-hours`. API trả về tổng `hours` theo `label`, sắp xếp tăng dần theo tên nhãn.

**Acceptance criteria:**

1. Response HTTP 200 là JSON array có `label` và `total_hours`.
2. Với seed data, `backend` là `4.5`, `frontend` là `4.0`, `testing` là `4.5`.
3. Có unit test cho route mới.
4. `go test ./...` pass.

**Điều kiện:** Được dùng đúng một AI coding tool đã khai báo. Lưu prompt và model/tool.

## T3 — NoAI: Báo cáo giờ đã được duyệt theo thành viên

**Yêu cầu:** Thêm `GET /api/v1/reports/approved-member-hours`. API chỉ cộng các worklog có `approved=true`, nhóm theo `member_id`, sắp xếp tăng dần theo member ID.

**Acceptance criteria:**

1. Response HTTP 200 là JSON array có `member_id` và `total_hours`.
2. Với seed data, `hai` là `3.0`, `hung` là `1.5`, `trang` là `4.0`; không có `chi`.
3. Có unit test cho route mới.
4. `go test ./...` pass.

**Điều kiện:** Không dùng AI coding tool, chatbot hoặc completion sinh mã. Chỉ dùng tài liệu chính thức và tìm kiếm web theo quy ước của nhóm.

## T4 — NoAI: Báo cáo giờ theo ngày làm việc

**Yêu cầu:** Thêm `GET /api/v1/reports/daily-hours`. API trả về tổng `hours` theo `work_date`, định dạng ngày `YYYY-MM-DD`, sắp xếp tăng dần theo ngày.

**Acceptance criteria:**

1. Response HTTP 200 là JSON array có `work_date` và `total_hours`.
2. Với seed data, mỗi ngày từ `2026-08-10` đến `2026-08-14` có đúng một phần tử; tổng tương ứng lần lượt là `3.0`, `2.5`, `4.0`, `1.5`, `2.0`.
3. Có unit test cho route mới.
4. `go test ./...` pass.

**Điều kiện:** Không dùng AI coding tool, chatbot hoặc completion sinh mã. Chỉ dùng tài liệu chính thức và tìm kiếm web theo quy ước của nhóm.
