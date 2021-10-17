package docker

import (
	"bytes"
	"fmt"
	"goload/config"
	"os"
	"os/exec"
)

func RunDocker(port int) {

	// defer wg.Done()

	imageId := config.GetDockerImageName()

	dockerCmd := fmt.Sprintf("docker run -p %d:6000 -d %s", port, imageId)

	fmt.Println(dockerCmd)

	dockerExec := exec.Command("/bin/sh", "-c", dockerCmd)

	var outb, errb bytes.Buffer
	dockerExec.Stdout = os.Stdout
	dockerExec.Stderr = os.Stderr

	// if there is an error with our execution
	// handle it here
	err := dockerExec.Run()

	fmt.Println("Ran this")

	if err != nil {
		fmt.Println("out:", outb.String(), "err:", errb.String())
		panic(fmt.Sprintf("%s", err))
	}
}

func RunMultipleDocker() []string {
	// var wg sync.WaitGroup
	startPort := 3010
	numPorts := 5
	output := []string{}
	// wg.Add(5)
	for i := startPort; i <= startPort+numPorts-1; i++ {
		RunDocker(i)
		output = append(output, fmt.Sprintf("http://localhost:%d", i))
	}
	// wg.Wait()
	fmt.Println("Ran multiple")
	return output
}
