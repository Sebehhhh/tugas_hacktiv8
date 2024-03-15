// internal/utils/file_writer.go

package utils

import (
	"encoding/json"
	"io/ioutil"
	"tugas3-hacktiv8/internal/model"
)

func WriteStatusToFile(status model.Status, filePath string) error {
	data, err := json.Marshal(status)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
