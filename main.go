package main

import (
	"fmt"
	a "goreloaded/function"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Get arguments passed to the programme (excluding the program name)
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: go run . sample.txt result.txt")
		return
	}
	//Read the contents of the first argument file
	File, _ := os.ReadFile(args[0])
	data := string(File)
	contents := strings.Split(data,  " ")
	for i, word := range contents {
		if word == "(hex)" {
			if i-1 >= 0 {
				contents[i-1] = strconv.Itoa(a.HexToDec(contents[i-1]))
			}
			contents = append(contents[:i], contents[i+1:]...)
		}
		if word == "(bin)" {
			if i-1 >= 0 {
				contents[i-1] = strconv.Itoa(a.BinToInt(contents[i-1]))
			}
			contents = append(contents[:i], contents[i+1:]...)
		}
	}
	// call functions from "go-reloaded/function" to manipulate the text
	mydata := a.ToUpper(strings.Join(contents,  " ")) 
	mydata = a.ToLower(mydata)
	mydata = a.Title(mydata)
	words := strings.Split(mydata, " ")
	modifiedwords := a.CheckVowel(words)
	mydata = strings.Join(modifiedwords, " ")
	mydata = a.Pun(mydata)
	mydata = a.Punc(mydata) 

	// Write to the output file "result.txt"
	Filename := "result.txt"
	err := os.WriteFile(Filename, []byte(mydata), 0644)
	if err != nil {
		fmt.Println("err", err)
	}
}
