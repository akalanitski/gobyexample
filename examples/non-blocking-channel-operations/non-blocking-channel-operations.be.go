// Звычайныя адпраўкі і прыёмы па каналах блакуюцца.
// Аднак мы можам выкарыстоўваць `select` з пунктам `default` для
// рэалізацыі _неблакуючых_ адпраўак, прыёмаў і нават
// неблакуючых шматбаковых `select`.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Вось неблакіруючы прыём. Калі значэнне даступнае для `messages`, то `select` прыме
	// `<-messages` `case` з гэтым значэннем. Калі не,
	// ён адразу прыме `default` case.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Неблакіруючая адпраўка працуе аналагічна. Тут `msg`
	// нельга адправіць у канал `messages`, таму што
	// у канала няма буфера і няма атрымальніка.
	// Таму абраны выпадак `default`.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Мы можам выкарыстоўваць некалькі `case` над `default`
	// пунктам для рэалізацыі шматбаковага неблакіруючага
	// выбару. Тут мы спрабуем неблакіруючыя прыёмы
	// для `messages` і `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
