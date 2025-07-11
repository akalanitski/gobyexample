// Go мае ўбудаваную падтрымку _некалькіх вяртаемых значэнняў_.
// Гэтая функцыя часта выкарыстоўваецца ў ідыяматычным Go, напрыклад
// для вяртання як выніку, так і значэння памылкі з функцыі.

package main

import "fmt"

// `(int, int)` у гэтай сігнатуры функцыі паказвае, што
// функцыя вяртае 2 `in`.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Тут мы выкарыстоўваем 2 розныя значэнні, якія вяртаюцца з
	// выкліку з _множным прысваеннем_.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Калі патрэбна толькі адзін з вернутых значэнняў,
	// выкарыстоўвайце пусты ідэнтыфікатар `_`.
	_, c := vals()
	fmt.Println(c)
}
