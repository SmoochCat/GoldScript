package main

import (
	"fmt"
	"os"
)

func getFileAndParse(fileName string) {
	var buffer string = ""
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(data)

	for _, r := range text { //я еб*л того все. пушу и спать
		buffer += string(r)
		switch buffer {
		case " ", "\n", "\t":
			buffer = ""
			continue
		case ":=":
			fmt.Println("ravno")
			buffer = ""
		case ".":
			fmt.Println("dot")
			buffer = ""
		case "(":
			fmt.Println("leftScob")
			buffer = ""
		case ")":
			fmt.Println("RightScob")
			buffer = ""
		case ";":
			fmt.Println("end_line")
			buffer = ""
		case "prog":
			fmt.Println("PROGRAM")
			buffer = ""
		case "then":
			fmt.Println("NACHALO")
			buffer = ""
		case "string":
			fmt.Println("STRING")
			buffer = ""
		default:
			fmt.Println("WHAT IS THIS")
			buffer = ""
		}
	}
}

func main() {
	if len(os.Args) == 3 && os.Args[1] == "set" {
		getFileAndParse(os.Args[2])
	} else {
		fmt.Println("error")
	}
}
