package cmd

import (
	"errors"
	"fmt"
	"goload/config"
	"goload/docker"
	"goload/globals"
	"goload/loadbalancer"
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
			return errors.New("requires a replicas argument")
		}

		_, err := strconv.Atoi(replicasFlag)
		if err != nil {
			return errors.New("replicasFlag must be a integer")
		}

		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		replicasFlag, _ := cmd.Flags().GetString("replicas")
		serverList := []string{}
		goloadConfig := config.LoadConfig()
		numberOfReplicas, _ := strconv.Atoi(replicasFlag)
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
