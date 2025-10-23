package fmanager

import (
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var ErrUnknownCharset error = errors.New("unknown charset")

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
		panic(err)
	}

	return configData
}

func ParseOutputFilePath(outputFilePath string) (string, string) {
	var path string

	var filename string

	if strings.Contains(outputFilePath, "/") {
		pathAndFile := strings.Split(outputFilePath, "/")
		path = ""

		for i := range len(pathAndFile) - 1 {
			path = filepath.Join(path, pathAndFile[i])
		}

		filename = pathAndFile[len(pathAndFile)-1]
	} else {
		path = ""
		filename = outputFilePath
	}

	return path, filename
}

func MakeDirectory(dirPath string) {
	if dirPath != "" {
		err := os.MkdirAll(dirPath, 0777)
		if err != nil {
			panic(err)
		}
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
		return nil, ErrUnknownCharset
	}
}
