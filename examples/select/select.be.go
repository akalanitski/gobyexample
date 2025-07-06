// Функцыя _select_ у Go дазваляе чакаць некалькіх каналаў
// Спалучэнне горуцін, каналаў і select — магутная інтсрументы Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// У нашым прыкладзе мы будзем выбіраць з двух каналах.
	c1 := make(chan string)
	c2 := make(chan string)

	// Кожны канал атрымае значэнне праз пэўны час,
	// каб імітаваць, напрыклад, блакіроўку аперацый RPC
	// якія выконваюцца ў паралельных горуцінах.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// `select` выкарыстоўваецца для чакання абодвух значэнняў
	// адначасова, друкуючы кожнае з іх па меры яго паступлення.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
