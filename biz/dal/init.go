package dal

import (
	"github.com/juzi0911/gomall_PJ/biz/dal/mysql"
	"github.com/juzi0911/gomall_PJ/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
