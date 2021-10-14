package hook

import (
	"sync"

	"github.com/fumiama/ZeroBot-Hook/hook/message"
)

// Ctx represents the Context which hold the event.
// 代表上下文
type Ctx struct {
	ma     *Matcher
	Event  *Event
	State  State
	caller APICaller

	// lazy message
	once    sync.Once
	message string
}

// GetMatcher ...
func (ctx *Ctx) GetMatcher() *Matcher {
	return ctx.ma
}

var parse func(ctx *Ctx, model interface{}) (err error)

// Parse 将 Ctx.State 映射到结构体
func (ctx *Ctx) Parse(model interface{}) (err error) {
	return parse(ctx, model)
}

// CheckSession 判断会话连续性
func (ctx *Ctx) CheckSession() Rule {
	return func(ctx2 *Ctx) bool {
		return ctx.Event.UserID == ctx2.Event.UserID &&
			ctx.Event.GroupID == ctx2.Event.GroupID // 私聊时GroupID为0，也相等
	}
}

// Send 快捷发送消息
func (ctx *Ctx) Send(message interface{}) int64 {
	if ctx.Event.GroupID != 0 {
		return ctx.SendGroupMessage(ctx.Event.GroupID, message)
	}
	return ctx.SendPrivateMessage(ctx.Event.UserID, message)
}

// SendChain 快捷发送消息-消息链
func (ctx *Ctx) SendChain(message ...message.MessageSegment) int64 {
	if ctx.Event.GroupID != 0 {
		return ctx.SendGroupMessage(ctx.Event.GroupID, message)
	}
	return ctx.SendPrivateMessage(ctx.Event.UserID, message)
}

// FutureEvent ...
func (ctx *Ctx) FutureEvent(Type string, rule ...Rule) *FutureEvent {
	return ctx.ma.FutureEvent(Type, rule...)
}

// ExtractPlainText 提取消息中的纯文本
func (ctx *Ctx) ExtractPlainText() string {
	if ctx == nil || ctx.Event == nil || ctx.Event.Message == nil {
		return ""
	}
	return ctx.Event.Message.ExtractPlainText()
}

// Block 阻止后续触发
func (ctx *Ctx) Block() {
	ctx.ma.SetBlock(true)
}

// MessageString 字符串消息便于Regex
func (ctx *Ctx) MessageString() string {
	ctx.once.Do(func() {
		if ctx.Event != nil && ctx.Event.Message != nil {
			ctx.message = ctx.Event.Message.String()
		}
	})
	return ctx.message
}
