package cmd

import (
	"fmt"
	"gitee/zeqjone/gen/conf"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 配置工具指令
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init gen ",
	Long:  "init gen with project config, such as db,output dir etc",
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("init args: %v\n", viper.GetString("dsn"))
		// 将配置写入配置文件
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
		}
	},
}

// init 定义 init 指令
func init() {
	initCmd.Flags().StringP(conf.MysqlDsn, "c", "root:123456@tcp(127.0.0.1:3306)/test", "to connect mysql for getting the structrue of db")
	viper.BindPFlag(conf.MysqlDsn, initCmd.Flags().Lookup(conf.MysqlDsn))
	initCmd.Flags().StringP(conf.MysqlTables, "t", "", "to connect mysql for getting the structrue of db")
	viper.BindPFlag(conf.MysqlTables, initCmd.Flags().Lookup(conf.MysqlTables))

	initCmd.Flags().StringP(conf.OutputDir, "d", "table", "指定 go structure 存放路径")
	viper.BindPFlag(conf.OutputDir, initCmd.Flags().Lookup(conf.OutputDir))
	initCmd.Flags().StringP(conf.OutputNameSpace, "n", "table", "指定 go structure 文件的命名空间")
	viper.BindPFlag(conf.OutputNameSpace, initCmd.Flags().Lookup(conf.OutputNameSpace))

	initCmd.Flags().StringP(conf.MysqlOrm, "m", "gorm", "指定驱动mysql 的orm，支持 gorm、 beego")
	viper.BindPFlag(conf.MysqlOrm, initCmd.Flags().Lookup(conf.MysqlOrm))

	rootCmd.AddCommand(initCmd)
}
