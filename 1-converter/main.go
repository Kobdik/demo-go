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
