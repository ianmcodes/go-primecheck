package main

import (
	"syscall/js"
	"strconv"
	"math"
	"math/rand"
	"time"
)

func main() {
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("wasmPrimeCheck", js.FuncOf(jsPrimeCheck))
	global.Set("wasmModPow", js.FuncOf(jsModPow))
	global.Set("wasmRngGenRange", js.FuncOf(jsRngGenRange))
	<-done
}

/**
 * [JS]
 * args[0]: string; Number to test.
 * args[1]: number; Number of times to run the test.
 */
func jsPrimeCheck(this js.Value, args []js.Value) interface{} {
	strN := args[0].String()
	k := args[1].Int()

	n, err := strconv.ParseInt(strN, 10, 64)

	if err != nil {
		return nil
	}

	return PrimeCheck(n, k)
}

/**
 * [JS]
 * args[0]: string; Base.
 * args[1]: string; Exponent.
 * args[2]: string; Modulus.
 */
func jsModPow(this js.Value, args []js.Value) interface{} {
	strBase := args[0].String()
	strExp := args[1].String()
	strMod := args[2].String()

	// Convert to int64
	base, err := strconv.ParseInt(strBase, 10, 64)
	exp, err := strconv.ParseInt(strExp, 10, 64)
	mod, err := strconv.ParseInt(strMod, 10, 64)

	if err != nil {
		return nil
	}
	return strconv.FormatInt(ModPow(base, exp, mod), 10)
}

/**
 * [JS]
 * args[0]: string; Minimum value.
 * args[1]: string; Maximum value.
 */
func jsRngGenRange(this js.Value, args []js.Value) interface{} {
	strMin := args[0].String()
	strMax := args[1].String()

	min, err := strconv.ParseInt(strMin, 10, 64)
	max, err := strconv.ParseInt(strMax, 10, 64)

	if err != nil {
		return nil
	}

	return strconv.FormatInt(RngGenRange(min, max), 10)
}

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
