package cmd

import (
	"fmt"
	"gitee/zeqjone/gen/conf"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cn string
)
var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "gen project config detail",
		Long:  "gen 工具相关的配置信息",
		Run: func(c *cobra.Command, args []string) {
			if cn != "" {
				for _, s := range conf.ConfigKeys {
					if strings.Contains(s, cn) {
						fmt.Printf("配置文件里的配置如下: \n %s: %v\n", s, viper.GetString(s))
						return
					}
					fmt.Printf("配置文件里目前没有模糊匹配到输入的配置项：%s", cn)
				}
			}
			fmt.Sprintln("配置文件里的配置如下:")
			for _, s := range conf.ConfigKeys {
				fmt.Printf("%18s: %v\n", s, viper.GetString(s))
			}
		},
	}
)

func init() {
	configCmd.Flags().StringVarP(&cn, "list", "l", "", "查看当前指定的配置项，若为空，则返回所有的配置项")
	rootCmd.AddCommand(configCmd)
}
