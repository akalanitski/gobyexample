// Вырыятыўныя функцыі [_Variadic functions_](https://en.wikipedia.org/wiki/Variadic_function)
// можна выклікаць з любой колькасцю аргументаў.
// Напрыклад, `fmt.Println` — гэта распаўсюджаная функцыя
// з пераменнай колькасцю параметраў.

package main

import "fmt"

// Тут функцыя, якая прымае пераменную колькасць
// цэлых лікаў у якасці аргументаў.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// Унутры функцыі тып `nums` эквівалентны
	// `[]int`. Мы можам выклікаць `len(nums)`,
	// перабіраць яго з `range` і г.д.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Функцыі з пераменнай колькасцю параметраў можна выклікаць
	// звычайным спосабам.
	sum(1, 2)
	sum(1, 2, 3)

	// Калі ў вас ужо ёсць некалькі аргументаў у слайсе,
	// застасуўце іх да функцыі са зменнай колькасцю параметрыў, наступным чынам
	// `func(зрэз...)`.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
