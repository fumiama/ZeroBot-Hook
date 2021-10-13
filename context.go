package main

import (
	"sync"

	zero "github.com/wdvxdr1123/ZeroBot"
)

// Ctx represents the Context which hold the event.
// 代表上下文
type Ctx = zero.Ctx

// context 用来暴露内部成员
type context struct {
	ma     *Matcher
	Event  *Event
	State  State
	caller APICaller

	// lazy message
	once    sync.Once
	message string
}
