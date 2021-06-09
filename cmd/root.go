package cmd

import (
	"fmt"
	"gen/conf"
	"gen/lib/utils"
	"gen/repo"
	"os"
	"os/exec"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile = ""
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "cfgfile", "", "设置配置文件地址，默认在用户home目录下")
}
func initConfig() {
	if cfgFile == "" {
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigName("cobra")
		viper.SetConfigType("yaml")
	} else {
		// use config file from flags
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
		f, err := os.Open(viper.ConfigFileUsed())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = viper.ReadConfig(f)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		dsn := viper.GetString(conf.MysqlDsn)
		if dsn == "" {
			fmt.Println("第一次使用请先指定 gen init 配置数据库连接字符串， 否则不能使用")
			return
		}
		repo.NewDB(&repo.MysqlCfg{
			Dsn: dsn,
		})
	} else {
		fmt.Println(err.Error())
	}
}

var rootCmd = &cobra.Command{
	Use:   "gen",
	Short: "to gen to struct file from mysql",
	Long:  "to gen go struct file from mysql",
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("args: %v", args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
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
	str.WriteString(fmt.Sprintf("package %s\n", viper.GetString(conf.OutputNameSpace)))
	for _, t := range tbls {
		fmt.Printf("t:%s\n", t.Name)
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
	d := viper.GetString(conf.OutputDir)
	fname := viper.GetString(conf.OutputNameSpace)
	f, err := os.Create(fmt.Sprintf("%s/%s.go", d, fname))
	if err != nil {
		panic(fmt.Errorf("创建文件失败：%v", err))
	}
	defer f.Close()
	f.WriteString(str.String())
	fmt.Println("table model file finished")
}

func GetColDesp(col repo.Column) string {
	cs := ""
	orm := viper.GetString(conf.MysqlOrm)
	fmt.Println(orm, col)
	if repo.GenOrm(orm) == repo.GenOrmGorm {
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
