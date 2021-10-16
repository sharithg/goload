package docker

import (
	"fmt"
	"os/exec"
)

func BuildDocker(name string) {

	dockerStr := "docker build . -t " + name + "-" + RandomId(10)

	out, err := exec.Command(dockerStr).Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Print(string(out))
}
