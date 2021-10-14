package hook

import (
	"encoding/json"
	"unsafe"

	"github.com/fumiama/ZeroBot-Hook/hook/message"
)

// GjsonResult represents a json value that is returned from Get().
// 使用时请将其强制转换为 gjson.Result 类型
type GjsonResult = message.GjsonResult

func HookCtx(sndgrpmsg interface{}, sndprivmsg interface{}, getmsg interface{}, parsectx interface{}) {
	d1 := getdata(&sndgrpmsg)
	d2 := getdata(&sndprivmsg)
	d3 := getdata(&getmsg)
	d4 := getdata(&parsectx)
	sendGroupMessage = *(*(func(ctx *Ctx, groupID int64, message interface{}) int64))(unsafe.Pointer(&d1))
	sendPrivateMessage = *(*(func(ctx *Ctx, groupID int64, message interface{}) int64))(unsafe.Pointer(&d2))
	getMessage = *(*(func(ctx *Ctx, messageID int64) Message))(unsafe.Pointer(&d3))
	parse = *(*(func(ctx *Ctx, model interface{}) (err error)))(unsafe.Pointer(&d4))
}

var sendGroupMessage func(ctx *Ctx, groupID int64, message interface{}) int64
var sendPrivateMessage func(ctx *Ctx, userID int64, message interface{}) int64
var getMessage func(ctx *Ctx, messageID int64) Message

// CallAction 调用 cqhttp API
func (ctx *Ctx) CallAction(action string, params Params) APIResponse {
	req := APIRequest{
		Action: action,
		Params: params,
	}
	rsp, _ := ctx.caller.CallApi(req)
	return rsp
}

// SendGroupMessage 发送群消息
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#send_group_msg-%E5%8F%91%E9%80%81%E7%BE%A4%E6%B6%88%E6%81%AF
func (ctx *Ctx) SendGroupMessage(groupID int64, message interface{}) int64 {
	return sendGroupMessage(ctx, groupID, message)
}

// SendPrivateMessage 发送私聊消息
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#send_private_msg-%E5%8F%91%E9%80%81%E7%A7%81%E8%81%8A%E6%B6%88%E6%81%AF
func (ctx *Ctx) SendPrivateMessage(userID int64, message interface{}) int64 {
	return sendPrivateMessage(ctx, userID, message)
}

// DeleteMessage 撤回消息
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#delete_msg-%E6%92%A4%E5%9B%9E%E6%B6%88%E6%81%AF
func (ctx *Ctx) DeleteMessage(messageID int64) {
	ctx.CallAction("delete_msg", Params{
		"message_id": messageID,
	})
}

// GetMessage 获取消息
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_msg-%E8%8E%B7%E5%8F%96%E6%B6%88%E6%81%AF
func (ctx *Ctx) GetMessage(messageID int64) Message {
	return getMessage(ctx, messageID)
}

// GetForwardMessage 获取合并转发消息
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_forward_msg-%E8%8E%B7%E5%8F%96%E5%90%88%E5%B9%B6%E8%BD%AC%E5%8F%91%E6%B6%88%E6%81%AF
func (ctx *Ctx) GetForwardMessage(id int64) GjsonResult {
	rsp := ctx.CallAction("get_forward_msg", Params{
		"id": id,
	}).Data
	return rsp
}

// SetGroupKick 群组踢人
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#set_group_kick-%E7%BE%A4%E7%BB%84%E8%B8%A2%E4%BA%BA
func (ctx *Ctx) SetGroupKick(groupID, userID int64, rejectAddRequest bool) {
	ctx.CallAction("set_group_kick", Params{
		"group_id":           groupID,
		"user_id":            userID,
		"reject_add_request": rejectAddRequest,
	})
}

// SetGroupBan 群组单人禁言
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#set_group_ban-%E7%BE%A4%E7%BB%84%E5%8D%95%E4%BA%BA%E7%A6%81%E8%A8%80
func (ctx *Ctx) SetGroupBan(groupID, userID, duration int64) {
	ctx.CallAction("set_group_ban", Params{
		"group_id": groupID,
		"user_id":  userID,
		"duration": duration,
	})
}

// SetGroupWholeBan 群组全员禁言
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#set_group_whole_ban-%E7%BE%A4%E7%BB%84%E5%85%A8%E5%91%98%E7%A6%81%E8%A8%80
func (ctx *Ctx) SetGroupWholeBan(groupID int64, enable bool) {
	ctx.CallAction("set_group_whole_ban", Params{
		"group_id": groupID,
		"enable":   enable,
	})
}

// SetGroupAdmin 群组设置管理员
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#set_group_whole_ban-%E7%BE%A4%E7%BB%84%E5%85%A8%E5%91%98%E7%A6%81%E8%A8%80
func (ctx *Ctx) SetGroupAdmin(groupID, userID int64, enable bool) {
	ctx.CallAction("set_group_admin", Params{
		"group_id": groupID,
		"user_id":  userID,
		"enable":   enable,
	})
}

