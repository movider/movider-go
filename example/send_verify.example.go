package main

import (
	"log"

	"github.com/movider/movider-go/client"
	"github.com/movider/movider-go/verify"
)

func main() {
	c := client.New("your-api-key", "your-api-secret")

	_, err := verify.Send(c, "your-recipient-number", &verify.Params{})
	if err != nil {
		log.Fatal(err)
	}
}
