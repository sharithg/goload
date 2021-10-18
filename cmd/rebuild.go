package cmd

import (
	"fmt"
	"goload/docker"

	"github.com/spf13/cobra"
)

var rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild your project",
	Long:  `Rebuilds your project. Run this if you make any changes and you wish to rerun.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		docker.RebuildDocker()
		// config.GetDockerImageName()
	},
}

func init() {

	rootCmd.AddCommand(rebuildCmd)

}
