package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(url string) ([]byte, error) {
	data, err := os.ReadFile(url)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func IsJson(filePath string) (bool, error) {
	if filePath == "" {
		return false, errors.New("file is empty")
	}
	return strings.ToLower(filepath.Ext(filePath)) == ".json", nil
}
