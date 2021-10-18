package config

import (
	"encoding/json"
	"fmt"
	"goload/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

const CONFIG_FILE_NAME = ".goload.json"

type GoloadConfig struct {
	ImageId    string `json:"imageId"`
	ProjectDir string `json:"projectDir"`
}

var ProjectDirectoryPath = ""

// get the image name in the project directory
func GetDockerImageName() string {

	var goloadConfig GoloadConfig

	// open the json file
	jsonFile, err := os.Open(filepath.Join(ProjectDirectoryPath, CONFIG_FILE_NAME))

	// read bytes and unmarshall
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &goloadConfig)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	// return the image name
	return goloadConfig.ImageId
}

// check if a attribute exists and that the config file exists
func DoesAttributeAndFileExist(projectDir string, name string) bool {
	_, err := os.Stat(filepath.Join(projectDir, CONFIG_FILE_NAME))

	if err != nil {
		return false
	}

	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return false
	}

	// set the global ProjectDirectoryPath
	ProjectDirectoryPath = projectDir

	// if image name exists attribute and file exist
	return GetDockerImageName() != ""
}

// write the initial goload config file
func WriteInitialConfig(name string, projectDir string) {
	dockerImageName := name + "-" + utils.RandomId(10)

	data := GoloadConfig{ImageId: dockerImageName, ProjectDir: projectDir}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(filepath.Join(projectDir, CONFIG_FILE_NAME), file, 0644)

	ProjectDirectoryPath = projectDir
}

// get the current project directory
func GetProjectDir() string {
	var goloadConfig GoloadConfig

	jsonFile, err := os.Open(filepath.Join(ProjectDirectoryPath, CONFIG_FILE_NAME))

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &goloadConfig)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	return goloadConfig.ProjectDir
}
