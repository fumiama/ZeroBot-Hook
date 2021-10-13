package main

import "C"
import (
	"fmt"
	"unsafe"

	ctrl "github.com/fumiama/ZeroBot-Hook/control"
	zero "github.com/fumiama/ZeroBot-Hook/hook"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func Inita(reg interface{}, del interface{}) {
	rd := getdata(&reg)
	dd := getdata(&del)
	ctrl.Register = *(*(func(service string, o *ctrl.Options) *zero.Engine))(unsafe.Pointer(&rd))
	ctrl.Delete = *(*(func(service string)))(unsafe.Pointer(&dd))
	// fmt.Printf("[plugin]set reg: %x, del: %x\n", ctrl.Register, ctrl.Delete)
	// -------------在此下书写插件内容-------------
	en := ctrl.Register("demo", &ctrl.Options{
		DisableOnDefault: false,
		Help:             "help from demo",
	})
	en.OnCommand("demo", zero.AdminPermission).SetBlock(true).SecondPriority().
		Handle(func(ctx *zero.Ctx) {
			fmt.Println("msg recv.")
			ctx.SendChain(message.Text("回复"))
		})
	// -------------在此上书写插件内容-------------
}

// 以下勿动
// Hook 改变本插件的环境变量以加载插件
func Hook(botconf interface{}, apicallers interface{}, hooknew interface{}, matlist interface{}, matlock interface{}) {
	zero.Hook(botconf, apicallers, hooknew, matlist, matlock)
}

// 没有方法的interface
type eface struct {
	_type uintptr
	data  unsafe.Pointer
}

func getdata(ptr *interface{}) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(ptr)).data
}

func main() {
	// stub!
}
