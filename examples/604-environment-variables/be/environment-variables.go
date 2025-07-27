// [Зменныя асяроддзя](https://en.wikipedia.org/wiki/Environment_variable)
// з'яўляюцца універсальным механізмам для [перадачы канфігурацыйнай
// інфармацыі праграмам Unix](https://www.12factor.net/config).
// Давайце паглядзім, як усталёўваць, атрымліваць і пералічваць зменныя асяроддзя.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Каб задаць пару ключ/значэнне, выкарыстоўвайце `os.Setenv`.
	// Каб атрымаць значэнне па ключу, выкарыстоўвайце `os.Getenv`.
	// Калі ключ зменная адсутнічыае у асяродзі `os.Getenv` павярне
	// нулявое значэнне
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Фнукцыя `os.Environ` павяртая слайс тэкставых зменных, кожная апісвае пару
	// "ключ=значэнне" зменнай асяродзя. Ніжэй функцыя `strings.SplitN` разбівае
	// пары на асобныя значэнні і выводзіць толькі ключы
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
