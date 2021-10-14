# ZeroBot-Hook
为 ZeroBot-Plugin 生成动态库插件。

# 使用方法
### 编写插件
复制`main.go`到本地，修改`main.go`的`Init`函数为插件内容，语法与[ZeroBot](https://github.com/wdvxdr1123/ZeroBot)基本相同。
### 编译为动态库
在本地运行
```bash
go build -ldflags "-s -w" -buildmode=plugin -o demo.so
```
### 开始使用
放置动态库到[ZeroBot-Plugin](https://github.com/FloatTech/ZeroBot-Plugin)的`plugins/`目录下，给机器人发送`/刷新插件`即可，或重启也可加载。
