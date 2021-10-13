package main

import "C"
import (
	"fmt"
	"unsafe"

	ctrl "github.com/fumiama/ZeroBot-Hook/control"
	zero "github.com/fumiama/ZeroBot-Hook/hook"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func Init(reg uintptr, del uintptr) {
	go func() {
		fmt.Println("enter init.")
		ctrl.Register = *(*(func(service string, o *ctrl.Options) *zero.Engine))(unsafe.Pointer(reg))
		ctrl.Delete = *(*(func(service string)))(unsafe.Pointer(del))
		// -------------在此下书写插件内容-------------
		fmt.Println("fill control success.")
		en := ctrl.Register("demo", &ctrl.Options{
			DisableOnDefault: false,
			Help:             "help from demo",
		})
		fmt.Println("register demo success.")
		en.OnCommand("demo", zero.AdminPermission).SetBlock(true).SecondPriority().
			Handle(func(ctx *zero.Ctx) {
				fmt.Println("msg recv.")
				ctx.SendChain(message.Text("回复"))
			})
		fmt.Println("register matcher success.")
		// -------------在此上书写插件内容-------------
		fmt.Println("quit init.")
	}()
}

// 以下勿动
// Hook 改变本插件的环境变量以加载插件
func Hook(botconf unsafe.Pointer, apicallers unsafe.Pointer, hooknew unsafe.Pointer) {
	fmt.Println("enter hook.")
	zero.Hook(botconf, apicallers, hooknew)
	fmt.Println("quit hook.")
}

func main() {
	// stub!
}
