# TKVX DPM

Repository này chứa mã nguồn và tài liệu phục vụ thí nghiệm về năng suất lập
trình khi dùng AI.

## Cách đọc tài liệu

- [Yêu cầu đề tài](docs/requirements.md)
- [Hướng dẫn chạy worklog-service](docs/service-guide.md)
- [Báo cáo môi trường thí nghiệm](docs/experiment-report.md)
- [Bốn task thí nghiệm AI → AI → NoAI → NoAI](docs/task-cards.md)

## Chạy service

```bash
cd worklog-service
go test ./...
go run ./cmd/worklog-service
```
