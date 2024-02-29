package main

import (
	"fmt"

	"github.com/btwiuse/rng"
)

func main() {
	fmt.Println(rng.NewDocker())
	fmt.Println(rng.NewDockerSep("-"))
	fmt.Println(rng.NewDigits(4))
	fmt.Println(rng.NewDockerSepDigits("-", 4))
	fmt.Println(rng.NewUUID())
	fmt.Println(rng.NewPetUUID(""))
}
