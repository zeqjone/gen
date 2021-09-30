package repo

// GetGoType 将数据库的字段类型映射到 go 的类型
func GetGoType(s string) string {
	var gs string
	switch s {
	case "int", "tinyint", "smallint", "mediumint":
		gs = "int"
	case "bigint":
		gs = "int64"
	case "longtext", "text", "char", "enum", "mediumtext", "varchar", "nvarchar", "json":
		gs = "string"
	case "blob", "longblob", "tinyblob", "varbinary", "binary", "mediumblob":
		gs = "[]byte"
	case "datetime", "timestamp", "time", "date":
		gs = "time.Time"
	case "decimal", "float", "real":
		gs = "float32"
	case "double":
		gs = "float64"
	}
	return gs
}
