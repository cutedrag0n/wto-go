# wto-go

基于 Resty v2 的 WTO 非官方 Go SDK，要求 Go 1.24 或更高版本。覆盖官方目录的 27 个操作：[接口覆盖清单](docs/api-coverage.md)。完整选项和示例见 [英文 README](README.md)。

## 安装与认证

```sh
go get github.com/cutedrag0n/wto-go
export WTO_API_KEY='你的订阅密钥'
```

到 [WTO 开发者门户](https://apiportal.wto.org/signup)注册并订阅免费产品，获取密钥。SDK 使用 `Ocp-Apim-Subscription-Key` Header 认证，环境变量由应用读取。

```go
client, err := wto.NewClient(os.Getenv("WTO_API_KEY"))
if err != nil { /* 处理错误 */ }
limit := 100
rows, err := client.Timeseries.Data(ctx, timeseries.DataParams{
    Indicator: "ITS_MTV_AX",
    Reporters: "840",
    Periods: "2020-2023",
    Max: &limit,
})
```

导入根包 `github.com/cutedrag0n/wto-go`（包名 `wto`），参数和响应类型分别放在 `timeseries`、`eping`、`qr`、`tfad` 子包。所有方法接收 Context。

- `client.Timeseries`：数据、计数、元数据、指标及代码表。`DataPost` 使用 JSON 请求体，适合较长查询。
- `client.EPing`：成员及 SPS/TBT 通知搜索。
- `client.QR`：成员、产品、通知、限制措施及 HS 版本；分页结果保留 `Data/Links/Meta`。
- `client.TFAD`：程序、联系信息及单一窗口查询。

数据查询使用机器可读列名，从实际 `Dataset` 包装提取数据；支持 full 和 codes 模式。分页显式设置 `Offset/Max`。`DataRaw/DataPostRaw` 保留人类可读列名、元数据和原始输出；`Download/DownloadPost` 将 CSV 的 ZIP 输出写入 `io.Writer`。下载先缓冲响应，大数据查询应限制范围。

## 限流与配置

默认启用 `golang.org/x/time/rate` 开源令牌桶，并组合小时计数：全部请求平滑限制为 10 次/秒（burst=1）；滚动一小时最多 10,000 次；Timeseries data 与 data_count 合计 1 次/秒（包含 POST）。首次请求和每次重试均计数。

默认总调用超时 30 秒，包含限流等待、网络执行和重试退避。等待可被 Context 取消。小时额度用尽后等待最早记录过期，也可能先达到调用超时。

```go
limiter := wto.NewRateLimiter()
client, err := wto.NewClient(key,
    wto.WithRateLimiter(limiter),
    wto.WithTimeout(2*time.Minute),
    wto.WithRetry(2),
)
```

同一订阅的多个 Client 应显式复用 `limiter`。限流状态仅在进程内存中保存，无法协调其他进程、外部应用或进程重启前的调用。官方未说明小时重置边界，因此使用保守的滚动窗口。

默认重试次数为零；开启后重试临时网络故障和 429/502/503/504。支持 `Retry-After` 秒数、HTTP 日期，429 冷却由共享限流器统一执行。

其他配置包括 `WithBaseURL`、`WithServiceBaseURL`、`WithHTTPClient`、`WithProxy`、`WithHeader`。服务 URL 覆盖需包含完整服务前缀；自定义 HTTP Client 配置复制后使用，Transport/Jar 共享；代理克隆标准 Transport，不能应用到自定义 RoundTripper 时返回错误。跨源重定向停止以保护密钥。

## 错误与模型缺口

使用 `errors.As(err, &apiErr)` 获取 `*wto.APIError`，检查状态码、官方错误字段和原始 `Body`。使用 `errors.Is` 判断 Context 取消、截止时间。错误字符串不输出服务端正文；正文可通过错误对象读取。

官方缺少模型的 ePing Members 和 TFAD 返回 `json.RawMessage`，未知嵌套字段同样保留原始 JSON。QR schema 与示例的冲突及真实联调差异见 [API 说明](docs/api-notes.md)。

## 测试

```sh
go test ./...
go test -race -cover ./...
go vet ./...
go build ./...
WTO_API_KEY='你的密钥' go test -tags=integration -v ./integration
```

单元测试完全离线；真实联调通过 build tag 隔离，未提供密钥时跳过，执行时消耗配额。不要将真实密钥写入代码或提交到仓库。

详细测试范围与联调结果见 [验证报告](docs/testing.md)。
