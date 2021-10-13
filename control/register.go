package control

import (
	zero "github.com/fumiama/ZeroBot-Hook/hook"
)

// Register 注册插件控制器
var Register func(service string, o *Options) *zero.Engine

// Delete 删除插件控制器，不会删除数据
var Delete func(service string)
