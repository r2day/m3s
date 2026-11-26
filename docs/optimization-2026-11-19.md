# 2026-11-19 扫描：问题与优化步骤

只读扫描本仓库（不含 `vendor`）。本文不记录密钥明文。总览见 [r2day/docs](https://github.com/r2day/docs/blob/main/architecture/2026-11-19-api-optimization.md)。

公共模型库。finance-api、pay-api、fileserver、device-api、client-api 等 require 本模块。本次按 `origin/main` 编写，不包含工作区里未提交的 `client/home/model.go`。

## 问题

### P1 支付密钥结构体会被 JSON 序列化

`finance/keys/model.go` 的 `Model.Private`、`Merchant.APIKey`、`Merchant.AppSecret`、`LakalaConf.CredentialCode` 都有 json 标签。finance-api 的 `POST /rpc/pay/keys` 在内部令牌通过后返回的就是这份文档（见 finance-api 的优化文档）。任何把 `Model` 直接写入响应的处理函数都会带出私钥、微信 API 密钥、小程序 AppSecret 和拉卡拉凭证。

`Merchant.Secret` 已是 `bson:"-"`，只作入参别名。读路径没有同样的处理。

## 步骤

1. 读响应使用不含密钥的结构，或给 `Private`、`APIKey`、`AppSecret`、`CredentialCode` 加上 `json:"-"`。写入仍接受单独的请求结构。
2. finance-api 的密钥接口改为返回「是否已配置」，不返回密钥正文。已发出的响应视为泄露，轮换对应商户的微信和拉卡拉密钥。
3. 不要把本模型当日志字段打印。
