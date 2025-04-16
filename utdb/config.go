package utdb

import (
	"fmt"
	"net/url"
)

// Config 数据库连接配置
type Config struct {
	Type     string
	Hostaddr string
	Username string
	Password string
	Database string
	Params   map[string]string // 额外参数（如超时、编码）
}

// DSN 生成数据库连接字符串
func (c *Config) DSN() string {
	if c.Type == "mysql" {
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
			c.Username, c.Password, c.Hostaddr, c.Database, encodeParams(c.Params))
	} else if c.Type == "dm" {
		//"dm://SYSDBA:SYSDBA001@172.20.30.102:5236?schema=yfk_basedata"
		return fmt.Sprintf("dm://%s:%s@%s?schema=%s",
			c.Username, c.Password, c.Hostaddr, c.Database)
	}
	fmt.Printf("Unknown type: %s", c.Type)
	//c.Username, c.Password, c.Host, c.Port, c.Database, encodeParams(c.Params))
	return ""
}

// encodeParams 将参数转换为URL查询格式
func encodeParams(params map[string]string) string {
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}
	return values.Encode()
}
