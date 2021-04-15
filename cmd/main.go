package main

import (
	"flag"
	"fmt"
	"gen/lib/utils"
	"gen/repo"
	"os"
	"os/exec"
	"strings"

	"github.com/BurntSushi/toml"
)

var (
	conf string
	cfg  = &Cfg{}
)

type GoCfg struct {
	Namespace string
	Path      string
}

type Cfg struct {
	Mysql repo.MysqlCfg
	Gocfg GoCfg
}

func init() {
	flag.StringVar(&conf, "conf", "cmd/default.toml", "配置文件地址")
	flag.Parse()
	if conf == "" {
		panic("conf 参数不能为空")
	}
	_, err := toml.DecodeFile(conf, &cfg)
	if err != nil {
		fmt.Printf("toml decode error %v", err)
	}
	fmt.Printf("\n%v\n", cfg)
}

func main() {
	// 初始化数据库
	repo.NewDB(&cfg.Mysql)
	tbls := repo.GetAllTables(cfg.Mysql.Db)
	for _, t := range tbls {
		repo.GetTable(cfg.Mysql.Db, t)
	}
	SaveGoStruct(tbls)
	fmt.Println("waiting for fmt")
	fmtTableModel(cfg.Gocfg.Path)
	fmt.Println("fmt for finished")
}

func fmtTableModel(f string) {
	cmd := exec.Command("goimports", "-w", f)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("fmt err22: %v\n\n", err)
	}
}

func SaveGoStruct(tbls []*repo.Table) {
	str := &strings.Builder{}
	str.WriteString(fmt.Sprintf("// desc: code generate by tools\n"))
	str.WriteString(fmt.Sprintf("// github: https://github.com/zeqjone/gen.git\n"))
	str.WriteString(fmt.Sprintf("// mail: zeq_jone@163.com\n"))
	str.WriteString(fmt.Sprintf("// version: v1.0.0\n\n"))
	str.WriteString(fmt.Sprintf("package %s\n", cfg.Gocfg.Namespace))
	for _, t := range tbls {
		fmt.Printf("t:%s\n", t.Name)
		if len(cfg.Mysql.Tables) == 0 || utils.Has(cfg.Mysql.Tables, t.Name) {
			pascalName := utils.Snake2Pascal(t.Name)
			str.WriteString(fmt.Sprintf("// %s %s\n", pascalName, t.Comment))
			str.WriteString(fmt.Sprintf("type %s struct {\n", pascalName))
			for _, c := range t.Cols {
				cs := GetColDesp(c)
				str.WriteString(cs)
			}
			str.WriteString("}\n")
			strFunc := GetTableNameFunc(t.Name)
			str.WriteString(strFunc)
		}
	}
	f, err := os.Create(cfg.Gocfg.Path)
	if err != nil {
		panic(fmt.Errorf("创建文件失败：%v", err))
	}
	defer f.Close()
	f.WriteString(str.String())
	fmt.Println("table model file finished")
}

func GetColDesp(col repo.Column) string {
	cs := ""
	if cfg.Mysql.Orm == repo.GenOrmGorm {
		colName := utils.Snake2Pascal(col.Name)
		if colName == "Id" {
			colName = "ID"
		}
		tagGorm := ""
		if strings.ToLower(col.Key) == "pri" || strings.ToLower(col.Key) == "mul" {
			tagGorm = fmt.Sprintf(`gorm:"%scolumn:%s"`, "primaryKey;", col.Name)
		} else if strings.ToLower(col.Key) == "uni" {
			tagGorm = fmt.Sprintf(`gorm:"%scolumn:%s"`, "unique;", col.Name)
		} else {
			tagGorm = fmt.Sprintf(`gorm:"%scolumn:%s"`, "", col.Name)
		}
		jsonName := utils.Camel2Snake(col.Name)
		if col.Comment != "" {
			cs = fmt.Sprintf("%s %s `json:\"%s\" %s` // %s\n", colName, repo.GetGoType(col.Type), jsonName, tagGorm, col.Comment)
		} else {
			cs = fmt.Sprintf("%s %s `json:\"%s\" %s` \n", colName, repo.GetGoType(col.Type), jsonName, tagGorm)
		}
	}
	return cs
}

func GetTableNameFunc(tn string) string {
	return fmt.Sprintf("func (ins *%s) TableName () string {\n return \"%s\"\n}\n", utils.Snake2Pascal(tn), tn)
}
