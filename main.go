package main

import "C"
import (
	"unsafe"

	ctrl "github.com/fumiama/ZeroBot-Hook/control"
	zero "github.com/fumiama/ZeroBot-Hook/hook"
	"github.com/wdvxdr1123/ZeroBot/message"
)

//export Init
func Init(reg unsafe.Pointer, del unsafe.Pointer) {
	ctrl.Register = *(*(func(service string, o *ctrl.Options) *zero.Engine))(reg)
	ctrl.Delete = *(*(func(service string)))(del)
	// -------------在此书写插件内容-------------
	en := ctrl.Register("demo", &ctrl.Options{
		DisableOnDefault: false,
		Help:             "help from demo",
	})
	en.OnCommand("demo", zero.AdminPermission).SetBlock(true).SecondPriority().
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("回复"))
		})
	// -------------在此书写插件内容-------------
}

// 以下勿动
//export Hook
// Hook 改变本插件的环境变量以加载插件
func Hook(botconf unsafe.Pointer, apicallers unsafe.Pointer, hooknew unsafe.Pointer, defaultengine unsafe.Pointer) {
	zero.Hook(botconf, apicallers, hooknew, defaultengine)
}

func main() {
	// stub!
}
