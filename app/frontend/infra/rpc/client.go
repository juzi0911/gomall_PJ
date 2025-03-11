package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/juzi0911/gomall_PJ/app/frontend/conf"
	frontendUtils "github.com/juzi0911/gomall_PJ/app/frontend/utils"
	"github.com/juzi0911/gomall_PJ/rpc_gen/kitex_gen/user/userservice"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	UserClient userservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		iniUserClient()
	})
}

func iniUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
