// Go падтрымлівае _канстанты_ сімвальных, тэкставых, булевых і лікавых
// тыпаў

package main

import (
	"fmt"
	"math"
)

// `const` аб'яўляе канстантнае значэнне.
const s string = "constant"

func main() {
	fmt.Println(s)

	// Аператар `const` можа з'яўляцца ў любым месцы,
	// дзе можа з'яўляцца аператар `var`
	const n = 500000000

	// Канстантныя выразы выконваюць арыфметычныя аперацыі з
	// адвольнай дакладнасцю.
	const d = 3e20 / n
	fmt.Println(d)

	// Лікавая канстанта не мае тыпу, пакуль ёй не зададзены тып,
	// напрыклад, шляхам відавочнага пераўтварэння.
	fmt.Println(int64(d))

	// Ліку можна надаць тып, выкарыстоўваючы яго ў
	// кантэксце, які патрабуе тыпу, напрыклад, зменнай
	// прысвойванні або выкліку функцыі. Напрыклад, тут
	// `math.Sin` чакае `float64`.
	fmt.Println(math.Sin(n))
}
