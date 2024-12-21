package authentication

import (
	"fmt"
	"math/rand"
	"time"
)

func generateToken() string {
	randomNumber := rand.Intn(900000) + 100000

	return fmt.Sprintf("%d", randomNumber)
}

func getTokenExpiresAt() time.Time {
	return time.Now().Add(5 * time.Minute)
}

func isTokenExpired(expiresAt time.Time) bool {
	currentTime := time.Now()

	return currentTime.After(expiresAt)
}

func isTokenValid(reqOtp *string, vOtp string) bool {
	if reqOtp == nil {
		return false
	}

	return *reqOtp == vOtp
}

func BoolPtr(b bool) *bool {
	return &b
}
