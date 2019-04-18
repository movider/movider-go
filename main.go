package main

import (
	"log"
	"movider"
	"movider/sms"
	"time"
)

func main() {
	client := movider.New("your-api-key", "your-api-secret")
	// client := movider.New("1JezeMGZO10gP5bgev7qKFxrakn", "1z5E7XQdcHGZX6FiLQRM9kz721Z6GSfjxNwCv0K9")

	d, err := sms.Send(client, []string{
		"66847827374",
	}, "Hello from golang sdk", &sms.Params{})
	if err != nil {
		log.Fatalf("[Movider] Error: %s\n", err.Error())
	}

	for _, v := range d.PhoneNumberList {
		log.Printf("[Movider] Send SMS to: %s completed.\n", v.Number)
		time.Sleep(time.Second)
	}

	for _, v := range d.BadPhoneNumberList {
		log.Printf("[Movider] Cannot send SMS to: %d because %s\n", v.Number, v.Msg)
		time.Sleep(time.Second)
	}

	log.Printf("[Movider] Your balance is $ %.8f \n", d.RemainingBalance)
	log.Println("[Movider] Sent SMS complete.")
}
