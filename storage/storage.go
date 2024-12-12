package storage

import (
	"encoding/json"
	"errors"
	"json-cli/bins"
	"os"
)

type Storage interface {
	Save(filePath string, binList *bins.BinList) error
	Load(filePath string) (*bins.BinList, error)
}

type JSONStorage struct{}

func (s *JSONStorage) Save(filePath string, binList *bins.BinList) error {
	data, err := json.Marshal(binList)
	if err != nil {
		return errors.New("FILE PATH ERROR")
	}
	return os.WriteFile(filePath, data, 0644)
}
func (s *JSONStorage) Load(filePath string) (*bins.BinList, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("FILE PATH ERROR")
	}
	var binList bins.BinList
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return nil, errors.New("JSON not be Unmarchal")
	}
	return &binList, nil
}
