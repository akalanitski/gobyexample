// Пакет `slices` у Go рэалізуе сартаванне для ўбудаваных
// і вашых тыпаў. Спачатку мы разгледзім сартаванне для
// ўбудаваных тыпаў.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Функцыі сартавання з'яўляюцца агульнымі і працуюць для любога
	// _ordered_ убудаванага тыпу. Спіс упарадкаваных
	// тыпаў глядзіце ў [cmp.Ordered](https://pkg.go.dev/cmp#Ordered).
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// Прыклад сартавання `цэлых` лікаў.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// Мы таксама можам выкарыстаць пакет `slices`, каб праверыць, ці
	// слайс ужо адсартаваны.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
