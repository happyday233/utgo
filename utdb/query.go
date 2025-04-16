package utdb

import (
	"context"
	"database/sql"
	"fmt"
)

// Query 通用查询方法
func Query(db *sql.DB, sqlStr string, args ...interface{}) ([]map[string]interface{}, error) {
	// 定义每页的记录数
	limit := 500
	// 初始化页码为 0
	offset := 0
	// 用于存储最终的查询结果
	var allResults []map[string]interface{}

	for {
		// 构建带有分页参数的 SQL 查询语句
		paginatedSQL := fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlStr, limit, offset)
		// 执行查询
		rows, err := db.Query(paginatedSQL)
		if err != nil {
			return nil, err
		}
		// 确保在函数结束时关闭查询结果集
		defer rows.Close()

		// 获取结果集的列名
		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		// 创建用于存储每一行数据的切片
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// 存储当前页的查询结果
		var pageResults []map[string]interface{}
		for rows.Next() {
			// 扫描当前行的数据到 values 切片中
			if err := rows.Scan(valuePtrs...); err != nil {
				return nil, err
			}
			// 创建一个映射用于存储当前行的数据
			row := make(map[string]interface{})
			for i, col := range columns {
				var v interface{}
				val := values[i]
				b, ok := val.([]byte)
				if ok {
					v = string(b)
				} else {
					v = val
				}
				row[col] = v
			}
			// 将当前行的数据添加到当前页的结果中
			pageResults = append(pageResults, row)
		}

		// 如果当前页没有查询到数据，说明已经查询完所有记录，退出循环
		if len(pageResults) == 0 {
			break
		}

		// 将当前页的结果添加到最终结果中
		allResults = append(allResults, pageResults...)
		// 增加偏移量，准备查询下一页
		offset += limit
	}

	return allResults, nil
}

// Exec 执行非查询语句（INSERT/UPDATE/DELETE）
func (db *DB) Exec(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	return db.DB.ExecContext(ctx, sql, args...)
}

// Get 查询单行数据
func (db *DB) Get(ctx context.Context, dest interface{}, sql string, args ...interface{}) error {
	row := db.DB.QueryRowContext(ctx, sql, args...)
	return row.Scan(dest) // 需结合反射处理结构体字段，此处简化
}
