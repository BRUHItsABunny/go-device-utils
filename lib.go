package go_device_utils

import (
	"math/rand"
	"time"
)

func init() {
	// Guarantee a seed
	rand.Seed(time.Now().UnixNano())
}
