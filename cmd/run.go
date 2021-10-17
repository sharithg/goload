package cmd

import (
	"errors"
	"fmt"
	"goload/docker"
	"goload/loadbalancer"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var replicas string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run a initialized project",
	Long:  `Runs the existing project witht he given parameters`,
	Args: func(cmd *cobra.Command, args []string) error {
		replicasFlag, _ := cmd.Flags().GetString("replicas")

		if replicasFlag == "" {
			return errors.New("requires a replicas argument")
		}

		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		serverList := docker.RunMultipleDocker()
		loadbalancer.RunBackend(serverList)
		// config.GetDockerImageName()
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	runCmd.PersistentFlags().StringVar(&replicas, "replicas", "", "number of replicas to balance with")
	viper.BindPFlag("replicas", runCmd.PersistentFlags().Lookup("replicas"))

	rootCmd.AddCommand(runCmd)

}
