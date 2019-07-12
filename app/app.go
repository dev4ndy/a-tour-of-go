package main

import (
	"fmt"
	"math"

	"github.com/dev4ndy/a-tour-of-go/exercises"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func main() {
	a := 75.0
	fmt.Println(fmt.Sprintf("Using math.Sqrt function %f", math.Sqrt(a)))
	fmt.Println(fmt.Sprintf("My function %f", exercises.Sqrt(15)))
	pic.Show(exercises.Pic)
	wc.Test(exercises.WordCount)
	hosts := map[string]exercises.IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
