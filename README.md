# Go-OA-Proxy

```bash
go build -ldflags "-w -s" -tags=jsoniter .

// 交叉编译
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s"
```