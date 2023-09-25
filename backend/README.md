# 后端

## 1. 创建项目

```bash
kratos new verifyCode
kratos new verifyCode -r https://gitee.com/go-kratos/kratos-layout.git
go mod tidy
```

## 2. 依赖注入工具安装

```bash
go get github.com/google/wire/cmd/wire
go generate ./...
```

## 3. 运行项目

```bash
kratos run
```

## 4. 编写 proto

```bash
# 编写 proto
kratos proto add api/verifyCode/verifyCode.go

# 生成go代码
kratos proto client api/verifyCode/verifyCode.proto

# 生成业务逻辑相关代码
kratos proto server api/verifyCode/verifyCode.proto -t internal/service

# 修改service.go
var ProviderSet = wire.NewSet(NewGreeterService, NewVerifyCodeService)

# 增加服务参数及注册 internal/server/grpc.go
# verifyCodeService *service.VerifyCodeService
verifyCode.RegisterVerifyCodeServer(srv, verifyCodeService)

# 再次生成
go generate ./...
```

```go
package main
// 1. 全局设置随机数种子
// rand.NewSource(time.Now().UnixNano())

// 2. 一次随机会获取 63 位随机数, 浪费

// 3. 借用位运算优化，重复利用一次随机数
```
