package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Hello, World!")
	fileContent, err := os.Open("../thai-wordlist/thai-wordlist.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Read file success!")
	defer fileContent.Close()
}
