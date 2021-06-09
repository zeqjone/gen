package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init gen ",
	Long:  "init gen with project config, such as db,output dir etc",
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("init args: %v\n", viper.GetString("dsn"))
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	viper.SetDefault("dsn", "root:123456@tcp(127.0.0.1:3306)/rdc_manager?parseTime=true&loc=Local")
	viper.SetDefault("tabels", "")
	initCmd.Flags().StringP("dsn", "c", "mysql connection config", "to connect mysql for getting the structrue of db")
	viper.BindPFlag("dsn", initCmd.Flags().Lookup("dsn"))
	rootCmd.AddCommand(initCmd)
}
