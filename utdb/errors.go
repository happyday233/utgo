package utdb

import "errors"

var (
	ErrInvalidConfig = errors.New("无效的数据库配置")
	ErrPingFailed    = errors.New("数据库连接测试失败")
)
