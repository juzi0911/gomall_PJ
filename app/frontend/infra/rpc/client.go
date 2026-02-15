package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/juzi0911/gomall_PJ/app/frontend/conf"
	frontendUtils "github.com/juzi0911/gomall_PJ/app/frontend/utils"
	"github.com/juzi0911/gomall_PJ/rpc_gen/kitex_gen/user/userservice"
	"github.com/juzi0911/gomall_PJ/rpc_gen/kitex_gen/product/productcatalogservice"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	UserClient userservice.Client
	ProductClient productcatalogservice.Client
	once sync.Once
)

func Init() {
	once.Do(func() {
		iniUserClient()
		iniProductClient()
	})
}

func iniUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func iniProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}