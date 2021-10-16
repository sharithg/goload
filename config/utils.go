package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const CONFIG_FILE_NAME = ".goload.json"

type GoloadConfig struct {
	ImageId string `json:"imageId"`
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

func WriteDockerImageName(name string) {

	data := GoloadConfig{
		ImageId: name,
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(CONFIG_FILE_NAME, file, 0644)
}
