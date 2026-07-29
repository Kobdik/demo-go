package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	op, data := readValues()
	trs := parseRawData(data)

	fmt.Printf("Oper=%s, result=%.2f\n", op, calcOper(op, trs))
}

func readValues() (op string, raw string) {
	prompt := "Введите операцию (AVG, SUM, MED)"
	done := false
	for !done {
		fmt.Println(prompt)
		fmt.Scan(&op)
		switch op {
		case "AVG":
			done = true
		case "SUM":
			done = true
		case "MED":
			done = true
		default:
			fmt.Println("Такой операции не предусмотрено, повторите ввод.")
			continue
		}
	}
	fmt.Println("Введите числа, без пробелов, разделенные запятыми (1,2,3):")
	fmt.Scan(&raw)
	return
}

func parseRawData(raw string) []float64 {
	ss := strings.Split(raw, ",")
	tr := make([]float64, 0, len(ss))
	var (
		f float64
		e error
	)
	for _, s := range ss {
		f, e = strconv.ParseFloat(s, 64)
		if e == nil {
			tr = append(tr, f)
		}
	}
	return tr
}

func calcOper(op string, ts []float64) float64 {
	var res float64
	switch op {
	case "AVG":
		res = calcAvg(ts)
	case "SUM":
		res = calcSum(ts)
	case "MED":
		res = calcMed(ts)
	}
	return res
}

func calcAvg(ts []float64) float64 {
	return calcSum(ts) / float64(len(ts))
}

func calcSum(ts []float64) float64 {
	var sum float64 = 0.0
	for _, val := range ts {
		sum += val
	}
	return sum
}

func calcMed(ts []float64) float64 {
	if len(ts) == 0 {
		panic("Пустой массив не допустим!")
	}
	var (
		min float64 = ts[0]
		max float64 = ts[0]
	)
	for _, val := range ts {
		if min > val {
			min = val
		}
		if max < val {
			max = val
		}
	}
	return (min + max) / 2
}
