# ZeroBot-Hook
为 ZeroBot-Plugin 生成动态库插件。

# 使用方法
### 编写插件
修改`main.go`的`Init`函数为插件内容，语法与[ZeroBot](https://github.com/wdvxdr1123/ZeroBot)基本相同。
### 编译为动态库
1. 使用`gc`
```bash
go build -x -v -ldflags "-s -w" -buildmode=c-shared
```
2. 使用`gccgo`
```bash
go build -x -v -ldflags "-s -w" -buildmode=c-shared -compiler gccgo
```
### 开始使用
放置动态库到[ZeroBot-Plugin](https://github.com/FloatTech/ZeroBot-Plugin)的`plugins/`目录下，给机器人发送`/刷新插件`即可，或重启也可加载。