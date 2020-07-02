package main

import (
	"syscall/js"
	"strconv"
	p "ianmccall.codes/go/primecheck/primecheck"
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

	return p.PrimeCheck(n, k)
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
	return strconv.FormatInt(p.ModPow(base, exp, mod), 10)
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

	return strconv.FormatInt(p.RngGenRange(min, max), 10)
}
