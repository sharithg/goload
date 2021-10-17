package docker

import (
	"bytes"
	"fmt"
	"goload/config"
	"goload/utils"
	"os/exec"
	"path/filepath"
)

// Builds a docker image given a name
func BuildDocker(projectDir string, dockerImageName string) (string, string) {

	dockerBuildDir := "."
	dockerFileBuildPath := ""

	if projectDir != "" {
		dockerBuildDir = projectDir
		dockerFileBuildPath = fmt.Sprintf("-f %s", filepath.Join(projectDir, "Dockerfile"))
	}

	dockerArgs := fmt.Sprintf("docker build %s -t %s %s", dockerFileBuildPath, dockerImageName, dockerBuildDir)

	fmt.Println("Building docker image...")

	dockerExec := exec.Command("/bin/sh", "-c", dockerArgs)

	var outb, errb bytes.Buffer
	dockerExec.Stdout = &outb
	dockerExec.Stderr = &errb

	// if there is an error with our execution
	// handle it here
	err := dockerExec.Run()

	if err != nil {
		utils.FatalError(errb.String())
	}

	fmt.Println("out:", outb.String(), "err:", errb.String())

	return dockerImageName, projectDir
}

func RebuildDocker() string {

	dockerImageName := config.GetDockerImageName()
	projectDir := config.GetProjectDir()

	dockerfilePath := filepath.Join(projectDir, "Dockerfile")

	dockerArgs := fmt.Sprintf("docker build -f %s -t %s %s", dockerfilePath, dockerImageName, projectDir)

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
