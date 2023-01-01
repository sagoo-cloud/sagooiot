package datahub

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"strconv"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

type sDataTemplateRecord struct{}

func init() {
	service.RegisterDataTemplateRecord(dataTemplateRecordNew())
}

func dataTemplateRecordNew() *sDataTemplateRecord {
	return &sDataTemplateRecord{}
}

// 更新数据记录，定时任务触发
func (s *sDataTemplateRecord) UpdateData(ctx context.Context, tid uint64) error {
	dt, err := service.DataTemplate().Detail(ctx, tid)
	if err != nil {
		return err
	}
	if dt == nil {
		return gerror.New("数据模型不存在")
	}

	// 获取节点
	tplNodes, err := service.DataTemplateNode().List(ctx, tid)
	if err != nil {
		return err
	}
	if len(tplNodes) == 0 {
		return gerror.New("数据模型未创建模型节点")
	}

	// 合并数据
	var insertData []*gmap.AnyAnyMap
	if dt.MainSourceId > 0 && dt.SourceNodeKey != "" {
		if insertData, err = s.mergeForRelation(ctx, dt, tplNodes); err != nil {
			return err
		}
	} else {
		if insertData, err = s.merge(ctx, dt, tplNodes); err != nil {
			return err
		}
	}

	// 入库
	if len(insertData) > 0 {
		table := getTplTableName(tid)
		if err = g.DB(DataCenter()).GetCore().ClearTableFields(ctx, table); err != nil {
			return err
		}
		for i := 0; i < len(insertData); i = i + 1000 {
			s := 1000 + i
			if len(insertData)-i < 1000 {
				s = len(insertData)
			}
			d := insertData[i:s]
			if _, err = g.DB(DataCenter()).Save(ctx, table, d); err != nil {
				return err
			}
		}
	}

	return nil
}

// 按关联字段合并
func (s *sDataTemplateRecord) mergeForRelation(ctx context.Context, dt *model.DataTemplateOutput, nodes []*model.DataTemplateNodeOutput) (data []*gmap.AnyAnyMap, err error) {
	// 统计数据源
	sourceIds := make(map[uint64]struct{})
	for _, node := range nodes {
		if node.SourceId > 0 {
			sourceIds[node.SourceId] = struct{}{}
		}
	}
	if len(sourceIds) == 0 {
		err = gerror.New("数据模型未关联数据源节点")
		return
	}

	// 获取数据源数据
	var mainSource map[string]*gmap.StrAnyMap
	recordData := make(map[uint64]map[string]*gmap.StrAnyMap)
	for sid := range sourceIds {
		rs, err := s.getSourceRecordForRelation(ctx, sid, dt, nodes)
		if err != nil {
			return nil, err
		}
		recordData[sid] = rs

		if sid == dt.MainSourceId || len(sourceIds) == 1 {
			mainSource = rs
		}
	}

	// 合并数据
	for rKey, row := range mainSource {
		m := gmap.New()
		for _, node := range nodes {
			if node.From == 2 {
				// 数据源关联节点
				if node.SourceId == dt.MainSourceId {
					m.Set(node.Key, row.Get(node.Key))
				} else {
					if row, ok := recordData[node.SourceId]; ok {
						if rowD, ok := row[rKey]; ok {
							m.Set(node.Key, rowD.Get(node.Key))
						} else {
							m.Set(node.Key, "")
						}
					}
				}
			} else {
				// 自定义节点
				if node.Default != "" {
					m.Set(node.Key, node.Default)
				} else {
					m.Set(node.Key, guid.S())
				}
			}
		}
		data = append(data, m)
	}
	return
}

func (s *sDataTemplateRecord) getSourceRecordForRelation(ctx context.Context, sourceId uint64, dt *model.DataTemplateOutput, nodes []*model.DataTemplateNodeOutput) (rs map[string]*gmap.StrAnyMap, err error) {
	data, err := service.DataSourceRecord().GetForTpl(ctx, sourceId, dt.Id)
	if err != nil {
		return
	}

	rs = make(map[string]*gmap.StrAnyMap, data.Len())
	for i, row := range data {
		m := gmap.NewStrAnyMap()
		for _, node := range nodes {
			if node.From == 2 && node.SourceId == sourceId {
				m.Set(node.Key, row[node.Node.Key])
			}
		}

		relationKey := row[dt.SourceNodeKey].String()
		if relationKey == "" {
			relationKey = strconv.Itoa(i)
		}
		rs[relationKey] = m
	}

	return
}

// 按索引行合并
func (s *sDataTemplateRecord) merge(ctx context.Context, dt *model.DataTemplateOutput, nodes []*model.DataTemplateNodeOutput) (data []*gmap.AnyAnyMap, err error) {
	// 统计数据源
	sourceIds := make(map[uint64]struct{})
	for _, node := range nodes {
		if node.SourceId > 0 {
			sourceIds[node.SourceId] = struct{}{}
		}
	}
	if len(sourceIds) == 0 {
		err = gerror.New("数据模型未关联数据源节点")
		return
	}

	// 聚合数据源数据
	recordNum := 0
	recordData := make(map[uint64][]*gmap.StrAnyMap)
	for sid := range sourceIds {
		rs, err := s.getSourceRecord(ctx, sid, dt.Id, nodes)
		if err != nil {
			return nil, err
		}
		if recordNum == 0 {
			recordNum = len(rs)
		}
		recordData[sid] = rs
	}

	// 合并数据
	for i := 0; i < recordNum; i++ {
		m := gmap.New()
		for _, node := range nodes {
			if node.From == 2 {
				// 数据源关联节点
				if row, ok := recordData[node.SourceId]; ok {
					m.Set(node.Key, row[i].Get(node.Key))
				}
			} else {
				// 自定义节点
				if node.Default != "" {
					m.Set(node.Key, node.Default)
				} else {
					m.Set(node.Key, guid.S())
				}
			}
		}
		data = append(data, m)
	}
	return
}

func (s *sDataTemplateRecord) getSourceRecord(ctx context.Context, sourceId uint64, tid uint64, nodes []*model.DataTemplateNodeOutput) (rs []*gmap.StrAnyMap, err error) {
	data, err := service.DataSourceRecord().GetForTpl(ctx, sourceId, tid)
	if err != nil {
		return
	}

	for _, row := range data {
		m := gmap.NewStrAnyMap()
		for _, node := range nodes {
			if node.From == 2 && node.SourceId == sourceId {
				m.Set(node.Key, row[node.Node.Key])
			}
		}
		rs = append(rs, m)
	}

	return
}
