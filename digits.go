package rng

import (
	"fmt"
	"math/rand/v2"
)

func NewDigits(n uint) (digits string) {
	for i := n; i > 0; i-- {
		digits += fmt.Sprintf("%d", rand.IntN(10))
	}
	return
}
