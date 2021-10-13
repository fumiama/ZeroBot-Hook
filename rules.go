package main

import (
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
var SuffixRule = zero.PrefixRule

// CommandRule check if the message is a command and trim the command name
var CommandRule = zero.PrefixRule

// RegexRule check if the message can be matched by the regex pattern
var RegexRule = zero.PrefixRule

// ReplyRule check if the message is replying some message
var ReplyRule = zero.PrefixRule

// KeywordRule check if the message has a keyword or keywords
var KeywordRule = zero.PrefixRule

// FullMatchRule check if src has the same copy of the message
var FullMatchRule = zero.PrefixRule

// OnlyToMe only triggered in conditions of @bot or begin with the nicknames
var OnlyToMe = zero.PrefixRule

// CheckUser only triggered by specific person
var CheckUser = zero.PrefixRule

// OnlyPrivate requires that the ctx.Event is private message
var OnlyPrivate = zero.PrefixRule

// OnlyGroup requires that the ctx.Event is public/group message
var OnlyGroup = zero.PrefixRule

// SuperUserPermission only triggered by the bot's owner
var SuperUserPermission = zero.PrefixRule

// AdminPermission only triggered by the group admins or higher permission
var AdminPermission = zero.PrefixRule

// OwnerPermission only triggered by the group owner or higher permission
var OwnerPermission = zero.PrefixRule
