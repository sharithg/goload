package config

import (
	"encoding/json"
	"fmt"
	"goload/utils"
	"io/ioutil"
	"os"
)

const CONFIG_FILE_NAME = ".goload.json"

type GoloadConfig struct {
	ImageId    string `json:"imageId"`
	ProjectDir string `json:"projectDir"`
}

func GetDockerImageName() string {

	var goloadConfig GoloadConfig

	jsonFile, err := os.Open(CONFIG_FILE_NAME)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &goloadConfig)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	return goloadConfig.ImageId
}

func DoesAttributeAndFileExist(name string) bool {
	_, err := os.Stat(CONFIG_FILE_NAME)

	if err != nil {
		return false
	}

	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return false
	}

	// if image name exists attribute and file exist
	return GetDockerImageName() != ""
}

func WriteInitialConfig(name string, projectDir string) {
	dockerImageName := name + "-" + utils.RandomId(10)

	data := GoloadConfig{ImageId: dockerImageName, ProjectDir: projectDir}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(CONFIG_FILE_NAME, file, 0644)
}

func GetProjectDir() string {
	var goloadConfig GoloadConfig

	jsonFile, err := os.Open(CONFIG_FILE_NAME)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &goloadConfig)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	return goloadConfig.ProjectDir
}
