package reader

import (
	"encoding/json"
	"io"
	"os"

	"github.com/maeva911/otus-basic-hw/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	var data []types.Employee

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	userData, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(userData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
