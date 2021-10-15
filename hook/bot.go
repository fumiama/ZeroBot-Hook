package hook

import (
	"sync"
	"unsafe"
)

// Config is config of zero bot
type Config struct {
	NickName      []string `json:"nickname"`       // 机器人名称
	CommandPrefix string   `json:"command_prefix"` // 触发命令
	SuperUsers    []string `json:"super_users"`    // 超级用户
	SelfID        string   `json:"self_id"`        // 机器人账号
	Driver        []Driver `json:"-"`              // 通信驱动
}

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

// Hook 改变本插件的环境变量以加载插件
func Hook(botconf unsafe.Pointer, apicallers unsafe.Pointer, hooknew unsafe.Pointer, matlist unsafe.Pointer, matlock unsafe.Pointer, defen unsafe.Pointer) {
	BotConfig = (*Config)(botconf)
	APICallers = (*callerMap)(apicallers)
	New = *(*(func() *Engine))(unsafe.Pointer(&hooknew))
	matcherList = (*[]*Matcher)(matlist)
	matcherLock = (*sync.RWMutex)(matlock)
	defaultEngine = (*Engine)(defen)
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
