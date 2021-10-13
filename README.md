# ZeroBot-Hook
为 ZeroBot-Plugin 生成动态库插件。

# 使用方法
### 编写插件
修改`main.go`的`Init`函数为插件内容。
### 编译为动态库
1. 使用`gc`
```bash
go build -x -v -ldflags "-s -w" -buildmode=c-shared
```
2. 使用`gccgo`
```bash
go build -x -v -ldflags "-s -w" -buildmode=c-shared -compiler gccgo
```