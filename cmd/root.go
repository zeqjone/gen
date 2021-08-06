package cmd

import (
	"errors"
	"fmt"
	"gitee/zeqjone/gen/conf"
	"gitee/zeqjone/gen/lib/utils"
	"gitee/zeqjone/gen/repo"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile = ""
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "cfgfile", "", "设置配置文件地址，默认在用户home目录下")
}

// initConfig 工具初始化，并写入配置文件
func initConfig() {
	if cfgFile == "" {
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigName("gen")
		viper.SetConfigType("yaml")

		_, err = os.Open(path.Join(home, "gen"))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				_, err := os.Create(path.Join(home, "gen"))
				if err != nil {
					fmt.Println(err)
				}
			}
			if errors.Is(err, os.ErrPermission) {
				fmt.Printf("当前账户创建 %s 权限不足", path.Join(home, "gen"))
				return
			}
		}
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
		fmt.Println("正在使用的数据库连接信息：", dsn)
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

// fmtTableModel 格式化生成的 go 文件
func fmtTableModel(f string) {
	cmd := exec.Command("goimports", "-w", f)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("fmt err22: %v\n\n", err)
	}
}

// SaveGoStruct 将表接口映射到 go struct
func SaveGoStruct(tbls []*repo.Table) {
	str := &strings.Builder{}
	str.WriteString("// desc: code generate by tools\n")
	str.WriteString("// version: v1.0.0\n\n")
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
		strFunc := GetTableNameFunc(t)
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

// GetColDesp 将数据库的表字段映射到结构体的字段
func GetColDesp(col repo.Column) string {
	cs := ""
	orm := viper.GetString(conf.MysqlOrm)
	fmt.Println(orm, col)
	if repo.GenOrm(orm) == repo.GenOrmGorm {
		colName := utils.Snake2Pascal(col.Name)
		if colName == "Id" {
			colName = "ID"
		}
		tags := ""
		if strings.ToLower(col.Key) == "pri" {
			tags = fmt.Sprintf(`gorm:"%scolumn:%s" column:"%s"`, "primaryKey;", col.Name, col.Name)
		} else if strings.ToLower(col.Key) == "uni" {
			tags = fmt.Sprintf(`gorm:"%scolumn:%s" column:"%s"`, "unique;", col.Name, col.Name)
		} else {
			tags = fmt.Sprintf(`gorm:"%scolumn:%s" column:"%s"`, "", col.Name, col.Name)
		}
		jsonName := utils.Camel2Snake(col.Name)
		cs = fmt.Sprintf("%s %s `json:\"%s\" %s`", colName, repo.GetGoType(col.Type), jsonName, tags)
		if col.Comment != "" {
			cs += fmt.Sprintf(" // %s\n", col.Comment)
		} else {
			cs += "\n"
		}
	}
	return cs
}

// GetTableNameFunc 生成的结构体方法
func GetTableNameFunc(t *repo.Table) string {
	tablename := fmt.Sprintf("func (ins *%s) TableName () string {\n return \"%s\"\n}", utils.Snake2Pascal(t.Name), t.Name)
	getTableName := fmt.Sprintf("func(ins *%s) GetTableName()string{\n return \"%s\"\n}", utils.Snake2Pascal(t.Name), strings.ToUpper(t.Name))
	pks := make([]string, 0)
	for _, c := range t.Pks {
		pks = append(pks, c.Name)
	}
	pk := strings.Join(pks, ",")
	if pk == "" {
		pk = "id"
	}
	// todo: 如果是组合主键，改怎么返回？
	getPKColumnName := fmt.Sprintf("func(ins *%s) GetPKColumnName()string{\n return \"%s\"\n}", utils.Snake2Pascal(t.Name), pk)
	getPkSequence := fmt.Sprintf("func(ins *%s) GetPkSequence()map[string]string{\n return nil\n}", utils.Snake2Pascal(t.Name))
	return fmt.Sprintf("%s\n%s\n%s\n%s\n", tablename, getTableName, getPKColumnName, getPkSequence)
}