// SetGroupAnonymous 群组匿名
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#set_group_anonymous-%E7%BE%A4%E7%BB%84%E5%8C%BF%E5%90%8D
func (ctx *Ctx) SetGroupAnonymous(groupID int64, enable bool) {
	ctx.CallAction("set_group_anonymous", Params{
		"group_id": groupID,
		"enable":   enable,
	})
}

// SetGroupCard 设置群名片（群备注）
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#set_group_card-%E8%AE%BE%E7%BD%AE%E7%BE%A4%E5%90%8D%E7%89%87%E7%BE%A4%E5%A4%87%E6%B3%A8
func (ctx *Ctx) SetGroupCard(groupID, userID int64, card string) {
	ctx.CallAction("set_group_card", Params{
		"group_id": groupID,
		"user_id":  userID,
		"card":     card,
	})
}

// SetGroupName 设置群名
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#set_group_name-%E8%AE%BE%E7%BD%AE%E7%BE%A4%E5%90%8D
func (ctx *Ctx) SetGroupName(groupID int64, groupName string) {
	ctx.CallAction("set_group_name", Params{
		"group_id":   groupID,
		"group_name": groupName,
	})
}

// SetGroupLeave 退出群组
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#set_group_leave-%E9%80%80%E5%87%BA%E7%BE%A4%E7%BB%84
func (ctx *Ctx) SetGroupLeave(groupID int64, isDismiss bool) {
	ctx.CallAction("set_group_leave", Params{
		"group_id":   groupID,
		"is_dismiss": isDismiss,
	})
}

// SetGroupSpecialTitle 设置群组专属头衔
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#set_group_special_title-%E8%AE%BE%E7%BD%AE%E7%BE%A4%E7%BB%84%E4%B8%93%E5%B1%9E%E5%A4%B4%E8%A1%94
func (ctx *Ctx) SetGroupSpecialTitle(groupID int64, userID int64, specialTitle string) {
	ctx.CallAction("set_group_special_title", Params{
		"group_id":      groupID,
		"user_id":       userID,
		"special_title": specialTitle,
	})
}

// SetFriendAddRequest 处理加好友请求
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#set_friend_add_request-%E5%A4%84%E7%90%86%E5%8A%A0%E5%A5%BD%E5%8F%8B%E8%AF%B7%E6%B1%82
func (ctx *Ctx) SetFriendAddRequest(flag string, approve bool, remark string) {
	ctx.CallAction("set_friend_add_request", Params{
		"flag":    flag,
		"approve": approve,
		"remark":  remark,
	})
}

// SetGroupAddRequest 处理加群请求／邀请
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#set_group_add_request-%E5%A4%84%E7%90%86%E5%8A%A0%E7%BE%A4%E8%AF%B7%E6%B1%82%E9%82%80%E8%AF%B7
func (ctx *Ctx) SetGroupAddRequest(flag string, subType string, approve bool, reason string) {
	ctx.CallAction("set_group_add_request", Params{
		"flag":     flag,
		"sub_type": subType,
		"approve":  approve,
		"reason":   reason,
	})
}

// GetLoginInfo 获取登录号信息
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_login_info-%E8%8E%B7%E5%8F%96%E7%99%BB%E5%BD%95%E5%8F%B7%E4%BF%A1%E6%81%AF
func (ctx *Ctx) GetLoginInfo() GjsonResult {
	return ctx.CallAction("get_login_info", Params{}).Data
}

// GetStrangerInfo 获取陌生人信息
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_stranger_info-%E8%8E%B7%E5%8F%96%E9%99%8C%E7%94%9F%E4%BA%BA%E4%BF%A1%E6%81%AF
func (ctx *Ctx) GetStrangerInfo(userID int64, noCache bool) GjsonResult {
	return ctx.CallAction("get_stranger_info", Params{
		"user_id":  userID,
		"no_cache": noCache,
	}).Data
}

// GetFriendList 获取好友列表
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_friend_list-%E8%8E%B7%E5%8F%96%E5%A5%BD%E5%8F%8B%E5%88%97%E8%A1%A8
func (ctx *Ctx) GetFriendList() GjsonResult {
	return ctx.CallAction("get_friend_list", Params{}).Data
}

// GetGroupInfo 获取群信息
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_group_info-%E8%8E%B7%E5%8F%96%E7%BE%A4%E4%BF%A1%E6%81%AF
func (ctx *Ctx) GetGroupInfo(groupID int64, noCache bool) Group {
	rsp := ctx.CallAction("get_group_info", Params{
		"group_id": groupID,
		"no_cache": noCache,
	}).Data
	group := Group{}
	_ = json.Unmarshal([]byte(rsp.Raw), &group)
	return group
}

