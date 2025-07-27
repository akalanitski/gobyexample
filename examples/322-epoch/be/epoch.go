// Час у праграммах часта адлічваецца 1 студзеня 1970г
// (так званная дата пачатку [зпохі Unix](https://be.wikipedia.org/wiki/Час_Unix).
// Гэта дазваляе зберагаць час у лічбавым фармаце наўпрост як беззнакавая лічба
// Даты і час разрахоўваюцца ўбудованнымі функцыямі.
// Такі падыходы да працы з часам таксама называецца Час Unix.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Выкарыстоўвайце `time.Now` з `Unix`, `UnixMilli` або `UnixNano`
	// каб атрымаць час, які прайшоў з эпохі Unix, у секундах,
	// мілісекундах або нанасекундах адпаведна.
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// Вы таксама можаце пераўтварыць цэлыя секунды або нанасекунды
	// з эпохі ў адпаведны `час`.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
