package product

import (
	"context"
	"encoding/json"
	"math"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sDevTSLFunction struct{}

func init() {
	service.RegisterDevTSLFunction(devTSLFunctionNew())
}

func devTSLFunctionNew() *sDevTSLFunction {
	return &sDevTSLFunction{}
}

func (s *sDevTSLFunction) ListFunction(ctx context.Context, in *model.ListTSLFunctionInput) (out *model.ListTSLFunctionOutput, err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, in.ProductKey).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return nil, gerror.New("产品不存在")
	}

	j, err := gjson.DecodeToJson(p.Metadata)
	if err != nil {
		return
	}
	tsl := new(model.TSL)
	if err = j.Scan(tsl); err != nil {
		return
	}

	out = new(model.ListTSLFunctionOutput)
	out.CurrentPage = in.PageNum

	if len(tsl.Functions) == 0 {
		return
	}

	length := len(tsl.Functions)
	out.Total = length

	if in.PageNum > int(math.Ceil(float64(length)/float64(in.PageSize))) {
		return
	}
	start := (in.PageNum - 1) * in.PageSize
	end := in.PageSize + start
	if end > length {
		end = length
	}
	out.Data = tsl.Functions[start:end]

	return
}

func (s *sDevTSLFunction) AllFunction(ctx context.Context, key string, inputsValueTypes string) (list []model.TSLFunction, err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, key).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return nil, gerror.New("产品不存在")
	}

	j, err := gjson.DecodeToJson(p.Metadata)
	if err != nil {
		return
	}
	tsl := new(model.TSL)
	if err = j.Scan(tsl); err != nil {
		return
	}
	if inputsValueTypes != "" {
		for _, function := range tsl.Functions {
			num := 0
			flag := false
			for _, input := range function.Inputs {
				if strings.EqualFold(input.ValueType.Type, inputsValueTypes) {
					flag = true
					num++
				}
			}
			if flag && num == 1 {
				list = append(list, function)
			}
		}
	} else {
		list = tsl.Functions
	}
	return
}

func (s *sDevTSLFunction) AddFunction(ctx context.Context, in *model.TSLFunctionAddInput) (err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, in.ProductKey).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("产品不存在")
	}

	tsl := new(model.TSL)
	j, err := gjson.DecodeToJson(p.Metadata)
	if err != nil {
		return
	}
	if err = j.Scan(tsl); err != nil {
		return
	}

	// 检查标识Key是否存在
	existKey := checkExistKey(in.Key, *tsl)
	if existKey {
		return gerror.New("标识已存在，物模型模块下唯一")
	}

	tsl.Functions = append(tsl.Functions, in.TSLFunction)
	metaData, _ := json.Marshal(tsl)

	_, err = dao.DevProduct.Ctx(ctx).
		Data(dao.DevProduct.Columns().Metadata, metaData).
		Where(dao.DevProduct.Columns().Key, in.ProductKey).
		Update()

	return
}

func (s *sDevTSLFunction) EditFunction(ctx context.Context, in *model.TSLFunctionAddInput) (err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, in.ProductKey).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("产品不存在")
	}

	j, err := gjson.DecodeToJson(p.Metadata)
	if err != nil {
		return
	}
	tsl := new(model.TSL)
	if err = j.Scan(tsl); err != nil {
		return
	}

	// 检查标识Key是否存在
	existKey := false
	existIndex := 0
	for i, v := range tsl.Functions {
		if strings.EqualFold(v.Key, in.Key) {
			existKey = true
			existIndex = i
			break
		}
	}
	if !existKey {
		return gerror.New("功能不存在")
	}

	newFunctions := append(tsl.Functions[:existIndex], in.TSLFunction)
	tsl.Functions = append(newFunctions, tsl.Functions[existIndex+1:]...)
	metaData, _ := json.Marshal(tsl)

	_, err = dao.DevProduct.Ctx(ctx).
		Data(dao.DevProduct.Columns().Metadata, metaData).
		Where(dao.DevProduct.Columns().Key, in.ProductKey).
		Update()

	return
}

func (s *sDevTSLFunction) DelFunction(ctx context.Context, in *model.DelTSLFunctionInput) (err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, in.ProductKey).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("产品不存在")
	}

	j, err := gjson.DecodeToJson(p.Metadata)
	if err != nil {
		return
	}
	tsl := new(model.TSL)
	if err = j.Scan(tsl); err != nil {
		return
	}

	// 检查标识Key是否存在
	existKey := false
	existIndex := 0
	for i, v := range tsl.Functions {
		if strings.EqualFold(v.Key, in.Key) {
			existKey = true
			existIndex = i
			break
		}
	}
	if !existKey {
		return gerror.New("功能不存在")
	}

	tsl.Functions = append(tsl.Functions[:existIndex], tsl.Functions[existIndex+1:]...)
	metaData, _ := json.Marshal(tsl)

	_, err = dao.DevProduct.Ctx(ctx).
		Data(dao.DevProduct.Columns().Metadata, metaData).
		Where(dao.DevProduct.Columns().Key, in.ProductKey).
		Update()

	return
}
