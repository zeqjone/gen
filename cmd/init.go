package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dsn string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init gen ",
	Long:  "init gen with project config, such as db,output dir etc",
	Run: func(c *cobra.Command, args []string) {

	},
}

func init() {
	viper.SetDefault("dsn", "root:123456@tcp(127.0.0.1:3306)/rdc_manager?parseTime=true&loc=Local")
	initCmd.Flags().StringVarP(&dsn, "dsn", "c", "mysql connection config", "to connect mysql for getting the structrue of db")
	rootCmd.AddCommand(initCmd)
}
