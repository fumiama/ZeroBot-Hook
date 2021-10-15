package main

import "C"
import (
	"fmt"
	"unsafe"

	ctrl "github.com/fumiama/ZeroBot-Hook/control"
	zero "github.com/fumiama/ZeroBot-Hook/hook"
	"github.com/fumiama/ZeroBot-Hook/hook/message"
)

var options = ctrl.Options{
	DisableOnDefault: false,
	Help:             "help from demo",
}

//export Inita
func Inita() {
	// --------------------------在此下书写插件内容--------------------------
	en := ctrl.Register("demo", &options)
	en.OnCommand("demo", zero.AdminPermission).SetBlock(true).SecondPriority().
		Handle(func(ctx *zero.Ctx) {
			fmt.Println("msg recv.")
			ctx.SendChain(message.Text("回复"))
		})
	// --------------------------在此上书写插件内容--------------------------
}

// 以下勿动
// Hook 改变本插件的环境变量以加载插件
//export Hook
func Hook(botconf unsafe.Pointer, apicallers unsafe.Pointer, hooknew unsafe.Pointer,
	matlist unsafe.Pointer, matlock unsafe.Pointer, defen unsafe.Pointer,
	reg unsafe.Pointer, del unsafe.Pointer,
	sndgrpmsg unsafe.Pointer, sndprivmsg unsafe.Pointer, getmsg unsafe.Pointer,
	parsectx unsafe.Pointer,
	custnode unsafe.Pointer, pasemsg unsafe.Pointer, parsemsgfromarr unsafe.Pointer,
) {
	zero.Hook(botconf, apicallers, hooknew, matlist, matlock, defen)
	ctrl.Register = *(*(func(service string, o *ctrl.Options) *zero.Engine))(unsafe.Pointer(&reg))
	ctrl.Delete = *(*(func(service string)))(unsafe.Pointer(&del))
	zero.HookCtx(sndgrpmsg, sndprivmsg, getmsg, parsectx)
	message.HookMsg(custnode, pasemsg, parsemsgfromarr)
	// fmt.Printf("[plugin]set reg: %x, del: %x\n", ctrl.Register, ctrl.Delete)
}

func main() {
	// stub!
}

type String struct {
	Data unsafe.Pointer
	Len  int
}
