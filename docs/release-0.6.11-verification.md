# 0.6.11 部署发布核验

日期：2026-10-28（Asia/Shanghai）。应用源码提交为 `85c705bb9271ed9907eeb5744fefacce1040157f`；[应用 Release](https://github.com/soulteary/flare/releases/tag/0.6.11)提供二进制与容器镜像，本仓库使用 `v0.6.11` 标签发布部署示例及文档。

## 应用产物

[正式发布工作流](https://github.com/soulteary/flare/actions/runs/40624785544)的构建、测试、Docker Hub 发布、GHCR 同步和发布报告派发均成功。Release 包含七个平台的二进制归档和 SHA-256 校验和。

Docker Hub 与 GHCR 的 0.6.11 及 latest 使用同一多架构摘要：

`sha256:6739e44562246af77f180fe4fe0fd86ab1d5eee7a2b1b4a4bb277c414664e2b7`

GHCR 发布流程已校验 linux/amd64、linux/arm64、linux/arm/v7、linux/arm/v6 的版本／源码元数据，版本摘要一致，latest 已更新且允许匿名访问。部署仓库不重新构建应用镜像。

## 部署配置与文档

- 四份 Compose 均通过 `docker compose config --format json` 解析，镜像固定为 `soulteary/flare:0.6.11`。
- 与同步前版本比较，Compose 只改变镜像版本；数据挂载、默认免登录或登录／全站私有模式、代理配置及 local 日志的 10m／3 轮转参数保留。
- login/private 示例在缺少 FLARE_PASS 时继续拒绝解析；验证使用独立的合成测试密码。
- 当前 Markdown 的本地链接及固定到应用 0.6.11 的源码文档目标已检查；历史升级和核验记录保持原有版本语义。

[机器可读结果](release-0.6.11-verification.json)记录静态检查断言和应用产物来源。

## 验证边界

本次同步未重复启动四种部署容器，未部署真实 Traefik／TLS 路由，也未验证实体 iPad Safari、Android、Via、树莓派。应用 CI 的 Chromium 自动化与多架构构建不能替代这些运行环境的实测。
