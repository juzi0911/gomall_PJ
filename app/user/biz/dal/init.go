package dal

import (
	"github.com/juzi0911/gomall_PJ/app/user/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
