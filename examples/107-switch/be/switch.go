// Канструкцыя switch (перакючэння) разгаліноўвае логіку ў багацёх кірунках

package main

import (
	"fmt"
	"time"
)

func main() {

	// Тут просты `switch`.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Праз коску, можна пералічыць некалькіх умоў аданаго выпадку (`case`).
	// У гэтым прыкладзе таксама выкарыстоўваецца безумоўны-выпадак (`default` case).
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Канстукцыя `switch` без выразу — гэта яшчэ адзін спосаб
	// запісу if/else які мая шмат умоў выбару.
	// Тут яшчэ бычым, умову замест канстанты перад  `case`.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Канструкцыя `type switch`) параўноўвае тыпы замест значэнняў.
	// Гэма можна выкарыстоўваць для вызначэння тыпу інтэрфейснага
	// значэння. У гэтым прыкладзе зменная `t` будзе мець
	// тып, які адпавядае яе блоку.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
