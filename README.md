# ZeroBot-Hook
Vars &amp; funcs' hook of ZeroBot to gen plugin lib.

# Compile
1. Normally
```bash
go build -x -v -ldflags "-s -w" -buildmode=c-shared
```
2. With `gccgo`
```bash
go build -x -v -ldflags "-s -w" -buildmode=c-shared -compiler gccgo
```