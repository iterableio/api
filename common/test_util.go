package common

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomEmail() string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x@example.com", b)
}
