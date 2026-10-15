package main

import (
	"fmt"
	"github.com/veraison/cmw"
	"os"
)

func readFile(fname string) []byte {
	b, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return b
}

func main() {
	fmt.Println("hello world")

	jsonCollection := readFile("collection-ok.json")

	fmt.Println(jsonCollection)

	var cmwCollection cmw.CMW
	cmwCollection.UnmarshalJSON(jsonCollection)

	fmt.Println(cmwCollection)
}
