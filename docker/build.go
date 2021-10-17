package docker

import (
	"bytes"
	"fmt"
	"goload/config"
	"os/exec"
)

func BuildDocker(name string) string {

	dockerImageName := name + "-" + RandomId(10)
	dockerArgs := fmt.Sprintf("docker build -f ./testdocker/Dockerfile -t %s ./testdocker", dockerImageName)

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

func RebuildDocker() string {

	dockerImageName := config.GetDockerImageName()
	dockerArgs := fmt.Sprintf("docker build -f ./testdocker/Dockerfile -t %s ./testdocker", dockerImageName)

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
