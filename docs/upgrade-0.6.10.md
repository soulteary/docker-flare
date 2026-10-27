# 升级到 0.6.10

应用发布与下载见 [Flare 0.6.10](https://github.com/soulteary/flare/releases/tag/0.6.10)。本仓库同步 Docker 部署示例，标签为 `v0.6.10`；应用镜像标签不带 `v`。

## 从 0.6.9 升级

1. 停止实例，备份整个 app 目录（含 `.flare-security/` 等隐藏文件）、Compose、外部环境变量和 Secret。
2. 保留原项目名、数据挂载和访问配置，只将镜像更新为 `soulteary/flare:0.6.10` 或 `ghcr.io/soulteary/flare:0.6.10`。
3. 执行 `docker compose pull flare` 与 `docker compose up -d --force-recreate flare`。仅执行 restart 不会替换镜像。
4. 检查 `docker compose exec flare flare --version`，应显示 0.6.10；用匿名窗口核验原有访问模式，再检查设置保存、编辑及重启后的状态。

0.6.9 恢复的默认免登录行为继续保留，内部安全状态仍需持久化且可写。已有私有部署必须保留 `FLARE_DISABLE_LOGIN=false`、`FLARE_VISIBILITY=PRIVATE` 和凭据来源；不要用根目录示例覆盖原配置，也不要重新初始化已有安全目录。

## 本版变化

设置页显式声明 same-origin Referrer 策略，修复反向代理将响应策略改成 no-referrer 后，原生表单 Origin 变为 null 导致的保存拒绝。若仍遇到 SOURCE_REJECTED，请核对 Host、外部协议／端口和可信代理配置，参见[反代说明](reverse-proxy.md)。本版不允许跨来源写入。

在「设置 → 外观 → 手机和平板布局」选择紧凑模式：最多覆盖 1024 CSS 像素，缩小导航间距，应用描述单行省略，工具按钮置于内容末尾。标题仍可换行；默认布局和字体／列数选项保留。紧凑书签行最小高度为 28px，默认手机布局保留 44px 点击区域。

新增隐藏版本信息选项，并优化导航读取与不必要的 CSRF 处理；需要保护的写入仍保留原有校验。

## 更早版本和回滚

从更早版本升级，先阅读 [0.6.9 默认访问行为变化](upgrade-0.6.9.md)及 [0.6.8 安全状态迁移](upgrade-0.6.8.md)。回滚需恢复升级前完整数据及外部配置，不应只替换镜像或删除安全状态目录。

公开产物及本次部署配置检查见[发布核验](release-0.6.10-verification.md)。物理 iPad Safari、Android、Via 未作本版真机验证。
