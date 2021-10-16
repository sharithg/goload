package docker

import (
	"bytes"
	"fmt"
	"os/exec"
)

func BuildDocker(name string) string {

	dockerImageName := name + "-" + RandomId(10)
	dockerArgs := "docker build -f ./testdocker/Dockerfile -t " + dockerImageName + " ."

	fmt.Print(dockerArgs + "\n")

	fmt.Println("Building docker image...")

	dockerExec := exec.Command("/bin/sh", "-c", dockerArgs)

	var outb, errb bytes.Buffer
	dockerExec.Stdout = &outb
	dockerExec.Stderr = &errb

	// if there is an error with our execution
	// handle it here
	err := dockerExec.Run()

	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}

	fmt.Println("out:", outb.String(), "err:", errb.String())

	return dockerImageName
}
