package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goload",
	Short: "Creates an instant load balancer",
	Long:  `Create a load balancer with minimal config, just initialize a project and run your project.`,
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) < 1 {
	// 		return errors.New("requires a color argument")
	// 	}

	// 	return fmt.Errorf("invalid color specified: %s", args[0])
	// },
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Print(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
