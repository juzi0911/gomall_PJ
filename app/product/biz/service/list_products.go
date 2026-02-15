package service

import (
	"context"
	product "github.com/juzi0911/gomall_PJ/rpc_gen/kitex_gen/product"
	"github.com/juzi0911/gomall_PJ/app/product/biz/model"
	"github.com/juzi0911/gomall_PJ/app/product/biz/dal/mysql"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)
	resp = &product.ListProductsResp{}
	for _, category := range c {
		for _, p := range category.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(p.ID),
				Name:        p.Name,
				Description: p.Description,
				Picture:     p.Picture,
				Price:       p.Price,
			})
		}
	}

	return resp, nil
}
