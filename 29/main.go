package main

import (
	"fmt"
	"math/big"
)

func division(a int64, b int64) *big.Int {
	return new(big.Int).Div(big.NewInt(a), big.NewInt(b))
}

func add(a int64, b int64) *big.Int {
	return new(big.Int).Add(big.NewInt(a), big.NewInt(b))
}

func mul(a int64, b int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(a), big.NewInt(b))
}

func sub(a int64, b int64) *big.Int {
	return new(big.Int).Sub(big.NewInt(a), big.NewInt(b))
}

func main() {
	// Ints
	var a int64 = 1e18
	var b int64 = 1e18
	fmt.Println("When overflows int64: ", a*b)

	// Division
	c := division(a, b)
	fmt.Println("division: ", c)

	// Sum
	c = add(a, b)
	fmt.Println("Sum: ", c)

	//Multiply
	c = mul(a, b)
	fmt.Println("Multiply: ", c)

	//Sub
	c = sub(a, b)
	fmt.Println("Sub: ", c)
}
