// [_Абмежаванне хуткасці_](https://en.wikipedia.org/wiki/Rate_limiting)
// з'яўляецца важным механізмам для кантролю выкарыстання рэсурсаў
// і падтрымання якасці абслугоўвання.
// Go элегантна падтрымлівае абмежаванне хуткасці з дапамогай goroutines,
// каналаў і [цікераў](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Спачатку мы разгледзім базавае абмежаванне хуткасці.
	// Уявім, што трэба абмежаваць апрацоўку ўваходных запытаў.
	// Гэтыя запыты абслугоўваюцца з канала з назвай `requests`.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Канал `limiter` будзе атрымліваць значэнне кожныя 200 мілісекунд.
	// Гэта канал-рэгулятар у нашай схеме абмежавання хуткасці.
	limiter := time.Tick(200 * time.Millisecond)

	// Блакуючы атрыманне з канала `limiter`
	// перад абслугоўваннем кожнага запыту, мы абмяжоўваем сябе
	// 1 запытам кожныя 200 мілісекунд.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Мы можам дазволіць кароткія пакеты запытаў у
	// нашай схеме абмежавання хуткасці, захоўваючы пры гэтым
	// агульны ліміт хуткасці. Мы можам дасягнуць гэтага шляхам
	// буферызацыі нашага канала абмежавальніка. Гэты канал `burstyLimiter`
	// дазволіць пакеты да 3 падзей.
	burstyLimiter := make(chan time.Time, 3)

	// Запаўняем канал, каб прадставіць дазволены пакетны набор дадзеных.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Кожныя 200 мілісекунд мы будзем спрабаваць дадаць новае
	// значэнне да `burstyLimiter`, да яго ліміту ў 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Цяпер прамадэлюйце яшчэ 5 уваходных запытаў. Першыя
	// 3 з іх атрымаюць перавагу ад магчымасці пакетнай апрацоўкі
	// `burstyLimiter`.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
