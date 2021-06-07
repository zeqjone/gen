package conf

type Conf struct {
	Dsn       string            `json:"dsn"`
	Output    string            `json:"output"`
	NameSpace string            `json:"namespace"`
	Types     map[string]string `json:"types"`
}
