# 0.6.8 部署核验

日期：2026-10-27（Asia/Shanghai）。应用源码提交为 `36d909159d2422a46934f5bb1396f14bccaa4f17`，应用镜像和二进制由 [Flare 源码仓库](https://github.com/soulteary/flare/releases/tag/0.6.8)发布；本仓库只发布部署示例及文档版本 `v0.6.8`。

## 应用发布与发布后检查

- [最终提交预检](https://github.com/soulteary/flare/actions/runs/40465694840)和[正式发布](https://github.com/soulteary/flare/actions/runs/40466217420)全部成功，覆盖七平台归档、四平台 Docker/QEMU 容器、归档二进制浏览器测试及双仓库发布。
- 发布后文档 PR 的 [CI](https://github.com/soulteary/flare/actions/runs/40466966038)、[容器预检](https://github.com/soulteary/flare/actions/runs/40466966062)、[安全扫描](https://github.com/soulteary/flare/actions/runs/40466966017)、[浏览器验收](https://github.com/soulteary/flare/actions/runs/40466965992)、[CodeQL](https://github.com/soulteary/flare/actions/runs/40466966020)均成功。
- 文档合入 main 后的 [CI](https://github.com/soulteary/flare/actions/runs/40467053084)、[浏览器验收](https://github.com/soulteary/flare/actions/runs/40467053134)、[CodeQL](https://github.com/soulteary/flare/actions/runs/40467053062)均成功；主仓库本地 Markdown 链接检查为零错误。

Docker Hub 和 GHCR 的 `0.6.8`、`latest` 已经独立匿名读取核验，四个平台集合一致，共享多平台摘要：

`sha256:4e01c92941b42cd28b8e5aa6b8db9f44af03a9ba03edee226a6fc6b05d80c24b`

## 本仓库 Compose 验收

四份实际 Compose 均通过 `docker compose config`，镜像固定为 0.6.8；确认 `/app` 整目录挂载以及 `local`、`10m`、`3` 日志轮转参数。login/private 示例在未提供 FLARE_PASS 时明确拒绝解析，不使用示例弱密码。

容器行为采用独立临时数据、合成书签和测试密码验证；没有访问或修改用户实际 app 数据。根目录和 Traefik 配置通过 `security init` 初始化，login/private 使用部署凭据。四组全部通过：

- 根目录、private 和 Traefik 后端在匿名访问时要求登录。
- login 示例公开首页，公开条目可见，private 条目对匿名不可见。
- 四组均能登录并读取私有条目；Traefik HTTPS origin 配置的会话 Cookie 含 Secure。
- 强制重建容器后账号、实例身份和原会话保留。
- 三个 YAML 字节保持不变；安全状态中没有测试密码明文；二进制报告 `0.6.8`。

机器可读结果见[核验 JSON](release-0.6.8-verification.json)。

本仓库全部 12 份 Markdown 文档的本地目标和锚点检查通过（78 个目标）；变更记录中 5 个历史失效目录锚点已修复。`git diff --check` 通过。

## 实际测试范围

本机 Docker 接口为 Podman 5.8.3，Compose 5.1.4，运行原生 Linux ARM64 容器。该引擎不支持 Docker 的 `local` 日志驱动，所以临时运行配置仅将日志驱动改为 `k8s-file`；正式 Compose 保留 `local`，轮转参数通过静态解析核验，不声称已在本机执行 Docker local 轮转。

临时配置使用 loopback 随机端口、独立网络、临时 `/app` 和离线天气设置。Traefik 组检查的是后端 origin、Host 和 Cookie 行为，测试客户端手动携带 Secure Cookie；没有启动真实 Traefik/TLS 代理，不能据此声称完成浏览器 HTTPS 或真实路由验收。其他平台的应用容器验证来自源码仓库上述四平台预检，不能替代本部署的实体树莓派或手机测试。Safari/Via/Android/iPad 真机和网络盘、多副本、完整子路径均不作未验证承诺。
