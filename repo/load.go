package repo

import (
	"fmt"
	"strconv"
)

type Table struct {
	Name    string
	Comment string
	Cols    []Column
}

type Column struct {
	Name    string // 列名
	Type    string // 类型
	Len     int    // 长度
	Key     string // 索引
	Comment string // 注释
}

// GetAllTables 获取数据库里的所有的表
func GetAllTables(dbname string) []*Table {
	sql := "select table_name, table_comment from information_schema.tables where table_schema = ?"
	rows, err := db.Query(sql, dbname)
	if err != nil {
		panic(fmt.Errorf("GetAllTables err: %v", err))
	}
	var tbls []*Table
	for rows.Next() {
		var tn, tc string
		rows.Scan(&tn, &tc)
		tbls = append(tbls, &Table{Name: tn, Comment: tc})
	}
	return tbls
}

// GetTable 得到表的信息
func GetTable(dbname string, tbl *Table) {
	sql := "select column_name, Data_type, character_maximum_length,column_key,convert(column_comment using utf8) COLLATE utf8_bin from information_schema.columns where table_schema = ? and  table_name = ?"
	rows, err := db.Query(sql, dbname, tbl.Name)
	if err != nil {
		panic(fmt.Errorf("GetAllTables err: %v", err))
	}
	for rows.Next() {
		var name, dataType, len, key string
		var comment []byte
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
			Comment: string(comment),
		}
		tbl.Cols = append(tbl.Cols, *col)
	}
}
