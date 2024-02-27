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

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sDevTSLTag struct{}

func init() {
	service.RegisterDevTSLTag(devTSLTagNew())
}

func devTSLTagNew() *sDevTSLTag {
	return &sDevTSLTag{}
}

func (s *sDevTSLTag) ListTag(ctx context.Context, in *model.ListTSLTagInput) (out *model.ListTSLTagOutput, err error) {
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

	out = new(model.ListTSLTagOutput)
	out.CurrentPage = in.PageNum

	if len(tsl.Tags) == 0 {
		return
	}

	length := len(tsl.Tags)
	out.Total = length

	if in.PageNum > int(math.Ceil(float64(length)/float64(in.PageSize))) {
		return
	}
	start := (in.PageNum - 1) * in.PageSize
	end := in.PageSize + start
	if end > length {
		end = length
	}
	out.Data = tsl.Tags[start:end]

	return
}

func (s *sDevTSLTag) AddTag(ctx context.Context, in *model.TSLTagInput) (err error) {
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

	tsl.Tags = append(tsl.Tags, in.TSLTag)
	metaData, _ := json.Marshal(tsl)

	err = dao.DevProduct.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.DevProduct.Ctx(ctx).
			Data(dao.DevProduct.Columns().Metadata, metaData).
			Where(dao.DevProduct.Columns().Key, in.ProductKey).
			Update()
		if err != nil {
			return err
		}

		// 增加TD表标签
		if p.MetadataTable == 1 {
			maxLength := 0
			if in.ValueType.TSLParamBase.MaxLength != nil {
				maxLength = *in.ValueType.TSLParamBase.MaxLength
			}
			err = service.TSLTable().AddTag(ctx, p.Key, in.Key, in.ValueType.Type, maxLength)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return
}

func (s *sDevTSLTag) EditTag(ctx context.Context, in *model.TSLTagInput) (err error) {
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
	for i, v := range tsl.Tags {
		if strings.EqualFold(v.Key, in.Key) {
			existKey = true
			existIndex = i
			break
		}
	}
	if !existKey {
		return gerror.New("标签不存在")
	}

	newTags := append(tsl.Tags[:existIndex], in.TSLTag)
	tsl.Tags = append(newTags, tsl.Tags[existIndex+1:]...)
	metaData, _ := json.Marshal(tsl)

	err = dao.DevProduct.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.DevProduct.Ctx(ctx).
			Data(dao.DevProduct.Columns().Metadata, metaData).
			Where(dao.DevProduct.Columns().Key, in.ProductKey).
			Update()
		if err != nil {
			return err
		}

		// 更新TD表结构
		if p.MetadataTable == 1 {
			maxLength := 0
			if in.ValueType.TSLParamBase.MaxLength != nil {
				maxLength = *in.ValueType.TSLParamBase.MaxLength
			}
			err = service.TSLTable().ModifyTag(ctx, p.Key, in.Key, in.ValueType.Type, maxLength)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return
}

func (s *sDevTSLTag) DelTag(ctx context.Context, in *model.DelTSLTagInput) (err error) {
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
	for i, v := range tsl.Tags {
		if strings.EqualFold(v.Key, in.Key) {
			existKey = true
			existIndex = i
			break
		}
	}
	if !existKey {
		return gerror.New("标签不存在")
	}

	tsl.Tags = append(tsl.Tags[:existIndex], tsl.Tags[existIndex+1:]...)
	metaData, _ := json.Marshal(tsl)

	err = dao.DevProduct.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.DevProduct.Ctx(ctx).
			Data(dao.DevProduct.Columns().Metadata, metaData).
			Where(dao.DevProduct.Columns().Key, in.ProductKey).
			Update()
		if err != nil {
			return err
		}

		// 删除TD表字段
		if p.MetadataTable == 1 {
			err = service.TSLTable().DelTag(ctx, p.Key, in.Key)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return
}
