# Redis：批量工具

## TODO

- [ ] 支持参数化配置 Redis
- [ ] `delete` 二次确认
- [ ] 批量过期时间
- [ ] 支持定时处理
- [ ] 支持可配置化的定时逻辑
- [ ] 打包成二进制命令可供下载

## 初始化

```bash
go mod download
go mod vendor

go build -o redis-batch main.go
# or
make build

./redis-batch XXX
```

## 功能

```bash
go run main.go generate # 生成测试数据
go run main.go delete # 批量删除数据
go run main.go expire # 批量设置过期时间
```

> 待完善……