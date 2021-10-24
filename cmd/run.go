package cmd

import (
	"fmt"
	"goload/config"
	"goload/docker"
	"goload/globals"
	"goload/loadbalancer"
	"goload/utils"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var replicas string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a initialized project",
	Long:  `Runs the existing project with the given parameters`,
	Args: func(cmd *cobra.Command, args []string) error {
		replicasFlag, _ := cmd.Flags().GetString("replicas")

		if replicasFlag == "" {
			utils.FatalError("requires a replicas argument")
		}

		num, err := strconv.Atoi(replicasFlag)
		if err != nil || num < 1 || num > 250 {
			utils.FatalError("replicas must be a number between 1 and 250")
		}

		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		replicasFlag, _ := cmd.Flags().GetString("replicas")
		serverList := []string{}
		goloadConfig := config.LoadConfig()

		numberOfReplicas, err := strconv.Atoi(replicasFlag)

		if err != nil {
			utils.FatalError("replicas must be a integer below 250")
		}

		// run numberOfReplicas docker containers
		portList := docker.RunMultipleDocker(numberOfReplicas)

		for _, port := range portList {
			serverList = append(serverList, fmt.Sprintf("http://localhost:%s", port))
			// append the image ids to the RUNNING_IDS global, which will be use in the cleanup
			globals.RUNNING_IDS = append(globals.RUNNING_IDS, fmt.Sprintf("%s-%s", goloadConfig.ImageId, port))
		}
		loadbalancer.RunBackend(serverList)
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	runCmd.PersistentFlags().StringVarP(&replicas, "replicas", "r", "", "number of replicas to balance with")
	viper.BindPFlag("replicas", runCmd.PersistentFlags().Lookup("replicas"))

	rootCmd.AddCommand(runCmd)

}
