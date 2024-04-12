package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ramdomInt generates a ramdom between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// ramdomString generate a ramdom string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// randomOwner generates a ramdom owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomBalency generantes a random amount of money
func RandomBalency() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generantes a ramdom
func RandomCurrency() string {
	currencies := []string{"USD", "CAD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
