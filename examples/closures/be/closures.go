// Го падтрымлівае [_ананімныя функцыі_](https://en.wikipedia.org/wiki/Anonymous_function),
// якая ёсць формай <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>замыканняў</em></a>.
// Ананімны функцыі карсыня калі трэба стварыць функцыю на месцы
// без імені.

package main

import "fmt"

// Гэтая функцыя `intSeq` вяртае іншую функцыю, якую
// ствараем у целе `intSeq`. Вярнутая функцыя _замыкае_
// зменную `i`, і стварае замыканне.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Мы выклікаем `intSeq`, прысвойваючы вынік (функцыю)
	// зменнай `nextInt`. Гэта значэнне функцыі мае
	// уласнае значэнне `i`, якое будзе абнаўляцца кожны раз
	// калі выклікаецца `nextInt`.
	nextInt := intSeq()

	// Паглядзіце на эфект замыкання, выклікаўшы `nextInt`
	// некалькі разоў.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Каб пацвердзіць, што стан унікальны для гэтай
	// канкрэтнай функцыі, стварыце і пратэстуйце яшчэ адну функцыю
	// ў зменную newInts
	newInts := intSeq()
	fmt.Println(newInts())
}
