package docker

import (
	"bytes"
	"fmt"
	"os/exec"
)

func BuildDocker(name string) {

	dockerImageName := name + "-" + RandomId(10)
	dockerArgs := "docker build -f ./testdocker/Dockerfile -t " + dockerImageName + " ."

	fmt.Print(dockerArgs + "\n")

	dockerExec := exec.Command("/bin/sh", "-c", dockerArgs)

	var outb, errb bytes.Buffer
	dockerExec.Stdout = &outb
	dockerExec.Stderr = &errb

	// if there is an error with our execution
	// handle it here
	err := dockerExec.Run()

	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("out:", outb.String(), "err:", errb.String())
}
