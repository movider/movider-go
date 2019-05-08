package main

import (
	"log"

	"github.com/movider/movider-go/client"
	"github.com/movider/movider-go/verify"
)

func main() {
	c := client.New("your-api-key", "your-api-secret")

	_, err := verify.SendAcknowledge(c, "your-request-id", "your-code")
	if err != nil {
		log.Fatal(err)
	}
}
