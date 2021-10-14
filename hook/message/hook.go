package message

import "unsafe"

// var CustomNode func(nickname string, userID int64, content interface{}) MessageSegment
// var ParseMessage func(msg []byte) Message
// var ParseMessageFromArray func(msgs GjsonResult) Message
func HookMsg(custnode interface{}, pasemsg interface{}, parsemsgfromarr interface{}) {
	d1 := getdata(&custnode)
	d2 := getdata(&pasemsg)
	d3 := getdata(&parsemsgfromarr)
	CustomNode = *(*(func(nickname string, userID int64, content interface{}) MessageSegment))(unsafe.Pointer(&d1))
	ParseMessage = *(*(func(msg []byte) Message))(unsafe.Pointer(&d2))
	ParseMessageFromArray = *(*(func(msgs GjsonResult) Message))(unsafe.Pointer(&d3))
}

// 没有方法的interface
type eface struct {
	_type uintptr
	data  unsafe.Pointer
}

func getdata(ptr *interface{}) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(ptr)).data
}
