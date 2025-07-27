// Разбор лікаў з радкоў — гэта простая, але распаўсюджаная задача
// у многіх праграмах; вось як гэта зрабіць у Go.

package main

// Убудаваны пакет `strconv` забяспечвае лік
// разбор.
import (
	"fmt"
	"strconv"
)

func main() {

	// З дапамогай `ParseFloat`, гэта `64` паказвае, колькі біт
	// дакладнасці трэба разабраць.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// Для `ParseInt` `0` азначае вывесці аснову з
	// радка. `64` патрабуе, каб вынік змяшчаўся ў 64
	// бітах.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` распазнае лікі ў шаснаццатковым фармаце.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// Таксама даступны `ParseUint`.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` — гэта зручная функцыя для базавага дзесятковага ліку
	// разбору `int`.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Функцыі разбору вяртаюць памылку пры няправільным уводзе.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
