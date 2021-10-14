package message

import (
	"strings"
)

// GjsonResult represents a json value that is returned from Get().
// 使用时请将其强制转换为 gjson.Result 类型
type GjsonResult struct {
	// Type is the json type
	Type int
	// Raw is the raw json
	Raw string
	// Str is the json string
	Str string
	// Num is the json number
	Num float64
	// Index of raw value in original json, zero means index unknown
	Index int
	// Indexes of all the elements that match on a path containing the '#'
	// query character.
	Indexes []int
}

// Modified from https://github.com/catsworld/qq-bot-api

// ParseMessage parses msg, which might have 2 types, string or array,
// depending on the configuration of cqhttp, to a Message.
// msg is the value of key "message" of the data unmarshalled from the
// API response JSON.
var ParseMessage func(msg []byte) Message

// ParseMessageFromArray parses msg as type array to a Message.
// msg is the value of key "message" of the data unmarshalled from the
// API response JSON.
// ParseMessageFromArray cq字符串转化为json对象
var ParseMessageFromArray func(msgs GjsonResult) Message

// CQString 转为CQ字符串
// Deprecated: use method String instead
func (m Message) CQString() string {
	return m.String()
}

// ExtractPlainText 提取消息中的纯文本
func (m Message) ExtractPlainText() string {
	sb := strings.Builder{}
	for _, val := range m {
		if val.Type == "text" {
			sb.WriteString(val.Data["text"])
		}
	}
	return sb.String()
}
