package conf

type Conf struct {
	Dsn       string            `json:"dsn"`
	Output    string            `json:"output"`
	NameSpace string            `json:"namespace"`
	Types     map[string]string `json:"types"`
}

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