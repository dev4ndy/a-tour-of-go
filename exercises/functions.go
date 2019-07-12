package exercises

import (
	"fmt"
	"math"
	"strings"
)

func Sqrt(x float64) float64 {
	z := 1.0
	auxZ := 0.2
	iterations := 0
	for math.Abs(z-auxZ) > 0.005 {
		iterations += 1
		auxZ = z
		z -= (z*z - x) / (2 * z)
	}
	fmt.Println(fmt.Sprintf("# of Iterations: %d", iterations))
	return z
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
		for j := range pic[i] {
			pic[i][j] = uint8((i + j) / 2)
		}
	}
	return pic
}

func WordCount(s string) map[string]int {
	lsWords := strings.Fields(s)
	mCountWords := make(map[string]int)
	for _, v := range lsWords {
		if _, ok := mCountWords[v]; ok {
			mCountWords[v] += 1
		} else {
			mCountWords[v] = 1
		}
	}
	return mCountWords
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	return func(x int) int {
		if x == 0 {
			return 0
		} else if x == 1 {
			return 1
		}
		f := fibonacci()
		return f(x-1) + f(x-2)
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
