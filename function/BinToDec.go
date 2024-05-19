package function

import (
	"fmt"
	"strconv"
)

// function to convert binary to integer
func BinToInt(bin string) int {
	decimal_num, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Println("Error", err)
		return 0
	}
	return int(decimal_num)
}
