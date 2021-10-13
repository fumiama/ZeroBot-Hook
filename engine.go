package main

import (
	"unsafe"

	zero "github.com/wdvxdr1123/ZeroBot"
)

var hookNew func() unsafe.Pointer

// New 生成空引擎
func New() *Engine {
	return (*Engine)(hookNew())
}

// hook 的插件不允许使用 defaultEngine
// var defaultEngine *Engine

// Engine is the pre_handler, post_handler manager
type Engine = zero.Engine
