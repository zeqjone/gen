package conf

var (
	MysqlDsn    string = "mysql.dsn"
	MysqlTables string = "mysql.tables"
	MysqlOrm    string = "mysql.orm"
	DmDbName    string = "dm.dbname"

	OutputDir           string = "output.dir"
	OutputNameSpace     string = "output.namespace"
	TableNameWithSchema string = "output.tableNameWithSchema"
)

var ConfigKeys = []string{
	MysqlDsn,
	MysqlTables,
	MysqlOrm,
	OutputDir,
	OutputNameSpace,
	DmDbName,
	TableNameWithSchema,
}
