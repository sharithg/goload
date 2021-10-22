package docker

import (
	"bufio"
	"fmt"
	"goload/utils"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func GetExposedPort(projectDir string) int {
	dockerfilePath := filepath.Join(projectDir, "Dockerfile")

	file, err := os.Open(dockerfilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	// Read through 'tokens' until an EOF is encountered.
	for sc.Scan() {
		res := strings.Contains(sc.Text(), "EXPOSE")
		if res {
			words := strings.Fields(sc.Text())
			if len(words) == 2 {
				intVar, _ := strconv.Atoi(words[1])
				return intVar
			}
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	yellow := color.New(color.FgYellow).SprintFunc()
	errorMessage := fmt.Sprintf("Could not find exposed port, please make sure you have %s in your Dockerfile", yellow("EXPOSE <port number>"))
	utils.FatalError(errorMessage)
	return 0
}
