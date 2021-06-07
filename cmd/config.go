package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "gen project config detail",
	Long:  "gen project config detail",
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("config list as below: \n %s: %v\n", "dsn", viper.GetString("dsn"))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	viper.SetConfigFile(".gen.yaml")
	// 在系统的配置目录下找配置文件
	viper.AddConfigPath("/etc/gen")
	// 在系统的个人用户目录下找配置文件
	viper.AddConfigPath("$HOME")
	// 在当前目录下找配置文件
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			_, err := os.Create("gen.yaml")
			if err != nil {
				panic(fmt.Errorf("创建配置文件失败: %s \n", err))
			}
		} else {
			panic(fmt.Errorf("fatal: viper readinpath error: %s", err.Error()))
		}
	}
}
