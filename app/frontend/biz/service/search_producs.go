package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	product "github.com/juzi0911/gomall_PJ/app/frontend/hertz_gen/frontend/product"
	"github.com/juzi0911/gomall_PJ/app/frontend/infra/rpc"
	rpcproduct "github.com/juzi0911/gomall_PJ/rpc_gen/kitex_gen/product"
)

type SearchProducsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProducsService(Context context.Context, RequestContext *app.RequestContext) *SearchProducsService {
	return &SearchProducsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProducsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{Query: req.Q})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"items": p.Results,
		"q":    req.Q,
	}, nil
}
