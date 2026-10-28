# 反向代理与子路径边界

推荐独立域名根路径，例如 `https://flare.example.com/`。目前没有完整的 `/flare/`
子路径支持；资源、表单、路由和 Cookie 包含根路径引用，单独重写首页不能解决全部问题。

TLS 在代理终止、上游为 HTTP 时，设置 `FLARE_COOKIE_SECURE=true`，保持通过 HTTPS
访问，并保留原始 Host（包括非默认公共端口）。不要向公网暴露后端，也不要把未经
信任的 Forwarded 头当作 Cookie 或动态 URL 的通用协议来源。

在已有 Nginx TLS server 中配置：

```nginx
location / {
    proxy_pass http://flare:5005;
    proxy_set_header Host $http_host;
    proxy_set_header X-Forwarded-Proto $scheme;
    client_max_body_size 10m;
}
```

Nginx Proxy Manager 使用独立域名、HTTP 上游 flare:5005，不配置 `/flare/` 自定义位置。
Traefik 可参考[现有示例](../example/traefik/docker-compose.yml)，按实际网络和证书解析器调整。
代理不要缓存私有 HTML、设置或编辑文档；允许 8 MiB 导入所需的请求体和超时。

用新浏览器验证未登录访问、HTTPS 登录、Secure Cookie、编辑保存及退出。0.6.5 搜索
设置若提示请求拒绝，检查 Host/端口和表单是否在服务重启前打开，不要先关闭安全保护。
动态地址占位符使用应用所见的请求信息；后端 HTTP 与外部 HTTPS 是不同层级。

## 0.6.8 的 origin、Host 和受信代理

普通保留原 Host 的根路径反代无需填写全部高级选项。需指定外部 origin 时设置 `FLARE_EXTERNAL_URL=https://flare.example.com`；`FLARE_ALLOWED_HOSTS` 是逗号分隔的主机／端口列表，`FLARE_TRUSTED_PROXIES` 是受信代理 IP/CIDR。只有受信代理提供的来源头会参与判断，勿信任任意公网来源。

Traefik 示例设置了 HTTPS origin 和 Secure Cookie，必须同步修改路由 Host 与外部 URL。HTTPS Cookie 不会在代理配置故障时自动降级；发生来源、CSRF 或认证错误时修正源配置，不要关闭保护。Traefik 示例显式配置全站私有，需先初始化管理员；0.6.9 默认示例免登录，不能据此假定反向代理后自动受保护；已有公开首页部署还需用匿名窗口检查 private 分类／条目不可见。

完整字段与边界见[0.6.11 反向代理指南](https://github.com/soulteary/flare/blob/0.6.11/docs/reverse-proxy.md)。

## 0.6.10 设置保存修复

设置页通过文档中的 same-origin Referrer 策略，抵御代理响应头覆盖为 no-referrer 导致的原生表单 Origin:null。该修复不解决 Host 被改写、外部协议或端口不一致；来源校验和 CSRF 保护继续生效。升级后刷新设置页面再提交，仍被拒绝时按上面的代理配置排查。

## 0.6.11 可操作诊断

来源拒绝响应现在包含稳定的 `error.reason`、可读的 `error.message`、调整建议 `error.hint` 和 `request_id`。常见原因包括 `ORIGIN_SCHEME_MISMATCH`、`ORIGIN_HOST_MISMATCH`、`ORIGIN_PORT_MISMATCH`、`ORIGIN_OPAQUE` 与 `FETCH_SITE_NOT_SAME_ORIGIN`。编辑器会展示这些信息并保留草稿；服务日志记录相同原因码和请求 ID。

反馈问题时附上原因码、请求 ID、公开访问 URL 结构和经过脱敏的代理配置。不要发布 Cookie、密码、Secret 或完整私有 Referer。诊断增强没有放宽来源保护，也不会自动信任未配置的代理转发头。
