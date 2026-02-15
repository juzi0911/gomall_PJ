package service

import (
	"context"
	product "github.com/juzi0911/gomall_PJ/rpc_gen/kitex_gen/product"
	"github.com/juzi0911/gomall_PJ/app/product/biz/model"
	"github.com/juzi0911/gomall_PJ/app/product/biz/dal/mysql"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	products, err := productQuery.SearchProducts(req.Query)

	var results []*product.Product
	for _, p := range products {
		results = append(results, &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		})
	}

	return &product.SearchProductsResp{
		Results: results,
	}, nil
}
