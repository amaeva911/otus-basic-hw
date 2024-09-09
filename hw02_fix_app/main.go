package main

import (
	"fmt"
	"os"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func main() {
	var path string
	var err error
	var staff []types.Employee

	fmt.Printf("Enter data file path: ")
	fmt.Fscan(os.Stdin, &path)

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path, -1)
	if err != nil {
		fmt.Println(err)
	}

	printer.PrintStaff(staff)
}
