package printer

import (
	"fmt"

	"github.com/maeva911/otus-basic-hw/hw02_fix_app/types"
)

func PrintStaff(s []types.Employee) {
	for i := 0; i < len(s); i++ {
		str := fmt.Sprintf(
			"User ID: %d; Age: %d; Name: %s; Department ID: %d; ",
			s[i].UserID,
			s[i].Age,
			s[i].Name,
			s[i].DepartmentID)

		fmt.Println(str)
	}
}
