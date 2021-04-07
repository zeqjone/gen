package repo

func GetGoType(s string) string {
	var gs string
	switch s {
	case "int", "tinyint", "smallint":
		gs = "int"
	case "bigint":
		gs = "int64"
	case "longtext", "text", "char", "enum", "mediumtext", "varchar", "nvarchar", "json":
		gs = "string"
	case "datetime", "timestamp", "time", "date":
		gs = "time.Time"
	case "decimal", "float":
		gs = "float32"
	case "double":
		gs = "float64"
	}
	return gs
}
