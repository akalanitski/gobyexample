// Пачынаючы з версіі 1.18, Go дадала падтрымку
// _абагульненняў_, таксама вядомых як _параметры тыпу_.

package main

import "fmt"

// Напрыклад функцыя, `SlicesIndex` прымае слайс тыпу `comparable`
// і параметр гэтага ж тыпу і вяртае індэкс першага ўваходжання
// v у s, або -1, калі адсутнічае. Абмежаванне `comparable`
// азначае, што мы можам параўноўваць значэнні гэтага тыпу з дапамогай
// аператараў `==` і `!=`. Для больш падрабязнага тлумачэння
// гэтай сігнатуры тыпу глядзіце [гэты пост у блогу](https://go.dev/blog/deconstructing-type-parameters).
// Звярніце ўвагу, што гэтая функцыя існуе ў стандартнай бібліятэцы
// як [slices.Index](https://pkg.go.dev/slices#Index).
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// Прыклад агульнага тыпу, `List` — гэта
// адназвязаны спіс са значэннямі любога тыпу.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Мы можам вызначаць метады для агульных тыпаў гэтак жа, як мы
// робім для звычайных тыпаў, але мы павінны захоўваць параметры тыпу
// на месцы. Тып — `List[T]`, а не `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements вяртае ўсе элементы спісу ў выглядзе слайсу.
// У наступным прыкладзе мы ўбачым больш тыповы спосаб
// перабору ўсіх элементаў карыстальніцкіх тыпаў.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// Пры выкліку дженекік функцый мы часта можам спадзявацца
	// на _вывад тыпаў_. Звярніце ўвагу, што нам не трэба
	// паказваць тыпы для `S` і `E` пры
	// выкліку `SlicesIndex` - кампілятар выводзіць іх
	// аўтаматычна.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// ... але можна ўказаць і дакладныя тыпы
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
