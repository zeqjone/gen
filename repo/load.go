package repo

import (
	"fmt"
	"strconv"
)

type Table struct {
	Name string
	Cols []Column
}

type Column struct {
	Name    string // 列名
	Type    string // 类型
	Len     int    // 长度
	Key     string // 索引
	Comment string // 注释
}

// GetAllTables 获取数据库里的所有的表
func GetAllTables(dbname string) []string {
	sql := "select table_name from information_schema.tables where table_schema = ?"
	rows, err := db.Query(sql, dbname)
	if err != nil {
		panic(fmt.Errorf("GetAllTables err: %v", err))
	}
	var tbls []string
	for rows.Next() {
		var tn string
		rows.Scan(&tn)
		tbls = append(tbls, tn)
	}
	return tbls
}

func GetTable(dbname, tablename string) *Table {
	sql := "select column_name, Data_type, character_maximum_length,column_key,column_comment from information_schema.columns where table_schema = ? and  table_name = ?"
	rows, err := db.Query(sql, dbname, tablename)
	if err != nil {
		panic(fmt.Errorf("GetAllTables err: %v", err))
	}
	var tbl = &Table{Name: tablename}
	for rows.Next() {
		var name, dataType, len, key, comment string
		rows.Scan(&name, &dataType, &len, &key, &comment)
		var ln int
		if len != "" {
			ln, _ = strconv.Atoi(len)
		}
		var col = &Column{
			Name:    name,
			Type:    dataType,
			Len:     ln,
			Key:     key,
			Comment: comment,
		}
		tbl.Cols = append(tbl.Cols, *col)
	}
	return tbl
}
