package utdb

import (
	"database/sql"
	"fmt"

	_ "gitee.com/chunanyong/dm" // 达梦驱动
	_ "github.com/go-sql-driver/mysql"
)

// DB 封装数据库连接和配置
type DB struct {
	*sql.DB
	config *Config
}

// Connect 根据配置连接数据库
func Connect(cfg *Config) (*DB, error) {
	driverName := string(cfg.Type)
	dsn := cfg.DSN()
	//"dm://SYSDBA:SYSDBA001@172.20.30.102:5236?schema=yfk_basedata"
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("连接失败: %w", err)
	}

	// 验证连接
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("Ping失败: %w", err)
	}

	return &DB{db, cfg}, nil
}

// Close 关闭连接
func (db *DB) Close() error {
	return db.DB.Close()
}
