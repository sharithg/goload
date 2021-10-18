package main

import (
	"fmt"
	"goload/cmd"
	"goload/docker"
	"goload/globals"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func cleanup() {
	fmt.Println("cleanup")

	var wg sync.WaitGroup

	runningIdsLength := len(globals.RUNNING_IDS)

	// create a waitgroup for the number of running containers
	wg.Add(runningIdsLength)

	for _, element := range globals.RUNNING_IDS {

		// concurrently stop the containers
		go func(elem string) {
			docker.StopDocker(elem)
			wg.Done()
		}(element)
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
