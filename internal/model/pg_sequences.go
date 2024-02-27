package model

type PgSequenceOut struct {
	SchemaName    string `json:"schemaName" description:"模式名称"`
	SeqUesCeName  string `json:"seqUesCeName" description:"序号名称"`
	SeqUesCeOwner int64  `json:"seqUesCeOwner" description:"所有者"`
	DataType      int64  `json:"dataType" description:"数据类型"`
	StartVale     int64  `json:"startVale" description:"开始值"`
	MaxValue      int64  `json:"maxValue" description:"最大值"`
	MinValue      int64  `json:"minValue" description:"最小值"`
	IncrementBy   int64  `json:"incrementBy" description:"自增量"`
	Cycle         int64  `json:"cycle" description:"是否启用循环"`
	CacheSize     int64  `json:"cacheSize" description:"缓存大小"`
	LastVale      int64  `json:"lastVale" description:"当前值"`
}
