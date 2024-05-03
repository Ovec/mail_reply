package str

import (
	"strings"

	"mailer/email"
)

func FindString(emails []*email.Email, lookFor string) int {
	for i, e := range emails {
		if strings.Contains(e.Body, lookFor) {
			return i
		}

		if strings.Contains(e.Headers.Get("subject"), lookFor) {
			return i
		}
	}
	return -1
}
