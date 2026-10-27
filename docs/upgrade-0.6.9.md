# 升级到 0.6.9

0.6.9 恢复 0.4.0 的默认免登录、公开导航，内部安全状态自动准备。新用户在根目录运行 `docker compose up -d flare` 即可使用，无需初始化命令；能访问实例的人均可浏览和修改内容。

## 已有部署

1. 停止实例，备份整个 app 目录（含隐藏文件）、Compose、外部环境变量和 Secret。
2. 保留原项目名、数据挂载和访问配置，只将镜像更新为 `soulteary/flare:0.6.9` 或 `ghcr.io/soulteary/flare:0.6.9`。
3. 已有私有部署继续保留 `FLARE_DISABLE_LOGIN=false`、`FLARE_VISIBILITY=PRIVATE` 及凭据来源。不要用根目录的新免登录示例覆盖旧配置；不是所有旧部署都有隐式私有默认记录。
4. 执行 `docker compose pull flare` 与 `docker compose up -d --force-recreate flare`。
5. 检查 `docker compose exec flare flare --version`，用匿名窗口核验访问模式，再检查登录、编辑和重启后的状态。

已记录的私有默认会继续保留；公开内容不会因为安装目录或样例识别不同而突然要求管理员。登录示例、全站私有示例和 Traefik 示例仍保留各自登录要求。

## 保留的安全机制

- `.flare-security/`（或 FLARE_STATE_DIR）仍须持久化且可写，即使免登录也需要它。
- 免登录与 `private: true` 或显式 PRIVATE 冲突时拒绝开放；请启用登录或核对并修正配置。
- 已提交安全状态缺失、损坏或锁冲突时仍停止业务，不能删除目录来绕过。
- CSRF、来源校验、密码管理和会话撤销继续有效。显式启用登录但缺少凭据时仍需配置账户。

回滚时恢复升级前完整数据及外部配置。0.6.8 新安装采用不同的默认初始化策略，不能把换回镜像当作完整回滚。

应用发布与下载见 [Flare 0.6.9](https://github.com/soulteary/flare/releases/tag/0.6.9)，部署验收见[核验记录](release-0.6.9-verification.md)。
