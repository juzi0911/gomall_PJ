package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	category "github.com/juzi0911/gomall_PJ/app/frontend/hertz_gen/frontend/category"
	"github.com/juzi0911/gomall_PJ/app/frontend/infra/rpc"
	rpcproduct "github.com/juzi0911/gomall_PJ/rpc_gen/kitex_gen/product"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Category",
		"items": p.Products,
	}, nil
}
