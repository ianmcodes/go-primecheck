package main

import (
	"math"
	"math/rand"
	"syscall/js"
	"time"
)

func main() {

}

/**
 * [JS]
 * args[0]: number; Number to test.
 * args[1]: number; Number of times to run the test.
 */
func jsPrimeCheck(this js.Value, args []js.Value) interface{} {
	return "Not implemented"
}

func primeCheck(n int64, k int) bool {
	if n == 1 || n == 3 {
		return true
	} else if n%2 == 0 {
		return false
	}

	for k > 0 {
		a := rngGenRange(2, n-2)
		x := modPow(a, n-1, n)
		if x != 1 {
			return false
		}
		k--
	}

	return true
}

/**
 * [JS]
 * args[0]: number; Base.
 * args[1]: number; Exponent.
 * args[2]: number; Modulus.
 */
func jsModPow(this js.Value, args []js.Value) interface{} {
	return "Not implemented"
}

func modPow(base int64, exponent int64, modulus int64) int64 {
	if modulus == 1 {
		return 0
	}
	result := int64(1)
	base = base % modulus
	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % modulus
		}
		exponent = exponent >> 1
		base = (base * base) % modulus
	}
	return result
}

/**
 * [JS]
 * args[0]: number; Minimum value.
 * args[1]: number; Maximum value.
 */
func jsRngGenRange(this js.Value, args []js.Value) interface{} {
	return "Not implemented"
}

func rngGenRange(min int64, max int64) int64 {
	rand.Seed(time.Now().UnixNano())

	r := max - min
	return int64(math.Floor(rand.Float64()*float64(r))) + min
}
