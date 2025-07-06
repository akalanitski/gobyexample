// Можна выкарыстоўваць уласныя тыпы як `error` шляхам
// рэалізацыі на іх метаду `Error()`. Прыклад ніжэй
// выкарыстоўвае ўласны тып для відавочнага прадстаўлення
// памылкі агрумента.

package main

import (
	"errors"
	"fmt"
)

// Выкарыстанне суфікс "Error" добрая практыка для ўласных памылак
type argError struct {
	arg     int
	message string
}

// Даданне гэтага метаду `Error` дазваляе `argError` рэалізаваць
// інтэрфейс `error`.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// Павяртае памылку ўласнага тыпу
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// `errors.As` — гэта больш пашыраная версія `errors.Is`.
	// Правярае, ці адпавядае зададзеная памылка (ці любая памылка ў яе ланцужку)
	// пэўнаму тыпу памылкі і пераўтварае яе ў значэнне
	// гэтага тыпу, вяртаючы `true`. Калі супадзення няма, яна
	// вяртае `false`.
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
