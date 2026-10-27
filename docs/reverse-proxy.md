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

Traefik 示例设置了 HTTPS origin 和 Secure Cookie，必须同步修改路由 Host 与外部 URL。HTTPS Cookie 不会在代理配置故障时自动降级；发生来源、CSRF 或认证错误时修正源配置，不要关闭保护。新实例默认私有，需先初始化管理员；已有公开首页部署还需用匿名窗口检查 private 分类／条目不可见。

完整字段与边界见[0.6.8 反向代理指南](https://github.com/soulteary/flare/blob/0.6.8/docs/reverse-proxy.md)。
