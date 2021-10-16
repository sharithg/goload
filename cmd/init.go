package cmd

import (
	"errors"
	"fmt"
	"goload/docker"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var name string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new load balancer",
	Long: `Initializes a new load balancer by building the docker images and
			exposing the ports`,
	Args: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")

		if nameFlag == "" {
			return errors.New("requires a name argument")
		}
		// if len(args) < 1 {
		return nil
		// }

	},
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		nameFlag, _ := cmd.Flags().GetString("name")
		docker.BuildDocker(nameFlag)
		fmt.Print(args)
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	initCmd.PersistentFlags().StringVar(&name, "name", "", "name of the project")
	viper.BindPFlag("name", initCmd.PersistentFlags().Lookup("name"))

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(initCmd)

}
