package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const CONFIG_FILE_NAME = ".goload.json"

type GoloadConfig struct {
	ProjectName string `json:"projectName"`
	ImageId     string `json:"imageId"`
	ProjectDir  string `json:"projectDir"`
	ExposedPort int    `json:"exposedPort"`
}

var ProjectDirectoryPath = ""

// get the image name in the project directory
func LoadConfig() GoloadConfig {

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
	return goloadConfig
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
	return LoadConfig().ImageId != ""
}

// write the initial goload config file
func (gc *GoloadConfig) Write() {

	data := *gc

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(filepath.Join(gc.ProjectDir, CONFIG_FILE_NAME), file, 0644)

	ProjectDirectoryPath = gc.ProjectDir
}

// get the current project directory
func (gc *GoloadConfig) GetProjectDir() string {
	return gc.ProjectDir
}
