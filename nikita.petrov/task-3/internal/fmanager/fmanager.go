package fmanager

import (
	"io"
	"os"
	"path"
	"strings"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func GetConfigFile(configPathFlag *string) *os.File {
	configFile, err := os.Open(*configPathFlag)
	if err != nil && os.IsNotExist(err) {
		panic(err)
	}

	return configFile
}

func GetConfigData(configFile *os.File) []byte {
	configData, err := io.ReadAll(configFile)
	if err != nil {
		panic("cannot read file")
	}

	return configData
}

func ParseOutputFilePath(outputFilePath string) (string, string) {
	var dir string

	var filename string

	if strings.Contains(outputFilePath, "/") {
		outputFilePath := strings.Split(outputFilePath, "/")
		dir = outputFilePath[0]
		filename = outputFilePath[1]
	} else {
		filename = outputFilePath
	}

	return dir, filename
}

func MakeDirectory(dirName string) {
	if dirName != "" {
		errCreateDir := os.Mkdir(dirName, 0777)
		if errCreateDir != nil {
			panic(errCreateDir)
		}
	}
}

func CreateFile(dirName string, fileName string) {
	_, errCreateFile := os.OpenFile(path.Join(dirName, fileName), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if errCreateFile != nil {
		panic(errCreateFile)
	}
}
