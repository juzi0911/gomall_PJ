package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/juzi0911/gomall_PJ/app/frontend/hertz_gen/frontend/common"
	"github.com/juzi0911/gomall_PJ/app/frontend/infra/rpc"
	rpcproduct "github.com/juzi0911/gomall_PJ/rpc_gen/kitex_gen/product"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	p, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Hot Sale",
		"items": p.Products,
	}, nil
}
