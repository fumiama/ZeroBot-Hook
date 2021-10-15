package message

import "unsafe"

// HookMsg hook message funcs
func HookMsg(custnode unsafe.Pointer, pasemsg unsafe.Pointer, parsemsgfromarr unsafe.Pointer) {
	CustomNode = *(*(func(nickname string, userID int64, content interface{}) MessageSegment))(unsafe.Pointer(&custnode))
	ParseMessage = *(*(func(msg []byte) Message))(unsafe.Pointer(&pasemsg))
	ParseMessageFromArray = *(*(func(msgs GjsonResult) Message))(unsafe.Pointer(&parsemsgfromarr))
}
