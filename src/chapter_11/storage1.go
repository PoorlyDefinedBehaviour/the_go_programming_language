package storage1

import (
	"fmt"
	"log"
	"net/smtp"
)

// Context:
// This is supposed to be a web service that provides
// networked storage to users. When users exceed 90% of their quota,
// the system sends them a warning email.

func bytesInUse(username string) int64 {
	// get bytes in use from somewhere
	return 0
}

// NOTE: never put credentials in source code
const sender = "example@example.com"
const password = "password"
const hostname = "smtp.example.com"

const template = "Warning: you are using %d bytes of storage, %d%% of your quota."

const gigabyte = 1000000000

// This could be better, is not easy to test this code
// unless smtp has mocking capabilities.
// If it does not, emails would be sent during tests,
// which is undesirable.
func CheckQuota(username string) {
	used := bytesInUse(username)

	const quota = 1 * gigabyte

	percentageInUse := 100 * used / quota

	if percentageInUse < 90 {
		return
	}

	message := fmt.Sprintf(template, used, percentageInUse)

	auth := smtp.PlainAuth("", sender, password, hostname)

	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(message))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}
