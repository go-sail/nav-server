package schedule

import (
	"github.com/keepchen/go-sail/v3/sail"
	"nav-server/app/admin/config"
	"nav-server/app/admin/http/service"
	"time"
)

func SaveNavListToMemory() {
	//服务启动的时候执行一次
	saveNavListToMemory()

	sail.Schedule(sail.Utils().String().WrapRedisKey(config.Get().AppName, "SaveNavListToMemory"), saveNavListToMemory).
		WithoutOverlapping().
		Every(time.Minute) //每分钟执行一次
}

func saveNavListToMemory() {
	service.Index.SaveNavCategoryToMem()
}
