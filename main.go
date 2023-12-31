package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")

	Port := os.Getenv("PORT")

	if Port == "" {
		log.Fatal("Port is not found in the Environmnet")
	}
	fmt.Println("Port:", Port)
}
