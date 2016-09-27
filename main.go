package main

import (
	"fmt"
	"walk/core"
)

func main() {
	if e := core.Walk(string(*Root), string(*Ignore), string(*Suffix), OutFilePath); e != nil {
		fmt.Println(e)
	}
	fmt.Println("over")
}
