package main

import "C"
import "unsafe"

//export Init
// 在此书写插件内容
func Init() {

}

//export Hook
// Hook 改变本插件的环境变量以加载插件
func Hook(botconf unsafe.Pointer, apicallers unsafe.Pointer, hooknew func() unsafe.Pointer) {
	BotConfig = (*Config)(botconf)
	APICallers = (*callerMap)(apicallers)
	hookNew = hooknew
}

func main() {
	// stub!
}
