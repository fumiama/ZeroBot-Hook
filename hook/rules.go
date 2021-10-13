package hook

import (
	"strconv"
	"strings"
	"unsafe"

	zero "github.com/wdvxdr1123/ZeroBot"
)

// Type check the ctx.Event's type
var Type = zero.Type

// PrefixRule check if the message has the prefix and trim the prefix
//
// 检查消息前缀
var PrefixRule = zero.PrefixRule

// SuffixRule check if the message has the suffix and trim the suffix
//
// 检查消息后缀
var SuffixRule = zero.SuffixRule

// CommandRule check if the message is a command and trim the command name
func CommandRule(commands ...string) Rule {
	f := func(ctx *Ctx) bool {
		if len(ctx.Event.Message) == 0 || ctx.Event.Message[0].Type != "text" {
			return false
		}
		first := ctx.Event.Message[0]
		firstMessage := first.Data["text"]
		if !strings.HasPrefix(firstMessage, BotConfig.CommandPrefix) {
			return false
		}
		cmdMessage := firstMessage[len(BotConfig.CommandPrefix):]
		for _, command := range commands {
			if strings.HasPrefix(cmdMessage, command) {
				ctx.State["command"] = command
				arg := strings.TrimLeft(cmdMessage[len(command):], " ")
				if len(ctx.Event.Message) > 1 {
					arg += ctx.Event.Message[1:].ExtractPlainText()
				}
				ctx.State["args"] = arg
				return true
			}
		}
		return false
	}
	return *(*zero.Rule)(unsafe.Pointer(&f))
}

// RegexRule check if the message can be matched by the regex pattern
var RegexRule = zero.RegexRule

// ReplyRule check if the message is replying some message
var ReplyRule = zero.ReplyRule

// KeywordRule check if the message has a keyword or keywords
var KeywordRule = zero.KeywordRule

// FullMatchRule check if src has the same copy of the message
var FullMatchRule = zero.FullMatchRule

// OnlyToMe only triggered in conditions of @bot or begin with the nicknames
var OnlyToMe = zero.OnlyToMe

// CheckUser only triggered by specific person
var CheckUser = zero.CheckUser

// OnlyPrivate requires that the ctx.Event is private message
var OnlyPrivate = zero.OnlyPrivate

// OnlyGroup requires that the ctx.Event is public/group message
var OnlyGroup = zero.OnlyGroup

// SuperUserPermission only triggered by the bot's owner
func SuperUserPermission(ctx *Ctx) bool {
	uid := strconv.FormatInt(ctx.Event.UserID, 10)
	for _, su := range BotConfig.SuperUsers {
		if su == uid {
			return true
		}
	}
	return false
}

// AdminPermission only triggered by the group admins or higher permission
var AdminPermission = zero.AdminPermission

// OwnerPermission only triggered by the group owner or higher permission
var OwnerPermission = zero.OwnerPermission
