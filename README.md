# TKVX DPM

Repository này chứa mã nguồn và tài liệu phục vụ thí nghiệm về năng suất lập
trình khi dùng AI.

## Cách đọc tài liệu

- [Yêu cầu đề tài](docs/requirements.md)
- [Hướng dẫn chạy worklog-service](docs/service-guide.md)
- [Hướng dẫn thí nghiệm: baseline, task và form đo](docs/experiment-report.md)

## Chạy service

```bash
cd worklog-service
go test ./...
go run ./cmd/worklog-service
```
