# Flare ✨

兼顾轻量快速、美观的个人导航页面，适用于 HomeLab 或其他场景。

应用资源消耗非常低：

- CPU: < 1%
- MEM: < 30M
- Docker Image: ~10M

## 如何使用

根据你的需求调整 `app/*yml` 文件，然后执行：

```bash
docker-compose up -d
```

![](./screenshots/ui.png)

![](./screenshots/lighthouse.png)

## 相比较 Flame

- 服务资源消耗极低，可以跑在任何规格的机器上。
- 程序页面性能非常好，渲染速度更快，支持渲染大量书签，而不必担心风扇起飞。
- 使用声明的配置来进行导航内容的管理，无需担心数据迁移问题。
- 简化了天气数据的获取方式，不再需要申请天气网站的 `API_KEY` ，避免了不必要的成本开销。
- 简化了 Flame 中的登陆、K8S、Docker 集成等不必要的功能。

## 支持图标

支持所有的 Material Design Icons，你可以在 [materialdesignicons.com](https://materialdesignicons.com/) 或 [pictogrammers.github.io/@mdi](https://pictogrammers.github.io/@mdi/font/6.5.95/) 6700 个精心设计的图标中找到你喜欢的图标。

在配置中使用的图标名称无需使用连字符 `-`，直接输入名称即可。（大小写不敏感）


## TODO

- [ ] 将程序完全配置化
- [ ] 完善程序设置功能
- [ ] 支持使用 API 进行内容管理

## Thanks

Inspired by https://github.com/pawelmalak/flame
