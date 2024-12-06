package storage

import (
	"encoding/json"
	"json-cli/bins"
	"os"
)

func SaveBinList(filePath string, binList *bins.BinList) error {
	data, err := json.MarshalIndent(binList, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}
func LoadBinList(filePath string) (*bins.BinList, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var binList bins.BinList
	if err := json.Unmarshal(data, &binList); err != nil {
		return nil, err
	}

	return &binList, nil
}
