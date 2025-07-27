// У Go выкарыстоўвае С-like або Unix-like стыл паведамлення пра памылкі --
// errorno які павяртаюць усе фукцыі. Гэта кантрастуе з
// выключэннямі (Exceptions), якія выкарыстоўваюцца ў такіх мовах, як Java і
// Ruby. GO  палепшыў гэты прыём дадаўшы функцыям магчымасць павяртаць шмат
// рэзультатаў. Падыход Go дазваляе лёгка ўбачыць, якія
// функцыі вяртаюць памылкі, і апрацоўваць іх з дапамогай
// тых жа моўных канструкцый, якія выкарыстоўваюцца для іншых,
// задач без памылак.
//
// Дакументацыя [errors](https://pkg.go.dev/errors)
// і [гэты пост у блогу](https://go.dev/blog/go1.13-errors) для атрымання дадатковай
// інфармацыі.

package main

import (
	"errors"
	"fmt"
)

// Паводле дамоўленасці, памылкі з'яўляюцца апошнім вяртаным значэннем і
// маюць тып `error`, убудаваны інтэрфейс.
func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` стварае базавае значэнне `error`
		// з паведамленнем пра памылку.
		return -1, errors.New("can't work with 42")
	}

	// Значэнне `nil` на месцы памылкі паказвае, што
	// памылкі не было.
	return arg + 3, nil
}

// Памылка-кантроль — гэта загадзя абвешчаная зменная, якая выкарыстоўваецца для
// абазначэння паўнай умовы памылкі.
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// Мы можам камбінаваць памылкі памылкамі, каб дадаць кантэкст.
		// Найпрасцейшы спосаб зрабіць гэта — з дапамогай дзеяслова `%w` у `fmt.Errorf`.
		// Абгарнуць памылкі стварыць лагічны ланцужок (A абгортвае B, якая абгортвае C і г.д.)
		// які можна запытаць з дапамогай функцый, такіх як `errors.Is`
		// і `errors.As`.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// Звычайна выкарыстоўваецца ўбудаваная праверка на памылкі
		// канструкцыяй галінавання
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			// `errors.Is` правярае, ці адпавядае дадзеная памылка (ці любая памылка ў яе ланцужку)
			// пэўнаму значэнню памылкі. Гэта асабліва карысна для скамбінаваных,
			// што дазваляе ідэнтыфікаваць пэўныя тыпы памылак ў ланцужку.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
