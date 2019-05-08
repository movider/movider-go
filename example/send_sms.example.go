package main

import (
	"log"

	"github.com/movider/movider-go/client"
	"github.com/movider/movider-go/sms"
)

func main() {
	c := client.New("your-api-key", "your-api-secret")

	_, err := sms.Send(c, []string{
		"your-recipient-number",
	}, "Hello from MOVIDER-SDK", &sms.Params{})
	if err != nil {
		log.Fatal(err)
	}
}
