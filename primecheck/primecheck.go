package primecheck

import (
	"math"
	"math/rand"
	"time"
)

// PrimeCheck is a Fermat Prime Check
func PrimeCheck(n int64, k int) bool {
	if n == 1 || n == 3 {
		return true
	} else if n%2 == 0 {
		return false
	}

	for k > 0 {
		a := RngGenRange(2, n-2)
		x := ModPow(a, n-1, n)
		if x != 1 {
			return false
		}
		k--
	}

	return true
}

// ModPow does modular exponentiation
func ModPow(base int64, exponent int64, modulus int64) int64 {
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

// RngGenRange generates a random number between min and max
func RngGenRange(min int64, max int64) int64 {
	rand.Seed(time.Now().UnixNano())

	r := max - min
	return int64(math.Floor(rand.Float64()*float64(r))) + min
}
