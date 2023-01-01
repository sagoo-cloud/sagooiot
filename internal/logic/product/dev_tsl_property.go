package product

import (
	"context"
	"encoding/json"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"math"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
)

type sDevTSLProperty struct{}

func init() {
	service.RegisterDevTSLProperty(devTSLPropertyNew())
}

func devTSLPropertyNew() *sDevTSLProperty {
	return &sDevTSLProperty{}
}

func (s *sDevTSLProperty) ListProperty(ctx context.Context, in *model.ListTSLPropertyInput) (out *model.ListTSLPropertyOutput, err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Id, in.ProductId).Scan(&p)
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

	out = new(model.ListTSLPropertyOutput)
	out.CurrentPage = in.PageNum

	if len(tsl.Properties) == 0 {
		return
	}

	if in.Name != "" {
		j := 0
		for _, v := range tsl.Properties {
			if gstr.Contains(v.Name, in.Name) {
				tsl.Properties[j] = v
				j++
			}
		}
		tsl.Properties = tsl.Properties[:j]
	}

	if in.DateType != "" {
		j := 0
		for _, v := range tsl.Properties {
			if gstr.Contains(v.ValueType.Type, in.DateType) {
				tsl.Properties[j] = v
				j++
			}
		}
		tsl.Properties = tsl.Properties[:j]
	}

	length := len(tsl.Properties)
	out.Total = length

	if in.PageNum > int(math.Ceil(float64(length)/float64(in.PageSize))) {
		return
	}
	start := (in.PageNum - 1) * in.PageSize
	end := in.PageSize + start
	if end > length {
		end = length
	}
	out.Data = tsl.Properties[start:end]

	return
}

func (s *sDevTSLProperty) AllProperty(ctx context.Context, key string) (list []model.TSLProperty, err error) {
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
	list = tsl.Properties

	return
}

func (s *sDevTSLProperty) AddProperty(ctx context.Context, in *model.TSLPropertyInput) (err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Id, in.ProductId).Scan(&p)
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

	// 检查属性标识Key是否存在
	existKey := checkExistKey(in.Key, *tsl)
	if existKey {
		return gerror.New("标识已存在，物模型模块下唯一")
	}

	tsl.Properties = append(tsl.Properties, in.TSLProperty)
	metaData, _ := json.Marshal(tsl)

	err = dao.DevProduct.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.DevProduct.Ctx(ctx).
			Data(dao.DevProduct.Columns().Metadata, metaData).
			Where(dao.DevProduct.Columns().Id, in.ProductId).
			Update()
		if err != nil {
			return err
		}

		// 增加TD表字段
		if p.MetadataTable == 1 {
			maxLength := 0
			if in.ValueType.TSLParamBase.MaxLength != nil {
				maxLength = *in.ValueType.TSLParamBase.MaxLength
			}
			err = service.TSLTable().AddDatabaseField(ctx, p.Key, in.Key, in.ValueType.Type, maxLength)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return
}

func (s *sDevTSLProperty) EditProperty(ctx context.Context, in *model.TSLPropertyInput) (err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Id, in.ProductId).Scan(&p)
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

	// 检查属性标识Key是否存在
	existKey := false
	existIndex := 0
	for i, v := range tsl.Properties {
		if v.Key == in.Key {
			existKey = true
			existIndex = i
			break
		}
	}
	if !existKey {
		return gerror.New("属性不存在")
	}

	old := tsl.Properties[existIndex]
	old.Name = in.Name
	old.AccessMode = in.AccessMode
	old.Desc = in.Desc
	in.ValueType.Type = old.ValueType.Type
	old.ValueType = in.ValueType

	newProperties := append(tsl.Properties[:existIndex], old)
	tsl.Properties = append(newProperties, tsl.Properties[existIndex+1:]...)
	metaData, _ := json.Marshal(tsl)

	_, err = dao.DevProduct.Ctx(ctx).
		Data(dao.DevProduct.Columns().Metadata, metaData).
		Where(dao.DevProduct.Columns().Id, in.ProductId).
		Update()

	return
}

func (s *sDevTSLProperty) DelProperty(ctx context.Context, in *model.DelTSLPropertyInput) (err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Id, in.ProductId).Scan(&p)
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

	// 检查属性标识Key是否存在
	existKey := false
	existIndex := 0
	for i, v := range tsl.Properties {
		if v.Key == in.Key {
			existKey = true
			existIndex = i
			break
		}
	}
	if !existKey {
		return gerror.New("属性不存在")
	}

	tsl.Properties = append(tsl.Properties[:existIndex], tsl.Properties[existIndex+1:]...)
	metaData, _ := json.Marshal(tsl)

	err = dao.DevProduct.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.DevProduct.Ctx(ctx).
			Data(dao.DevProduct.Columns().Metadata, metaData).
			Where(dao.DevProduct.Columns().Id, in.ProductId).
			Update()
		if err != nil {
			return err
		}

		// 删除TD表字段
		if p.MetadataTable == 1 {
			err = service.TSLTable().DelDatabaseField(ctx, p.Key, in.Key)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return
}

// 检查标识Key是否存在，物模型模块下唯一
func checkExistKey(key string, tsl model.TSL) bool {
	for _, v := range tsl.Properties {
		if v.Key == key {
			return true
		}
	}
	for _, v := range tsl.Functions {
		if v.Key == key {
			return true
		}
	}
	for _, v := range tsl.Events {
		if v.Key == key {
			return true
		}
	}
	for _, v := range tsl.Tags {
		if v.Key == key {
			return true
		}
	}
	return false
}
