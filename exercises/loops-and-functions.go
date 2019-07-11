/*
Exercise: Loops and Functions
As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.

Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:
z -= (z*z - x) / (2*z)
Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.

Implement this in the func Sqrt provided. A decent starting guess for z is 1, no matter what the input. To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.

Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small amount). See if that's more or fewer than 10 iterations. Try other initial guesses for z, like x, or x/2. How close are your function's results to the math.Sqrt in the standard library?

*/
package main

import (
	"fmt"
	"math"
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
	fmt.Println(fmt.Sprintf("Iterations: %d", iterations))
	return z
}

func main() {
	a := 75.0
	fmt.Println(fmt.Sprintf("Using math.Sqrt function %f", math.Sqrt(a)))
	fmt.Println(fmt.Sprintf("My function %f", Sqrt(a)))
}
