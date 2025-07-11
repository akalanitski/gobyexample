// _Функцыі_ цэнтральная канцэпцыя ў Go.
// Мы даведаемся пра функцыі па некалькіх прыкладах.

package main

import "fmt"

// Вось функцыя, якая прымае два лікі тыпу `int` і вяртае
// іх суму тыпу `int`.
func plus(a int, b int) int {

	// Go патрабуе відавочных вяртанняў, г.зн. ён не будзе
	// аўтаматычна вяртаць значэнне апошняга
	// выразу.
	return a + b
}

// Калі ў вас ёсць некалькі паслядоўных параметраў аднаго тыпу, вы можаце
// апусціць назву тыпу для параметраў падобнага тыпу да апошняга параметра,
// які аб'яўляе тып.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Выклікаем функцыю, як і чакалася, з дапамогай
	// `назва(аргументы)`.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
