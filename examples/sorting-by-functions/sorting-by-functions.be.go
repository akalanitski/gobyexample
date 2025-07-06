// Часам трэба адсартаваць калекцыю па чымсьці іншым,
// акрамя натуральнага парадку. Напрыклад, выкажам здагадку, што мы
// хочам адсартаваць тэкст па даўжыні, а не па алфавіце.
// Вось прыклад сартавання у Go.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// Мы рэалізуем функцыю параўнання тэкстаў за даўжынёй
	// `cmp.Compare` карысны для гэтага.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Цяпер мы можам выклікаць `slices.SortFunc` з нашай
	// параўнання для сартавання `фруктаў` па даўжыні назвы.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// Мы можам выкарыстоўваць той жа метад для сартавання часткі значэнняў,
	// якія не з'яўляюцца ўбудаванымі тыпамі.
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	// Сартаваць `людзей` па ўзросце з дапамогай `slices.SortFunc`.
	//
	// Заўвага: калі структура `Person` вялікая,
	// магчыма, вам захочацца, каб зрэз утрымліваў `*Person` замест гэтага
	// і адпаведна адрэгуляваць функцыю сартавання. Калі ёсць
	// сумневы, [benchmark](testing-and-benchmarking)!
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
