// Go падтрымлівае _ўбудоўванне_ структур і інтэрфейсаў
// для больш галдкай _кампазіцыяй_ тыпаў.
// Не варта блытаць з [`//go:embed`](embed-directive), якая з'яўляецца
// дырэктывай камплітара go, уведзенай у Go версіі 1.16+ для ўбудоўвання
// файлаў і папак у двайковы файл праграмы.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// `container` _ўбудоўвае_ `базу`. Убудаванне выглядае
// як поле без назвы.
type container struct {
	base
	str string
}

func main() {

	// Пры стварэнні структур з літэраламі мы павінны яўна
	// задаць значэнні убудаванай структуры; тут убудаваны тып служыць імем поля.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// Але выкарыстоўваць значэнні убудаванай структуры можам без указання яе назвы
	// напрыклад, `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Пры гэтым можно прапісаць і поўны шлях, выкарыстоўваючы
	// назву ўбудаванай структуры.
	fmt.Println("also num:", co.base.num)

	// Паколькі `container` умяшчае `base`, метады
	// `base` таксама становяцца метадамі `container`. Тут
	// мы выклікаем метад, які быў убудаваны з `base`
	// непасрэдна ў `co`.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Убудаванне структур з метадамі можа выкарыстоўвацца для перадачы
	// рэалізацый інтэрфейсаў іншым структурам. Тут
	// мы бачым, што `container` цяпер рэалізуе
	// інтэрфейс `describer`, таму што ён убудоўвае `base`.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
