# Hướng dẫn chạy worklog-service

`worklog-service` là REST API Go nhỏ quản lý nhật ký công việc của thành viên dự án. Service dùng thư viện chuẩn của Go, có dữ liệu seed, route HTTP, validation và unit test. Nó được tạo làm bối cảnh chung cho thí nghiệm năng suất AI: nhiệm vụ yêu cầu đọc cấu trúc service, domain model, store và handler thay vì giải thuật có sẵn trên mạng.

## Chạy service

```bash
go test ./...
go run ./cmd/worklog-service
curl http://localhost:8080/healthz
curl 'http://localhost:8080/api/v1/worklogs?member_id=hai'
```

## API baseline

- `GET /healthz`
- `GET /api/v1/worklogs?member_id={id}`
- `GET /api/v1/worklogs/{id}`
- `POST /api/v1/worklogs`

Chi tiết nhiệm vụ, bằng chứng và tiêu chí chấp nhận nằm ở
[`experiment-report.md`](experiment-report.md).
