package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	// output: 0°C
	fmt.Printf("%gF = %g°C\n", freezingF, fToC(freezingF))
	// output: 100°C
	fmt.Printf("%gF = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
