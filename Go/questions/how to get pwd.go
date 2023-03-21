package main

import (
	"fmt"
	"log"
	"os" // need "os" to get pwd
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}
