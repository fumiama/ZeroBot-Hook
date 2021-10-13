package hook

import (
	"sync"
	"unsafe"

	zero "github.com/wdvxdr1123/ZeroBot"
)

// Config is config of zero bot
type Config = zero.Config

// APICallers 所有的APICaller列表， 通过self-ID映射
var APICallers *callerMap

// APICaller is the interface of CallApi
type APICaller interface {
	CallApi(request APIRequest) (APIResponse, error)
}

// Driver 与OneBot通信的驱动，使用driver.DefaultWebSocketDriver
type Driver interface {
	Connect()
	Listen(func([]byte, APICaller))
}

// BotConfig 运行中bot的配置，是Run函数的参数的拷贝
var BotConfig *Config

// 没有方法的interface
type eface struct {
	_type uintptr
	data  unsafe.Pointer
}

func getdata(ptr *interface{}) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(ptr)).data
}

// Hook 改变本插件的环境变量以加载插件
func Hook(botconf interface{}, apicallers interface{}, hooknew interface{}, matlist interface{}, matlock interface{}) {
	BotConfig = (*Config)(getdata(&botconf))
	APICallers = (*callerMap)(getdata(&apicallers))
	n := getdata(&hooknew)
	New = *(*(func() *Engine))(unsafe.Pointer(&n))
	matcherList = (*[]*Matcher)(getdata(&matlist))
	matcherLock = (*sync.RWMutex)(getdata(&matlock))
	// fmt.Printf("[plugin]matlist: %p, matlock: %p\n", matcherList, matcherLock)
}

// GetBot 获取指定的bot (Ctx)实例
func GetBot(id int64) *Ctx {
	caller, ok := APICallers.Load(id)
	if !ok {
		return nil
	}
	return (*Ctx)(unsafe.Pointer(&Ctx{caller: caller}))
}

// RangeBot 遍历所有bot (Ctx)实例
//
// 单次操作返回 true 则继续遍历，否则退出
func RangeBot(iter func(id int64, ctx *Ctx) bool) {
	APICallers.Range(func(key int64, value APICaller) bool {
		return iter(key, (*Ctx)(unsafe.Pointer(&Ctx{caller: value})))
	})
}
