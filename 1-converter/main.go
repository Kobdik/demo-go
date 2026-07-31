package main

import (
	"fmt"
)

type DstMap = *map[string]float64

func main() {
	const (
		usdPerEur = 0.8791
		usdPerRub = 78.03
	)
	srcMap := createMap(usdPerEur, usdPerRub)
	src, qnt, dst, price := readValues(&srcMap)
	fmt.Printf("Source currency: %s, quantity: %.2f, target currency: %s\n", src, qnt, dst)
	cost := qnt * price // quantity of source * price of source in target currency
	fmt.Printf("Итого: %.2f %s\n", cost, dst)
}

// map in map, keys: source and target currencies, value: price of source in target currency
func createMap(usdPerEur, usdPerRub float64) map[string]DstMap {
	return map[string]DstMap{
		"usd": &map[string]float64{
			"eur": usdPerEur,
			"rub": usdPerRub,
		},
		"eur": &map[string]float64{
			"usd": 1.0 / usdPerEur,
			"rub": usdPerRub / usdPerEur,
		},
		"rub": &map[string]float64{
			"usd": 1.0 / usdPerRub,
			"eur": usdPerEur / usdPerRub,
		},
	}
}

func readValues(srcMap *map[string]DstMap) (src string, qnt float64, dst string, price float64) {
	var (
		dstMap DstMap
		ok     bool
	)
	for {
		fmt.Print("Введите исходную валюту (usd, eur, rub): ")
		fmt.Scan(&src)
		dstMap, ok = (*srcMap)[src]
		if !ok {
			fmt.Println("Недопустимый ввод, повторите ввод исходной валюты.")
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
		fmt.Print("Введите целевую валюту ( ")
		for dst = range *dstMap {
			fmt.Printf("%s ", dst)
		}
		fmt.Print("): ")
		fmt.Scan(&dst) // целевая валюта
		price, ok = (*dstMap)[dst]
		if !ok {
			fmt.Println("Недопустимый ввод, повторите ввод целевой валюты.")
			continue
		}
		break
	}
	return
}

// obsolete
// func calculate(src string, qnt float64, dst string, usdPerEur, usdPerRub float64) float64 {
// 	var res float64
// 	switch {
// 	case src == "usd" && dst == "eur":
// 		res = qnt * usdPerEur
// 	case src == "usd" && dst == "rub":
// 		res = qnt * usdPerRub
// 	case src == "eur" && dst == "usd":
// 		res = qnt / usdPerEur
// 	case src == "eur" && dst == "rub":
// 		res = qnt / usdPerEur * usdPerRub
// 	case src == "rub" && dst == "usd":
// 		res = qnt / usdPerRub
// 	case src == "rub" && dst == "eur":
// 		res = qnt * usdPerEur / usdPerRub
// 	}
// 	return res
// }
