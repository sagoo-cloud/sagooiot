package core

import (
	"context"
	"sagooiot/internal/service"
	"sagooiot/network/core/mapper"
	"sagooiot/network/model"
	"sync"
)

var products sync.Map

func LoadProduct(ctx context.Context, key string) (*model.Product, error) {
	v, ok := products.Load(key)
	if ok {
		return v.(*model.Product), nil
	}

	//加载产品
	var product model.Product
	productDetail, err := service.DevProduct().Detail(ctx, key)
	if err != nil {
		return nil, err
	}
	product = mapper.Product(*productDetail)

	products.Store(key, &product)

	return &product, nil
}
