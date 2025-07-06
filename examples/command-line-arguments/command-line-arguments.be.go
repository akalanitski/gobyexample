// [_Аргументы каманднага радка_](https://be.wikipedia.org/wiki/Інтэрфейс_каманднага_радка)
// з'яўляюцца распаўсюджаным спосабам параметрызацыі выканання праграм.
// Напрыклад, `go run hello.go` выкарыстоўвае `run` і
// `hello.go` як аргументы для праграмы `go`.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` забяспечвае доступ да аргументаў каманднага радка
	// Звярніце ўвагу, што першае значэнне ў гэтым зрэзе
	// — гэта шлях да праграмы, а `os.Args[1:]` змяшчае аргументы праграмы.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// Атрымаць асобныя аргументы можна па індэксу.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
