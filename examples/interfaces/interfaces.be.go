// _Інтэрфейс_ — гэта іменаваныя калекцыя сігнатур метадаў.

package main

import (
	"fmt"
	"math"
)

// Тут просты інтэрфейс для геаметрычных фігур.
type geometry interface {
	area() float64
	perim() float64
}

// У нашым прыкладзе мы рэалізуем гэты інтэрфейс на тыпах `rect` і `circle`.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Каб рэалізаваць інтэрфейс у Go, нам проста трэба
// рэалізаваць усе метады ў інтэрфейсе. Тут мы
// рэалізуем `geometry` на `rect`s`.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Рэалізацыя для `circle`
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Калі зменная мае тып інтэрфейсу, то мы можам выклікаць
// метады, якія знаходзяцца ў названым інтэрфейсе. Вось
// агульная функцыя `measure`, якая выкарыстоўвае гэта
// для працы з любой `геаметрыяй`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Часам карысна ведаць тып выканання значэння інтэрфейсу.
// Адзін з варыянтаў — выкарыстоўваць *сцвярджэнне тыпу*
// як паказана тут; іншы — [тыпу `switch`](switch.html).
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Структуры тыпаў `circle` і `rect`
	// рэалізуюць інтэрфейс `geometry`, каб мы маглі выкарыстоўваць
	// экзэмпляры гэтых структур як аргументы для `measure`.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
