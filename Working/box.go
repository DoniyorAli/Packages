package numbox

import (
	"strconv"
)

type Work struct {
	Value string
}

func Add(num1, num2 Work) Work {
	bag := Work{}
	first, _ := strconv.Atoi(num1.Value)
	second, _ := strconv.Atoi(num2.Value)
	bag.Value = strconv.Itoa(first + second)
	return bag
}

func Sub(num1, num2 Work) Work {
	bag := Work{}
	first, _ := strconv.Atoi(num1.Value)
	second, _ := strconv.Atoi(num2.Value)
	bag.Value = strconv.Itoa(first - second)
	return bag
}

func Multiples(num1, num2 Work) Work {
	bag := Work{}
	first, _ := strconv.Atoi(num1.Value)
	second, _ := strconv.Atoi(num2.Value)
	bag.Value = strconv.Itoa(first * second)
	return bag
}

func Division(num1, num2 Work) Work {
	bag := Work{}
	first, _ := strconv.Atoi(num1.Value)
	second, _ := strconv.Atoi(num2.Value)
	bag.Value = strconv.Itoa(first / second)
	return bag
}

func Mod(num1, num2 Work) Work {
	bag := Work{}
	first, _ := strconv.Atoi(num1.Value)
	second, _ := strconv.Atoi(num2.Value)
	bag.Value = strconv.Itoa(first % second)
	return bag
}

// func Abs(num Work) Work {
// 	bag := Work{}
// 	f := num.Value
// 	x, _ := strconv.ParseFloat(f, 64)
// 	if x < 0 {
// 		bag.Value = math.Abs(x)
// 		return bag
// 	}
// }