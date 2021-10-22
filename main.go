package main

import (
	"goload/cmd"
	"goload/docker"
	"goload/globals"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/fatih/color"
)

func cleanup() {
	color.Green("\nGracefully stopping...")

	var wg sync.WaitGroup

	runningIdsLength := len(globals.RUNNING_IDS)

	// create a waitgroup for the number of running containers
	wg.Add(runningIdsLength)

	for _, id := range globals.RUNNING_IDS {

		// concurrently stop the containers
		go func(id string) {
			docker.StopDocker(id)
			wg.Done()
		}(id)
	}
	wg.Wait()
}

func main() {
	// create a go channel to handle the SIGTERM signal (^C)
	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// attach a go channel to listen for SIGTERM
	go func() {
		<-c
		// run the cleanup when reciving the channel
		cleanup()
		os.Exit(1)
	}()
	cmd.Execute()
}
