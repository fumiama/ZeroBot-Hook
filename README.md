# ZeroBot-Hook
为 ZeroBot-Plugin 生成动态库插件。

# 使用方法
### 编写插件
复制`main.go`到本地，修改`main.go`的`Init`函数为插件内容，语法与[ZeroBot](https://github.com/wdvxdr1123/ZeroBot)基本相同。
### 编译为动态库
- 使用`Actions`
只要创建形如`v1.2.3`的`tag`，即可触发插件编译流程。编译好后前往`Release`页面下载即可。
- 本地编译
```bash
# 本机架构
go build -ldflags "-s -w" -buildmode=plugin -o demo.so
# 交叉编译：详见 workflow 相关代码
CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 CC=arm-linux-gnueabihf-gcc-9 CXX=g++-9-arm-linux-gnueabihf go build -ldflags="-s -w" -buildmode=plugin -o artifacts/zbpd-linux-armv6
```
### 开始使用
放置动态库到[ZeroBot-Plugin](https://github.com/FloatTech/ZeroBot-Plugin)的`plugins/`目录下，给机器人发送`/刷新插件`即可，或重启也可加载。
