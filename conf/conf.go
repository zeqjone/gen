package conf

var (
	MysqlDsn        string = "mysql.dsn"
	MysqlTables     string = "mysql.tables"
	MysqlOrm        string = "mysql.orm"
	OutputDir       string = "output.dir"
	OutputNameSpace string = "output.namespace"
)

var ConfigKeys = []string{
	MysqlDsn,
	MysqlTables,
	MysqlOrm,
	OutputDir,
	OutputNameSpace,
}
