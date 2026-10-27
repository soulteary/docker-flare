# 0.6.9 部署核验

日期：2026-10-27（Asia/Shanghai）。应用源码提交为 `9e8b34ba1bfa7f2ec0ca153d82f9fa5ce14897ea`；应用镜像和二进制由[上游 0.6.9](https://github.com/soulteary/flare/releases/tag/0.6.9)发布，本仓库发布部署示例与文档标签 `v0.6.9`。

## 上游发布

[正式发布](https://github.com/soulteary/flare/actions/runs/40473301985)全部成功，包含七平台归档、Docker Hub 和 GHCR。七份下载归档的 SHA-256、内嵌版本及提交号已经核对。正式 macOS arm64 二进制的空目录启动和重启验证通过。

[合并提交 CI](https://github.com/soulteary/flare/actions/runs/40472952039)、[浏览器验收](https://github.com/soulteary/flare/actions/runs/40472952074)、[CodeQL](https://github.com/soulteary/flare/actions/runs/40472952175)、[安全扫描](https://github.com/soulteary/flare/actions/runs/40472952208)全部成功。

Docker Hub 与 GHCR 的 `0.6.9`、`latest` 使用同一四平台摘要，GHCR 发布流程验证了平台元数据和匿名访问：

`sha256:dcada7886b44898daf29e9ab35ed4f72bd92ec49e225cadcec0fed377b88bad9`

## Compose 与文档检查

四份 Compose 通过解析：镜像固定为 0.6.9，保留 `/app` 数据挂载和 `local` 日志驱动的 `10m`、`3` 轮转参数。根目录不再强制开启登录；login/private/Traefik 保留登录保护，login/private 缺少密码时仍拒绝解析。

全部本地 Markdown 链接目标和当前 0.6.9 上游源码文档目标已检查。历史 0.6.8 核验记录保留原始版本语义。

## 隔离容器行为

四组 Linux ARM64 容器均通过：

- 根目录：没有初始化管理员即可匿名访问首页；重建后仍可用，实例身份保持。
- login：匿名可见公开条目，私有条目隐藏；登录后可见私有内容。
- private 与 Traefik 后端：匿名访问要求登录，管理员可读取私有内容。
- 三组登录部署重建后保留有效会话；四组 YAML 字节保持不变，安全状态没有测试密码明文。
- Traefik 后端按配置的 HTTPS origin 签发 Secure Cookie。
- 容器报告版本均为 `0.6.9`。

[机器可读核验结果](release-0.6.9-verification.json)记录具体断言。

## 验收范围

本机 Docker 接口为 Podman 5.8.3，临时容器使用独立数据、loopback 随机端口和测试密码，不访问用户实际 app 数据。生产 Compose 保留 `local` 日志驱动，临时运行配置使用 Podman 的 `k8s-file`，不声称在本机执行 Docker local 日志轮转。

Traefik 验收只覆盖后端 origin/Host/Cookie，不包括真实 Traefik/TLS 路由。实体 Safari、Via、Android、iPad、树莓派，以及网络盘、多副本和完整子路径未作额外验证。
