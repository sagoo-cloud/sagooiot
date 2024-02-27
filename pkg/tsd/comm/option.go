package comm

// Option 包含数据库连接的配置选项
type Option struct {
	Host         string
	Port         int
	Link         string
	Org          string
	Token        string
	Username     string
	Password     string
	Database     string
	DriverName   string // 驱动名称
	MaxIdleConns int    // 最大空闲连接数
	MaxOpenConns int    // 最大连接数
}
