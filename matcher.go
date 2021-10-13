package main

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

type (
	// Rule filter the event
	Rule func(ctx *Ctx) bool
	// Handler 事件处理函数
	Handler func(ctx *Ctx)
)

// Matcher 是 ZeroBot 匹配和处理事件的最小单元
type Matcher = zero.Matcher

// State store the context of a matcher.
type State = zero.State
