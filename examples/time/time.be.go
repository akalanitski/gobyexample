// Go прапануе шырокую падтрымку часу і працягласці;
// вось некалькі прыкладаў.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Пачнем з атрымання часу.
	now := time.Now()
	p(now)

	// Вы можаце стварыць структуру `time`, указаўшы
	// год, месяц, дзень і г.д. Час заўсёды звязаны
	// з месцам (`Location`) а г.зн. з часавым поясам.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Вы можаце атрымаць розныя кампаненты часу наступным чынам
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Таксама даступны `Weekday` з панядзелка па нядзелю.
	p(then.Weekday())

	// Гэтыя метады параўноўваюць два разы, правяраючы, ці адбываецца
	// першае падзея да, пасля або адначасова з другім адпаведна.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Метады `Sub` вяртаюць працягласць (`Duration`), які прадстаўляе
	// розніцу паміж двума момантамі часу.
	diff := now.Sub(then)
	p(diff)

	// Мы можам вылічыць працягласць у
	// розных адзінках вымярэння.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Вы можаце выкарыстоўваць `Add`, каб перасунуць час наперад на зададзены
	// час, або з дапамогай `-`, каб перасунуць назад на час.
	p(then.Add(diff))
	p(then.Add(-diff))
}
