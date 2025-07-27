// Галінаванне з дапамогай `if` і `else` у Go з'яўляецца простым.

package main

import "fmt"

func main() {

	// Вось просты прыклад.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Можна выкарыстоўваць аператар `if` без `else`.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// часта ва ўмовах выкарыстоўваюцца лагічныя аператары такія як `&&` і `||`
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// У блогу з умовай можа апісаць зменныя якію будуць даступны ва ўсех галінах
	// канструкцыі `if-else`
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Звярніце ўвагу, што ў Go не патрэбныя дужкі вакол умоў,
// але фігурныя дужкі абавязковыя.
