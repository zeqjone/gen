package cmd

import (
	"fmt"
	"gen/repo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export go structor file from mysql",
	Long:  "export go structure file from mysql",
	Run: func(c *cobra.Command, args []string) {
		dsn := viper.GetString("dsn")
		if dsn == "" {
			fmt.Println("请先通过 conf 指令配置 dns")
		}
		repo.NewDB(&repo.MysqlCfg{
			Dsn: dsn,
		})
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
