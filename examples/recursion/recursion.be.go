// Go падтрымлівае
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>рэкурсіўныя функцыі</em></a>.
// Вось класічны прыклад.

package main

import "fmt"

// Гэтая функцыя `fact` выклікае саму сябе, пакуль не дасягне
// базавага выпадку `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Ананімныя функцыі таксама могуць быць рэкурсіўнымі, але гэта патрабуе
	// відавочнага аб'яўлення зменнай з дапамогай `var` для захоўвання
	// функцыі перад яе вызначэннем.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Паколькі `fib` раней быў аб'яўлены ў `main`, Go
		// ведае, якую функцыю выклікаць з дапамогай `fib`.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
