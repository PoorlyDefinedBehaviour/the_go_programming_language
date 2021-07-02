package storage2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckQuota(t *testing.T) {
	t.Parallel()

	t.Run("980MB out of 1GB used", func(t *testing.T) {
		var notifiedUser string
		var notifiedMessage string

		// This is a bad solution because we need to change
		// notifyUser back after each test
		// and Go does not provide a clean way to do it.
		notifyUser = func(user, message string) error {
			notifiedUser = user
			notifiedMessage = message

			return nil
		}

		const user = "joe@example.org"

		fmt.Printf("\n\naaaaaaa  %+v\n\n", "here")
		err := CheckQuota(user)

		assert.Nil(t, err)

		assert.Equal(t, user, notifiedUser)
		assert.Equal(t, "Warning: you are using 980 bytes of storage, 98% of your quota.", notifiedMessage)
	})
}
