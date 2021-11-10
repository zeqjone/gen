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
	Long:  "初始化 gen 配置文件，可以配置数据库连接DSN(data source name)，映射输出的go文件的地址，指定文件的命名空间等等",
	Run: func(c *cobra.Command, args []string) {
		// 将配置写入配置文件
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
		}
	},
}

// init 定义 init 指令
func init() {
	initCmd.Flags().StringP(conf.MysqlDsn, "c", "root:123456@tcp(127.0.0.1:3306)/test", "数据库连接字符串，目前仅支持mysql")
	viper.BindPFlag(conf.MysqlDsn, initCmd.Flags().Lookup(conf.MysqlDsn))
	initCmd.Flags().StringP(conf.MysqlTables, "t", "", "配置导出结构体的数据库表信息，若为空，则导出数据库里所有的表；`,`多个表名之间的分隔符")
	viper.BindPFlag(conf.MysqlTables, initCmd.Flags().Lookup(conf.MysqlTables))

	initCmd.Flags().StringP(conf.OutputDir, "d", "table", "导出的结构体存放的位置，一般放在 table 目录下")
	viper.BindPFlag(conf.OutputDir, initCmd.Flags().Lookup(conf.OutputDir))
	initCmd.Flags().StringP(conf.OutputNameSpace, "n", "table", "指定 go structure 文件的命名空间")
	viper.BindPFlag(conf.OutputNameSpace, initCmd.Flags().Lookup(conf.OutputNameSpace))
	initCmd.Flags().StringP(conf.TableNameWithSchema, "s", "true", "返回数据库 table 名字的时候，是否携带返回数据库的名字")
	viper.BindPFlag(conf.TableNameWithSchema, initCmd.Flags().Lookup(conf.TableNameWithSchema))

	initCmd.Flags().StringP(conf.MysqlOrm, "m", "gorm", "指定驱动mysql 的orm，目前结构体支持 gorm，达梦。结构体上的方法仅支持gorm")
	viper.BindPFlag(conf.MysqlOrm, initCmd.Flags().Lookup(conf.MysqlOrm))

	rootCmd.AddCommand(initCmd)
}
