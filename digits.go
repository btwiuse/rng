package rng

import (
	"fmt"
	"math/rand"
)

func NewDigits(n uint) (digits string) {
	for i := n; i > 0; i-- {
		digits += fmt.Sprintf("%d", rand.Intn(10))
	}
	return
}
