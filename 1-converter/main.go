package main

import (
	"fmt"
)

func main() {
	const (
		usdPerEur = 0.8791
		usdPerRub = 78.03
	)
	src, qnt, dst := readValues()
	fmt.Printf("Source currency: %s, quantity: %.2f, target currency: %s\n", src, qnt, dst)
	res := calculate(src, qnt, dst, usdPerEur, usdPerRub)
	fmt.Printf("Итого: %.2f %s\n", res, dst)
}

func readValues() (src string, qnt float64, dst string) {
	prompt := "Введите исходную валюту (usd, eur, rub): "
	for {
		fmt.Print(prompt)
		fmt.Scan(&src)
		switch src {
		case "usd":
			prompt = "Выберите целевую валюту (eur, rub): "
		case "eur":
			prompt = "Выберите целевую валюту (usd, rub): "
		case "rub":
			prompt = "Выберите целевую валюту (usd, eur): "
		default:
			fmt.Print("\nТакой валюты не предусмотрено, повторите ввод.\n")
			continue
		}
		break
	}
	for {
		fmt.Printf("Введите количество (100.0) валюты (%s) для конвертации: ", src)
		if _, err := fmt.Scan(&qnt); err != nil {
			continue
		}
		break
	}
	for {
		fmt.Print(prompt)
		fmt.Scan(&dst)
		switch dst {
		case "usd":
		case "eur":
		case "rub":
		default:
			fmt.Print(" - валюта не предусмотрена, повторите ввод.\n")
			continue
		}
		if src == dst {
			fmt.Printf("Исходная (%s) и целевая валюты не должны совпадать!\n", src)
			continue
		}
		break
	}
	return
}

func calculate(src string, qnt float64, dst string, usdPerEur, usdPerRub float64) float64 {
	var res float64
	switch {
	case src == "usd" && dst == "eur":
		res = qnt * usdPerEur
	case src == "usd" && dst == "rub":
		res = qnt * usdPerRub
	case src == "eur" && dst == "usd":
		res = qnt / usdPerEur
	case src == "eur" && dst == "rub":
		res = qnt / usdPerEur * usdPerRub
	case src == "rub" && dst == "usd":
		res = qnt / usdPerRub
	case src == "rub" && dst == "eur":
		res = qnt * usdPerEur / usdPerRub
	}
	return res
}
