// Go падтрымлівае _метады_, вызначаныя ў тыпах структур.

package main

import "fmt"

type rect struct {
	width, height int
}

// Гэты метад `area` мае тып _прыёмніка_ `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Метады могуць быць вызначаны як для ўказальнікаў, так і для значэнняў
// тыпаў прыёмнікаў. Вось прыклад прыёмніка значэння.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Тут мы выклікаем 2 метады, вызначаныя для нашай структуры.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go аўтаматычна апрацоўвае пераўтварэнне паміж значэннямі
	// і ўказальнікамі для выклікаў метадаў. Вы можаце выкарыстоўваць
	// тып прыёмніка ўказальніка, каб пазбегнуць капіявання пры выкліках метаду
	// або дазволіць метаду змяняць
	// структуру прыёму.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
