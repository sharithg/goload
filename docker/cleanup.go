package docker

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func StopDocker(id string) {

	// defer wg.Done()

	dockerCmd := fmt.Sprintf("docker rm -f %s", id)

	dockerExec := exec.Command("/bin/sh", "-c", dockerCmd)

	var outb, errb bytes.Buffer
	dockerExec.Stdout = os.Stdout
	dockerExec.Stderr = os.Stderr

	// if there is an error with our execution
	// handle it here
	err := dockerExec.Run()

	if err != nil {
		fmt.Println("out:", outb.String(), "err:", errb.String())
	}
}
