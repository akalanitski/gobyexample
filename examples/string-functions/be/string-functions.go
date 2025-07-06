// Пакет `strings` стандартнай бібліятэкі забяспечвае шмат
// карысных функцый, звязаных з тэкстам. Вось некалькі прыкладаў,
// каб даць вам уяўленне аб пакеце.

package main

import (
	"fmt"
	s "strings"
)

// Мы выкарыстоўваем аліас для `fmt.Println` бо будзем выкарыстоўваць
// яго шмат разоў ніжэй.
var p = fmt.Println

func main() {

	// Вось прыклад функцый, даступных у
	// `strings`. Паколькі гэта функцыі з
	// пакета, а не метады самога аб'екта string,
	// нам трэба перадаць адпаведны радок у якасці першага
	// аргумента функцыі. Больш функцый можна знайсці ў дакументацыі пакета
	// [`strings`](https://pkg.go.dev/strings) дакументацыі пакета.
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
