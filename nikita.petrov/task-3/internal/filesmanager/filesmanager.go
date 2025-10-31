package filesmanager

import (
	"fmt"
	"os"
	"path"
)

const accessMask = 0o777

func ParseOutputFilePath(outputFilePath string) (string, string) {
	return path.Dir(outputFilePath), path.Base(outputFilePath)
}

func MakeDirectory(dirPath string) error {
	if dirPath != "." {
		err := os.MkdirAll(dirPath, accessMask)
		if err != nil {
			return fmt.Errorf("failed to make directory(ies) %s: %w", dirPath, err)
		}
	}

	return nil
}

func CreateFile(dirPath string, fileName string) error {
	_, err := os.OpenFile(path.Join(dirPath, fileName), os.O_APPEND|os.O_CREATE|os.O_RDWR, accessMask)
	if err != nil {
		return fmt.Errorf("failed to create file %s at the path %s: %w", fileName, dirPath, err)
	}

	return nil
}
