package cmd

import (
	"errors"
	"fmt"
	"goload/config"
	"goload/docker"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var name string
var projectDir string

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

		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		nameFlag, _ := cmd.Flags().GetString("name")
		projectDirFlag, _ := cmd.Flags().GetString("dir")

		if projectDirFlag != "" {
			exists, err := PathExits(projectDirFlag)

			if err != nil || !exists {
				errMsg := fmt.Sprintf("%s: no such directory", projectDirFlag)
				fmt.Fprintln(os.Stderr, errMsg)
				os.Exit(1)
			}
		}

		if config.DoesAttributeAndFileExist("imageId") {
			fmt.Fprintln(os.Stderr, "project already exists")
			os.Exit(1)
		}
		imageId := docker.BuildDocker(nameFlag)
		config.WriteDockerImageName(imageId)
		// config.GetDockerImageName()
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	initCmd.PersistentFlags().StringVar(&name, "name", "", "name of the project")
	viper.BindPFlag("name", initCmd.PersistentFlags().Lookup("name"))

	initCmd.PersistentFlags().StringVar(&projectDir, "dir", "", "project directory, will default to current directory")
	viper.BindPFlag("dir", initCmd.PersistentFlags().Lookup("dir"))
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(initCmd)

}
