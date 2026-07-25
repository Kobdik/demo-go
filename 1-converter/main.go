package main

import (
	"fmt"
)

func main() {
	const (
		usdPerEur = 0.8791
		usdPerRub = 78.03
	)
	var eurPerRub float64 = usdPerRub / usdPerEur
	fmt.Println(eurPerRub)
}

func readValues() (p string) {
	fmt.Scan(&p)
	return
}

func calculate(val float64, src, dst string) float64 {
	fmt.Printf("Sorce %s, target %s\n", src, dst)
	return val
}
