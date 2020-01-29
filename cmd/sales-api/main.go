package main

import (
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println("error: ", err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}
