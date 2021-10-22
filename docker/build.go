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
func BuildDocker(projectDir string, name string) (string, string) {

	dockerBuildDir := "."
	dockerFileBuildPath := ""

	dockerImageName := name + "-" + utils.RandomId(10)

	// if the project dir is not empty update params to use this directory
	if projectDir != "" {
		dockerBuildDir = projectDir
		dockerFileBuildPath = fmt.Sprintf("-f %s", filepath.Join(projectDir, "Dockerfile"))
	}

	// construct docker command
	dockerArgs := fmt.Sprintf("docker build %s -t %s %s", dockerFileBuildPath, dockerImageName, dockerBuildDir)

	fmt.Println("Building docker image...")

	// exec this command
	dockerExec := exec.Command("/bin/sh", "-c", dockerArgs)

	var outb, errb bytes.Buffer
	dockerExec.Stdout = &outb
	dockerExec.Stderr = &errb

	err := dockerExec.Run()

	// exit the process if there was an error
	if err != nil {
		utils.FatalError(errb.String())
	}

	fmt.Println("out:", outb.String(), "err:", errb.String())

	return dockerImageName, projectDir
}

// rebuilds a existing docker image
func RebuildDocker() string {

	goloadConfig := config.LoadConfig()
	projectDir := goloadConfig.GetProjectDir()

	dockerfilePath := filepath.Join(projectDir, "Dockerfile")

	dockerArgs := fmt.Sprintf("docker build -f %s -t %s %s", dockerfilePath, goloadConfig.ImageId, projectDir)

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

	return goloadConfig.ImageId

}
