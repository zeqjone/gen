package cmd

import (
	"fmt"
	"gitee/zeqjone/gen/conf"
	"gitee/zeqjone/gen/repo"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	db     string
	tables string
)

// 导出指令
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export go structor file from mysql",
	Long:  "export go structure file from mysql",
	Run: func(c *cobra.Command, args []string) {
		if repo.Ins() == nil {
			fmt.Println("请先执行 init 完成必须的配置")
			return
		}
		tbls := viper.GetStringSlice(conf.MysqlTables)
		if tables != "" {
			tbls = strings.Split(tables, ",")
		}
		var connstr = viper.GetString(conf.MysqlDsn)
		if connstr != "" {
			c, err := mysql.ParseDSN(connstr)
			if err != nil {
				fmt.Printf("dsn error: %v", err)
			}
			db = c.DBName
		}
		allTbls := repo.GetAllTables(db)
		var inTbls []*repo.Table
		if len(tbls) > 0 {
			for _, t := range allTbls {
				for _, it := range tbls {
					if t.Name == it {
						inTbls = append(inTbls, t)
					}
				}
			}
		} else {
			inTbls = allTbls
		}
		for _, t := range inTbls {
			repo.GetTable(db, t)
		}
		SaveGoStruct(inTbls)
		fmt.Println("waiting for fmt")
		fmtTableModel(viper.GetString(conf.OutputDir))
		fmt.Println("fmt for finished")
	},
}

// init
func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringVarP(&tables, "table", "t", "", "指定表导出")
	exportCmd.Flags().StringVarP(&db, "dbname", "n", "", "指定数据库导出")

}
