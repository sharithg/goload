package docker

import (
	"bytes"
	"fmt"
	"goload/config"
	"os"
	"os/exec"
	"sync"
)

func RunDocker(port int) {

	imageId := config.GetDockerImageName()

	dockerCmd := fmt.Sprintf("docker run -p %d:6000 -d --name %s %s", port, fmt.Sprintf("%s-%d", imageId, port), imageId)

	dockerExec := exec.Command("/bin/sh", "-c", dockerCmd)

	var outb, errb bytes.Buffer
	dockerExec.Stdout = os.Stdout
	dockerExec.Stderr = os.Stderr

	err := dockerExec.Run()

	if err != nil {
		fmt.Println("out:", outb.String(), "err:", errb.String())
		panic(fmt.Sprintf("%s", err))
	}
}

func RunMultipleDocker(numberOfPorts int) []string {
	var wg sync.WaitGroup
	startPort := 3010
	numPorts := numberOfPorts
	output := []string{}
	wg.Add(numPorts)
	for i := startPort; i <= startPort+numPorts-1; i++ {
		go func(portNum int) {
			RunDocker(portNum)
			output = append(output, fmt.Sprint(portNum))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(output)
	return output
}
