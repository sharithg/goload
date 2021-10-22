package docker

import (
	"bytes"
	"fmt"
	"goload/config"
	"os"
	"os/exec"
	"sync"
)

// runs a individual docker container given an id
func RunDocker(port int, mapPort int) {

	goloadConfig := config.LoadConfig()

	imageId := goloadConfig.ImageId

	dockerCmd := fmt.Sprintf("docker run -p %d:%d -d --name %s %s", port, mapPort, fmt.Sprintf("%s-%d", imageId, port), imageId)

	dockerExec := exec.Command("/bin/sh", "-c", dockerCmd)

	var outb, errb bytes.Buffer
	// dockerExec.Stdout = os.Stdout
	dockerExec.Stderr = os.Stderr

	err := dockerExec.Run()

	if err != nil {
		fmt.Println("out:", outb.String(), "err:", errb.String())
		panic(fmt.Sprintf("%s", err))
	}
}

// runs multiple docker containers given the number of ports
func RunMultipleDocker(numberOfPorts int) []string {
	var wg sync.WaitGroup
	startPort := 3010
	numPorts := numberOfPorts
	output := []string{}

	goloadConfig := config.LoadConfig()

	// add the number of concurrent processes to a wait group
	wg.Add(numPorts)
	for i := startPort; i <= startPort+numPorts-1; i++ {
		// run each docker run command concurrently
		go func(portNum int) {
			RunDocker(portNum, goloadConfig.ExposedPort)
			// append port number to a output array
			output = append(output, fmt.Sprint(portNum))
			wg.Done()
		}(i)
	}
	// wait until all concurrent processes are done running
	wg.Wait()
	return output
}