// GetGroupList 获取群列表
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_group_list-%E8%8E%B7%E5%8F%96%E7%BE%A4%E5%88%97%E8%A1%A8
func (ctx *Ctx) GetGroupList() GjsonResult {
	return ctx.CallAction("get_group_list", Params{}).Data
}

// GetGroupMemberInfo 获取群成员信息
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_group_member_info-%E8%8E%B7%E5%8F%96%E7%BE%A4%E6%88%90%E5%91%98%E4%BF%A1%E6%81%AF
func (ctx *Ctx) GetGroupMemberInfo(groupID int64, userID int64, noCache bool) GjsonResult {
	return ctx.CallAction("get_group_member_info", Params{
		"group_id": groupID,
		"user_id":  userID,
		"no_cache": noCache,
	}).Data
}

// GetGroupMemberList 获取群成员列表
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_group_member_list-%E8%8E%B7%E5%8F%96%E7%BE%A4%E6%88%90%E5%91%98%E5%88%97%E8%A1%A8
func (ctx *Ctx) GetGroupMemberList(groupID int64) GjsonResult {
	return ctx.CallAction("get_group_member_list", Params{
		"group_id": groupID,
	}).Data
}

// GetGroupHonorInfo 获取群荣誉信息
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_group_honor_info-%E8%8E%B7%E5%8F%96%E7%BE%A4%E8%8D%A3%E8%AA%89%E4%BF%A1%E6%81%AF
func (ctx *Ctx) GetGroupHonorInfo(groupID int64, hType string) GjsonResult {
	return ctx.CallAction("get_group_honor_info", Params{
		"group_id": groupID,
		"type":     hType,
	}).Data
}

// GetRecord 获取语音
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_record-%E8%8E%B7%E5%8F%96%E8%AF%AD%E9%9F%B3
func (ctx *Ctx) GetRecord(file string, outFormat string) GjsonResult {
	return ctx.CallAction("get_record", Params{
		"file":       file,
		"out_format": outFormat,
	}).Data
}

// GetImage 获取图片
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_image-%E8%8E%B7%E5%8F%96%E5%9B%BE%E7%89%87
func (ctx *Ctx) GetImage(file string) GjsonResult {
	return ctx.CallAction("get_image", Params{
		"file": file,
	}).Data
}

// GetVersionInfo 获取运行状态
// https://github.com/howmanybots/onebot/blob/master/v11/specs/api/public.md#get_status-%E8%8E%B7%E5%8F%96%E8%BF%90%E8%A1%8C%E7%8A%B6%E6%80%81
func (ctx *Ctx) GetVersionInfo() GjsonResult {
	return ctx.CallAction("get_version_info", Params{}).Data
}

// Expand API

// SetGroupPortrait 设置群头像
// https://github.com/Mrs4s/go-cqhttp/blob/master/docs/cqhttp.md#%E8%AE%BE%E7%BD%AE%E7%BE%A4%E5%A4%B4%E5%83%8F
func (ctx *Ctx) SetGroupPortrait(groupID int64, file string) {
	ctx.CallAction("set_group_portrait", Params{
		"group_id": groupID,
		"file":     file,
	})
}

// OCRImage 图片OCR
// https://github.com/Mrs4s/go-cqhttp/blob/master/docs/cqhttp.md#%E5%9B%BE%E7%89%87ocr
func (ctx *Ctx) OCRImage(file string) GjsonResult {
	return ctx.CallAction("ocr_image", Params{
		"image": file,
	}).Data
}

// SendGroupForwardMessage 发送合并转发(群)
// https://github.com/Mrs4s/go-cqhttp/blob/master/docs/cqhttp.md#%E5%9B%BE%E7%89%87ocr
func (ctx *Ctx) SendGroupForwardMessage(groupID int64, message message.Message) GjsonResult {
	return ctx.CallAction("send_group_forward_msg", Params{
		"group_id": groupID,
		"messages": message,
	}).Data
}

// GetGroupSystemMessage 获取群系统消息
// https://github.com/Mrs4s/go-cqhttp/blob/master/docs/cqhttp.md#%E8%8E%B7%E5%8F%96%E7%BE%A4%E7%B3%BB%E7%BB%9F%E6%B6%88%E6%81%AF
func (ctx *Ctx) GetGroupSystemMessage() GjsonResult {
	return ctx.CallAction("get_group_system_msg", Params{}).Data
}

// GetWordSlices 获取中文分词
// https://github.com/Mrs4s/go-cqhttp/blob/master/docs/cqhttp.md#%E8%8E%B7%E5%8F%96%E4%B8%AD%E6%96%87%E5%88%86%E8%AF%8D
func (ctx *Ctx) GetWordSlices(content string) GjsonResult {
	return ctx.CallAction(".get_word_slices", Params{
		"content": content,
	}).Data
}
