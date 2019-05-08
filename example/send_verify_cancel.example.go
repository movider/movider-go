package main

import (
	"log"

	"github.com/movider/movider-go/client"
	"github.com/movider/movider-go/verify"
)

func main() {
	c := client.New("your-api-key", "your-api-secret")

	_, err := verify.SendCancel(c, "your-request-id")
	if err != nil {
		log.Fatal(err)
	}
}
