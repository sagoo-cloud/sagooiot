package core

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/network/model"
	"strconv"
	"sync"
)

var products sync.Map

func LoadProduct(ctx context.Context, id string) (*model.Product, error) {
	v, ok := products.Load(id)
	if ok {
		return v.(*model.Product), nil
	}

	//加载产品
	var product model.Product
	productId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	productDetail, err := service.DevProduct().Detail(ctx, uint(productId))
	if err != nil {
		return nil, err
	}
	product = mapperProduct(*productDetail)

	products.Store(id, &product)

	return &product, nil
}
