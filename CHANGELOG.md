# 变更记录

项目应用中的所有变更记录，将在这个文件中进行存档。

## 目录

* [0.3.3](#033---2022-03-10)
* [0.3.2](#032---2022-02-23)
* [0.3.1](#031---2022-02-23)
* [0.2.10](#027---2022-02-21)
* [0.2.9](#027---2022-02-20)
* [0.2.8](#027---2022-02-18)
* [0.2.7](#027---2022-02-17)
* [0.2.6](#026---2022-02-17)
* [0.2.5](#025---2022-02-16)
* [0.2.4](#024---2022-02-11)
* [0.2.3](#023---2022-02-08)
* [0.2.1](#021---2022-02-07)
* [0.2.0](#020---2022-01-30)
* [0.1.0](#010---2022-01-27)

# 未发布 Unreleased

- [新增] 新增 README.md 英文翻译

# 0.3.3 - 2022-03-10

- [新增] 支持从 envfile 中设置应用启动选项。
- [修正] 现在支持完全从 env 中设置应用启动配置，实现诸如禁用离线模式，设置用户账号密码等。
- [修正] DockerHub 中镜像版本不是最新内容。https://github.com/soulteary/docker-flare/issues/43  （感谢 @flashliao @shuax @ZhXWei 反馈 ）
- [修正] 禁用界面应用选项后，帮助界面不展示内容。 https://github.com/soulteary/docker-flare/issues/39 （感谢 @harrison-guo-chn 反馈）
- [优化] 禁用 Chrome 默认弹出烦人的翻译对话框。 https://github.com/soulteary/docker-flare/issues/41 （感谢 @shuax 反馈）
- [优化] 新增界面选项，不再强制界面中的英文展示为大写。 https://github.com/soulteary/docker-flare/issues/31 （感谢 @llussy @kscu 反馈）
- [优化] 拆分和重构程序构建脚本，为自定义主题做准备。

# 0.3.2 - 2022-02-23

- [修正] 解决不能正确设置主题的问题，参考 https://github.com/soulteary/docker-flare/issues/26 感谢 @kx 反馈。

# 0.3.1 - 2022-02-23

当前版本新增了四个功能，并进行了大规模的重构，整理和开放了一部分代码。

- [新增] 图标支持使用外链、支持针对未填写图标名称或外链图标地址的项目，根据连接主机名称进行 favicon 获取和展示。 https://github.com/soulteary/docker-flare/issues/14 （感谢 @lmm214 @eallion 反馈）
- [新增] 将首页内容设置为登陆后可见。 https://github.com/soulteary/docker-flare/issues/23 （ 感谢 @OTritium 反馈）
- [新增] 将首页左下角的工具按钮（设置、帮助）在界面中进行隐藏。 https://github.com/soulteary/docker-flare/issues/20 （感谢 @fgprodigal 反馈）
- [新增] 支持将页面展示的链接进行简单的编码处理，避免直接被搜索引擎或机器人记录。
- [优化] 重构和持续迭代项目，整理了 10% 左右的代码到 https://github.com/soulteary/flare

# 0.2.10 - 2022-02-21

当前版本主要修正了编辑器的使用问题，以及更新了帮助页面展示的功能链接。

- [修正] 修正编辑器不能使用的问题。 https://github.com/soulteary/docker-flame/issues/2 （感谢 @Mantyke 反馈）
- [优化] 完善帮助页面的功能链接。

# 0.2.9 - 2022-02-20

当前版本新增了在线编辑功能，以及一个用于聚合内部工具地址的帮助页面，完善了程序启动时和环境变量以及命令行参数的交互，并进行了大量重构工作。

- [新增] 新增在线编辑器功能，解决一些小伙伴提到的“随时随地”编辑的需求。 https://github.com/soulteary/docker-flare/issues/11
- [新增] 考虑到接下来随着程序功能完善，内部工具链接会越来越多，新增一个“帮助页面”，将程序内置链接集中存放和展示。 https://github.com/soulteary/docker-flare/issues/19 https://github.com/soulteary/docker-flare/issues/14#issuecomment-1044430872 (感谢 @lmm214 提醒 )
- [优化] 调整部分程序内部链接，并新增废弃提醒（可通过命令行或环境变量进行关闭），完善程序日志输出实现，重构部分功能，为之后完全开源做准备。
- [优化] 重构程序初始化相关功能，程序目前支持使用命令行参数和环境变量两种方式来进行初始化。 https://github.com/soulteary/docker-flare/issues/5 （感谢 @flashliao 建议）

# 0.2.8 - 2022-02-18

当前版本优化了三个用户反馈的小细节。

- [优化] 添加了搜索返回功能。 https://github.com/soulteary/docker-flare/issues/10 （感谢 @eallion 建议）
- [优化] 简单调整了移动端展示，下一步计划先开放模版和自定义主题。 https://github.com/soulteary/docker-flare/issues/6 （感谢 @flashliao 建议）
- [新增] 为容器镜像添加了 `latest` 标签。 https://github.com/soulteary/docker-flare/issues/8 https://github.com/soulteary/docker-flare/issues/7 （感谢 @viticis @eallion 建议）

## 0.2.7 - 2022-02-17

当前版本优化了一个小细节，以及修正了一个用户反馈的问题。

- [优化] 在非离线模式下，当用户首次使用程序，或程序配置中的 Location 字段为空的时候，将尝试使用 IPIP 的免费接口来自动进行定位，减少不必要的用户输入操作。（感谢 IPIP.net 高春辉大叔的帮忙！）
- [修正] 程序出现设置选项回显不正确的问题，https://github.com/soulteary/docker-flare/issues/7 （感谢 @viticis 小伙伴的反馈）

## 0.2.6 - 2022-02-17

当前版本新增了一个简单的功能，“使用向导”，相比较阅读文章，十几秒钟的交互引导会更直观一些。

- [新增] 用户使用向导，访问 `/guide` 即可使用，如果不需要这个功能，可以在环境变量中配置 `FLARE_GUIDE=0` 来关闭功能。

## 0.2.5 - 2022-02-16

当前版本调整内容比较多，程序配置文件新增了一些配置项，请留意文档或阅读博客文章。

如果你不希望手动调整程序配置文件，可以在备份当前数据目录 `app` 后，将其中的程序配置文件 `app/config.yml` 删除，重启应用，让程序自动生成新版配置文件。

- [新增] 用户登陆，让公网部署的小伙伴不必担心被陌生人随便修改配置的问题。
- [新增] 免登陆模式，让 HomeLab、本地运行 Flare 的小伙伴，保持使用上的简单。
- [新增] 离线模式，让 Flare 不访问任何公网资源。目前 Flare 仅会调用 issue[#4](https://github.com/soulteary/docker-flare/issues/4) 中提到的 `https://wis.qq.com/weather/common` 接口。
- [新增] 允许用户在页脚自定义展示内容，方便设置诸如个人信息、备案号、用户自己的网站统计脚本等。
- [新增] 允许用户对首页的问候语进行自定义设置，支持设置固定问候语，或根据早晨、中午、下午、晚上几个不同时间段展示不同的内容。
- [优化] 完善应用启动日志输出，方便用户反馈问题。
- [优化] 整理和重构部分代码实现，优化应用渲染性能。

## 0.2.4 - 2022-02-11

解决用户 @ember-zhang 提出的特殊链接渲染和跳转问题，支持了服务端跳转[#3](https://github.com/soulteary/docker-flare/issues/3)。

## 0.2.3 - 2022-02-07

针对应用链接的样式进行调整，让“应用”之间的距离稍微大一些。

修正了用户 @SandZhSand 反馈的在 Firefox 下“应用”书签的一个样式兼容性问题。

## 0.2.1 - 2022-02-07

完善构建脚本，支持在常见的 ARM 设备上直接运行。

## 0.2.0 - 2022-01-30

项目数据存储相关功能重构，使用 yaml 配置进行数据持久化，让应用更加轻量。

## 0.1.0 - 2022-01-27

项目初始化。编写简单的文档，完成基础展示功能，使用 SQLite 进行数据持久化，添加链接图标辅助工具。
