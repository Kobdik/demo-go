package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// no operation selected
const NOP string = "NOP"

type Calc struct {
	operKey  string // selected operation
	operMap  map[string]func([]float64) float64
	operData []float64
}

// Создаем "пустой" калькулятор
func CreateCalc() *Calc {
	return &Calc{
		operKey: NOP,
		operMap: make(map[string]func([]float64) float64),
	}
}

// Добавляем в калькулятор операции
func (calc *Calc) AddOper(key string, oper func([]float64) float64) {
	if key != "" && key != NOP {
		calc.operMap[key] = oper
	}
}

// Выбираем допустимую операцию
func (calc *Calc) ReadOper() {
	var (
		prompt string = "Введите операцию (AVG, SUM, MED)"
		done   bool   = false
		key    string
	)
	for !done {
		fmt.Println(prompt)
		fmt.Scan(&key)
		if _, done = calc.operMap[key]; done {
			calc.operKey = key
		} else {
			calc.operKey = NOP
			fmt.Println("Такой операции не предусмотрено, повторите ввод.")
			continue
		}
	}
}

// Заполняем список аргументов операции
func (calc *Calc) ReadValues() {
	var (
		raw string
		num float64
		err error
	)
	fmt.Println("Введите числа, без пробелов, разделенные запятыми (1,2,3):")
	fmt.Scan(&raw)
	// слайс строк
	ss := strings.Split(raw, ",")
	nums := make([]float64, 0, len(ss))
	// конвертируем элементы в числа float64
	for _, s := range ss {
		num, err = strconv.ParseFloat(s, 64)
		if err == nil {
			nums = append(nums, num)
		}
	}
	// сохраняем для дальнейшего использования
	calc.operData = nums
}

func (calc *Calc) Calculate() (float64, error) {
	var key string = calc.operKey
	if key == NOP {
		return 0.0, errors.New("No operation selected!")
	}
	oper := calc.operMap[key]
	return oper(calc.operData), nil
}

func main() {
	calc := CreateCalc()
	calc.AddOper("AVG", calcAvg)
	calc.AddOper("MED", calcMed)
	calc.AddOper("SUM", calcSum)
	calc.ReadOper()
	calc.ReadValues()

	result, err := calc.Calculate()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Oper=%s, result=%.2f\n", calc.operKey, result)
	}
}

// func readValues() (op string, raw string) {
// 	prompt := "Введите операцию (AVG, SUM, MED)"
// 	done := false
// 	for !done {
// 		fmt.Println(prompt)
// 		fmt.Scan(&op)
// 		switch op {
// 		case "AVG":
// 			done = true
// 		case "SUM":
// 			done = true
// 		case "MED":
// 			done = true
// 		default:
// 			fmt.Println("Такой операции не предусмотрено, повторите ввод.")
// 			continue
// 		}
// 	}
// 	fmt.Println("Введите числа, без пробелов, разделенные запятыми (1,2,3):")
// 	fmt.Scan(&raw)
// 	return
// }

// func parseRawData(raw string) []float64 {
// 	ss := strings.Split(raw, ",")
// 	tr := make([]float64, 0, len(ss))
// 	var (
// 		f float64
// 		e error
// 	)
// 	for _, s := range ss {
// 		f, e = strconv.ParseFloat(s, 64)
// 		if e == nil {
// 			tr = append(tr, f)
// 		}
// 	}
// 	return tr
// }

// func calcOper(op string, ts []float64) float64 {
// 	var res float64
// 	switch op {
// 	case "AVG":
// 		res = calcAvg(ts)
// 	case "SUM":
// 		res = calcSum(ts)
// 	case "MED":
// 		res = calcMed(ts)
// 	}
// 	return res
// }

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
