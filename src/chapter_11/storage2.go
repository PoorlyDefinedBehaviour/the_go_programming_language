package storage2

import (
	"fmt"
	"net/smtp"

	"github.com/pkg/errors"
)

// Context:
// This is supposed to be a web service that provides
// networked storage to users. When users exceed 90% of their quota,
// the system sends them a warning email.

// NOTE: never put credentials in source code
const sender = "example@example.com"
const password = "password"
const hostname = "smtp.example.com"

var notifyUser = func(username, message string) error {
	auth := smtp.PlainAuth("", sender, password, hostname)

	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(message))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func bytesInUse(username string) int64 {
	// get bytes in use from somewhere
	return 0
}

const gigabyte = 1000000000

func percentOfStorageQuotaInUse(username string) int64 {
	used := bytesInUse(username)

	const quota = 1 * gigabyte

	return 100 * used / quota
}

func CheckQuota(username string) error {
	used := bytesInUse(username)

	const quota = 1 * gigabyte

	quotaPercentageInUse := 100 * used / quota

	if quotaPercentageInUse < 90 {
		return nil
	}

	message := fmt.Sprintf("Warning: you are using %d bytes of storage, %d%% of your quota.", used, quotaPercentageInUse)

	err := notifyUser(username, message)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
