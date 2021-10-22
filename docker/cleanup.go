package docker

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// cleanup function which stops an individual docker container given a id
func StopDocker(id string) {

	// defer wg.Done()

	dockerCmd := fmt.Sprintf("docker rm -f %s", id)

	dockerExec := exec.Command("/bin/sh", "-c", dockerCmd)

	var outb, errb bytes.Buffer
	// dockerExec.Stdout = os.Stdout
	dockerExec.Stderr = os.Stderr

	err := dockerExec.Run()

	if err != nil {
		fmt.Println("out:", outb.String(), "err:", errb.String())
	}
}
