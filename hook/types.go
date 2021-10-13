package hook

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

// Modified from https://github.com/catsworld/qq-bot-api

// Params is the params of call api
type Params = zero.Params

// APIResponse is the response of calling API
// https://github.com/howmanybots/onebot/blob/master/v11/specs/communication/ws.md
type APIResponse = zero.APIResponse

// APIRequest is the request sending to the cqhttp
// https://github.com/howmanybots/onebot/blob/master/v11/specs/communication/ws.md
type APIRequest = zero.APIRequest

// User is a user on QQ.
type User = zero.User

// Event is the event emitted form cqhttp
type Event = zero.Event

// Message 消息
type Message = zero.Message

// File 文件
type File = zero.File

// Group 群
type Group = zero.Group

// H 是 Params 的简称
type H = Params
