package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "gen project config detail",
		Long:  "gen project config detail",
		Run: func(c *cobra.Command, args []string) {
			fmt.Printf("config list as below: \n %s: %v\n", "dsn", viper.GetString("dsn"))
		},
	}
)

func init() {
	rootCmd.AddCommand(configCmd)
}
