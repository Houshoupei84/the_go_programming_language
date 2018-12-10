package main

import "fmt"

func main(){

	const freezingF , boilingF = 32.0, 212.0

	fmt.Printf("%g F = %g c\n", freezingF, ftoc(freezingF))
	fmt.Printf("%g F = %g c\n", boilingF, ftoc(boilingF))
}

func ftoc(f float64) float64 {
	return (f -32) * 5 / 9
}