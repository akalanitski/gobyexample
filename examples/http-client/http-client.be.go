// Стандартная бібліятэка Go мае выдатную падтрымку HTTP-кліентаў і сервераў
// у пакеце `net/http`
// Прыклад ніжэй пакажа выкананне простых HTTP-запытаў.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Адбравіць HTTP GET-запыт на сервер.
	// `http.Get` — гэта зручны спосаб стварэння аб'екта `http.Client`
	// і выкліку яго метаду `Get`;
	// ён выкарыстоўвае аб'ект `http.DefaultClient`,
	//які мае карысныя налады па змаўчанні.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Вывесці статус HTTP-адказу.
	fmt.Println("Response status:", resp.Status)

	// Вывесці першыя 5 радкоў цела адказу.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
