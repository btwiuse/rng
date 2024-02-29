package rng

import (
	"math/rand"
	"strings"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
)

func NewDocker() string {
	rand.Seed(time.Now().UnixNano())
	return namesgenerator.GetRandomName(0)
}

func NewDockerSep(s string) string {
	return strings.ReplaceAll(NewDocker(), "_", s)
}

func NewDockerSepDigits(s string, n uint) string {
	name := NewDockerSep(s)
	if n == 0 {
		return name
	}
	return name + s + NewDigits(n)
}
