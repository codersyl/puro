package main

import (
	"io" // io.EOF
	"log"
	"os" // open a file need "os"
)

func main() { // in fact, this program is to print the file os.Args[1]
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	// answer for "how to open a file?"
	f, err := os.Open(os.Args[1]) // type of f: *File
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// answer end

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
