package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type GoloadConfig struct {
	ImageId string `json:"imageId"`
}

func GetDockerImageName() string {

	var goloadConfig GoloadConfig

	jsonFile, err := os.Open(".config.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &goloadConfig)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	return goloadConfig.ImageId
}

func DoesAttributeAndFileExist(name string) bool {
	_, err := os.Stat(".config.json")

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

	_ = ioutil.WriteFile(".config.json", file, 0644)
}
