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

type sDevTSLEvent struct{}

func init() {
	service.RegisterDevTSLEvent(devTSLEventNew())
}

func devTSLEventNew() *sDevTSLEvent {
	return &sDevTSLEvent{}
}

func (s *sDevTSLEvent) Detail(ctx context.Context, deviceKey string, eventKey string) (event *model.TSLEvent, err error) {
	dout, err := service.DevDevice().Get(ctx, deviceKey)
	if err != nil {
		return
	}
	if dout.TSL == nil {
		err = gerror.Newf("设备 %s 物模型数据异常", deviceKey)
		return
	}

	for _, v := range dout.TSL.Events {
		if v.Key == eventKey {
			return &v, nil
		}
	}

	return
}

func (s *sDevTSLEvent) ListEvent(ctx context.Context, in *model.ListTSLEventInput) (out *model.ListTSLEventOutput, err error) {
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

	out = new(model.ListTSLEventOutput)
	out.CurrentPage = in.PageNum

	if len(tsl.Events) == 0 {
		return
	}

	length := len(tsl.Events)
	out.Total = length

	if in.PageNum > int(math.Ceil(float64(length)/float64(in.PageSize))) {
		return
	}
	start := (in.PageNum - 1) * in.PageSize
	end := in.PageSize + start
	if end > length {
		end = length
	}
	out.Data = tsl.Events[start:end]

	return
}

func (s *sDevTSLEvent) AllEvent(ctx context.Context, key string) (list []model.TSLEvent, err error) {
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
	list = tsl.Events

	return
}

func (s *sDevTSLEvent) AddEvent(ctx context.Context, in *model.TSLEventAddInput) (err error) {
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

	tsl.Events = append(tsl.Events, in.TSLEvent)
	metaData, _ := json.Marshal(tsl)

	_, err = dao.DevProduct.Ctx(ctx).
		Data(dao.DevProduct.Columns().Metadata, metaData).
		Where(dao.DevProduct.Columns().Key, in.ProductKey).
		Update()

	return
}

func (s *sDevTSLEvent) EditEvent(ctx context.Context, in *model.TSLEventAddInput) (err error) {
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
	for i, v := range tsl.Events {
		if strings.EqualFold(v.Key, in.Key) {
			existKey = true
			existIndex = i
			break
		}
	}
	if !existKey {
		return gerror.New("事件不存在")
	}

	newEvents := append(tsl.Events[:existIndex], in.TSLEvent)
	tsl.Events = append(newEvents, tsl.Events[existIndex+1:]...)
	metaData, _ := json.Marshal(tsl)

	_, err = dao.DevProduct.Ctx(ctx).
		Data(dao.DevProduct.Columns().Metadata, metaData).
		Where(dao.DevProduct.Columns().Key, in.ProductKey).
		Update()

	return
}

func (s *sDevTSLEvent) DelEvent(ctx context.Context, in *model.DelTSLEventInput) (err error) {
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
	for i, v := range tsl.Events {
		if strings.EqualFold(v.Key, in.Key) {
			existKey = true
			existIndex = i
			break
		}
	}
	if !existKey {
		return gerror.New("事件不存在")
	}

	tsl.Events = append(tsl.Events[:existIndex], tsl.Events[existIndex+1:]...)
	metaData, _ := json.Marshal(tsl)

	_, err = dao.DevProduct.Ctx(ctx).
		Data(dao.DevProduct.Columns().Metadata, metaData).
		Where(dao.DevProduct.Columns().Key, in.ProductKey).
		Update()

	return
}
