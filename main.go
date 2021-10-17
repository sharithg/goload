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

	wg.Add(runningIdsLength)

	for _, element := range globals.RUNNING_IDS {
		// index is the index where we are
		// element is the element from someSlice for where we are
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

	go func() {
		<-c
		// run the cleanup when reciving the channel
		cleanup()
		os.Exit(1)
	}()
	cmd.Execute()
}
