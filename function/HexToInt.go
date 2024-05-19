package function

import (
	"fmt"
	"strconv"
)

//function to convert hexadecimal to integer
func HexToDec(hex string) int {
	decimal, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		fmt.Println("An error", err)
		return 0
	}
	return int(decimal)
}
