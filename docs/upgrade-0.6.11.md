# 升级到 0.6.11

应用发布与下载见 [Flare 0.6.11](https://github.com/soulteary/flare/releases/tag/0.6.11)。本仓库同步 Docker 部署示例，标签为 `v0.6.11`；应用镜像标签不带 `v`。

## 从 0.6.10 升级

1. 停止实例，备份整个 app 目录（含 `.flare-security/` 等隐藏文件）、Compose、外部环境变量和 Secret。
2. 保留原项目名、数据挂载和访问配置，只将镜像更新为 `soulteary/flare:0.6.11` 或 `ghcr.io/soulteary/flare:0.6.11`。
3. 执行 `docker compose pull flare` 与 `docker compose up -d --force-recreate flare`。仅执行 restart 不会替换镜像。
4. 检查 `docker compose exec flare flare --version`，应显示 0.6.11；用匿名窗口核验原有访问模式，再检查设置保存、编辑及重启后的状态。

默认免登录行为没有变化。已有私有部署必须保留 `FLARE_DISABLE_LOGIN=false`、`FLARE_VISIBILITY=PRIVATE` 和凭据来源；不要用根目录示例覆盖原配置，也不要重新初始化已有安全目录。

## 本版变化

写入来源保护失败时，响应不再只有 `SOURCE_REJECTED`，还会提供：

- `error.reason`：稳定原因码，用于区分协议、Host、端口、不透明 Origin、重复来源头或跨站提交；
- `error.message`：可读的失败说明；
- `error.hint`：可执行的代理或访问方式调整建议；
- `request_id`：与服务日志关联的请求标识。

编辑器会展示这些诊断并保留未保存草稿。旧格式错误响应继续使用本地化说明。服务日志记录原因码和请求 ID，但不会复制不可信的 Origin 或可能包含私有路径的 Referer。

这项变化没有放宽来源保护：未配置可信代理时，转发头仍不参与来源判断，跨站写入继续被拒绝。常见 Nginx 部署应保留原始 Host 和非默认端口：

```nginx
proxy_set_header Host $http_host;
proxy_set_header X-Forwarded-Proto $scheme;
```

若 `FLARE_COOKIE_SECURE=true`，请始终从 HTTPS 公共地址访问；故意使用纯 HTTP 的部署应关闭该选项。完整说明见[反向代理与诊断](reverse-proxy.md)。

## 更早版本和回滚

从更早版本升级，先阅读 [0.6.10 升级说明](upgrade-0.6.10.md)、[0.6.9 默认访问行为变化](upgrade-0.6.9.md)及 [0.6.8 安全状态迁移](upgrade-0.6.8.md)。回滚需恢复升级前完整数据及外部配置，不应只替换镜像或删除安全状态目录。

公开产物及本次部署配置检查见[发布核验](release-0.6.11-verification.md)。物理 iPad Safari、Android、Via 未作本版真机验证。
