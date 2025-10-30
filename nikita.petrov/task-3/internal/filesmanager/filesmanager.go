package filesmanager

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const accessMask = 0o777

var (
	ErrNotAbleToMkDir error = errors.New("cannot make directory")
	ErrCreateFile     error = errors.New("cannot create file")
)

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
