// Пры выкарыстанні каналаў у якасці параметраў функцыі можна
// паказаць, ці прызначаны канал толькі для адпраўкі ці толькі
// для атрымання значэнняў. Такая дакладнасць павышае бяспеку тыпаў
// (type-safety) праграма.

package main

import "fmt"

// Функцыя `ping` прымае канал толькі для адпраўкі.
// Калі перадаць у яе канал прыёмнік, будзе памылка кампілыцыі
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Функцыя `pong` прымае адзін канал-прыёмнік
// (`pings`) і другі для адпраўкі (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
