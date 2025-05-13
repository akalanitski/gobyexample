// Пачынаючы з версіі 1.23, у Go дададзена падтрымка
// [ітэратараў](https://go.dev/blog/range-functions),
// што дазваляе нам ахопліваць практычна ўсё!

package main

import (
	"fmt"
	"iter"
	"slices"
)

// Давайце зноў паглядзім на тып `List` з
// [папярэдняга прыкладу](generics). У гэтым прыкладзе
// у нас быў метад `AllElements`, які вяртаў слайс
// усіх элементаў у спісе. З дапамогай ітэратараў Go мы
// можам зрабіць гэта лепш.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// Функцыя `All` вяртае _ітэратар_, які ў Go з'яўляецца функцыяй
// са [спецыяльнай сігнатурай](https://pkg.go.dev/iter#Seq).
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// Функцыя-ітэратар прымае іншую функцыю ў якасці
		// параметра, які паводле дамоўленасці называецца `yield` (але
		// назва можа быць любой). Яна будзе выклікаць `yield` для
		// кожнага элемента, які мы хочам перабраць, і натаваць значэнне,
		// якое вяртае `yield`, для патэнцыйнага датэрміновага завяршэння.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Ітэрацыя не патрабуе базавай структуры дадзеных,
// і нават не абавязкова павінна быць канечнай! Вось функцыя,
// якая вяртае ітэратар па ліках Фібаначы: яна працягвае
// выконвацца, пакуль `yield` вяртае `true`.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Паколькі `List.All` вяртае ітэратар, мы можам выкарыстоўваць яго
	// у звычайным цыкле `range`.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// Пакеты, падобныя да [slices](https://pkg.go.dev/slices), маюць
	// шэраг карысных функцый для працы з ітэратарамі.
	// Напрыклад, `Collect` бярэ любы ітэратар і збірае
	// усе яго значэнні ў слайс.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {

		// Пасля таго, як цыкл дасягне `break` або ранняга вяртання (калі функцыя `yield`
		// верне `false`).
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
