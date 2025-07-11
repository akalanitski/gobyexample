// [Таймеры](таймеры) прызначаны для адзіночных выпадкаў у будучыні
// _тыкеры_ прызначаны для шматразовых (рэгулярных) падзей з роўным
// інтэрвалам часу. Вось прыклад тыкера, які тыкае
// перыядычна, пакуль яго не спыніць.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Тыкеры выкарыстоўваюць падобны механізм да таймераў:
	// канал, у які адпраўляюцца значэнні. Тут мы будзем выкарыстоўваць
	// убудаваную канструкцыю `select` у канале, каб чакаць
	// значэнняў па меры іх паступлення кожныя 500 мс.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Тыкеры можна спыняць, як таймеры. Пасля таго, як тыкер спынены,
	// ён больш не будзе атрымліваць значэнні у свай канал.
	// Мы спынім наш праз 1600 мс.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
