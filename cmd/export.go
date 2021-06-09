package cmd

import (
	"fmt"
	"gen/conf"
	"gen/repo"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	db     string
	tables string
)
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
		fmt.Println(tbls)
		if tables != "" {
			tbls = strings.Split(tables, ",")
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

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringVarP(&tables, "table", "t", "", "指定表导出")
	exportCmd.Flags().StringVarP(&db, "dbname", "n", "", "指定数据库导出")
}