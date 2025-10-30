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

const accessMask = 0o777

var ErrUnknownCharset error = errors.New("unknown charset")
var ErrNotAbleToMkDir error = errors.New("cannot make directory")
var ErrCreateFile error = errors.New("cannot create file")

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func GetConfigFile(configPathFlag *string) (*os.File, error) {
	configFile, err := os.Open(*configPathFlag)

	if err != nil && os.IsNotExist(err) {
		return nil, os.ErrNotExist
	}

	return configFile, nil
}

func GetConfigData(configFile *os.File) ([]byte, error) {
	configData, err := io.ReadAll(configFile)
	if err != nil {
		return configData, io.EOF
	}

	return configData, nil
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

func MakeDirectory(dirPath string) error {
	if dirPath != "" {
		err := os.MkdirAll(dirPath, accessMask)
		if err != nil {
			return ErrNotAbleToMkDir
		}
	}

	return nil
}

func CreateFile(dirName string, fileName string) error {
	_, err := os.OpenFile(path.Join(dirName, fileName), os.O_APPEND|os.O_CREATE|os.O_RDWR, accessMask)
	if err != nil {
		return ErrCreateFile
	}

	return nil
}

func Charset(charset string, input io.Reader) (io.Reader, error) {
	switch charset {
	case "windows-1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	default:
		return nil, ErrUnknownCharset
	}
}
