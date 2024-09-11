package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	var data []types.Employee

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	userData, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	err = json.Unmarshal(userData, &data)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	res := data

	return res, nil
}
