package file

import (
	"errors"
	"path/filepath"
	"strings"
)

type FileHandler interface {
	IsJson(filePath string) (bool, error)
}

type DefaultFileHandler struct{}

func (d *DefaultFileHandler) IsJson(filePath string) (bool, error) {
	if filePath == "" {
		return false, errors.New("FILE PATH IS EMPTY")
	}
	return strings.ToLower(filepath.Ext(filePath)) == ".json", nil
}
