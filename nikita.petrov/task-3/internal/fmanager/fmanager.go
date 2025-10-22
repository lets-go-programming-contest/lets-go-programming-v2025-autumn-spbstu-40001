package fmanager

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"golang.org/x/text/encoding/charmap"
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
		dir = ""
		filename = outputFilePath
	}

	return dir, filename
}

func MakeDirectory(dirName string) {
	if dirName != "" {
		os.Mkdir(dirName, 0777)
	}
}

func CreateFile(dirName string, fileName string) {
	_, errCreateFile := os.OpenFile(path.Join(dirName, fileName), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if errCreateFile != nil {
		panic(errCreateFile)
	}
}

func Charset(charset string, input io.Reader) (io.Reader, error) {
	switch charset {
	case "windows-1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	default:
		return nil, fmt.Errorf("unknown charset: %s", charset)
	}
}
