// У папярэднім прыкладзе мы разглядалі наладу простага
// [HTTP-сервера](http-server). HTTP-серверы карысныя для
// дэманстрацыі выкарыстання `context.Context` для
// кіравання адменай. `Кантэкст` змяшчае тэрміны,
// сігналы адмены і іншыя значэнні, абмежаваныя запытам
// праз межы API і горуцін.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// Кантэкст `context.Context` ствараецца для кожнага запыту з дапамогай
	// механізму `net/http` і даступны з праз метад `Context()`.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Пачакайце некалькі секунд, перш чым адправіць адказ кліенту.
	// Гэта можа імітаваць нейкую працу, якую выконвае сервер.
	// Падчас працы сачыце за каналам кантэксту
	// `Done()` на наяўнасць сігналу аб тым, што нам трэба адмяніць
	// працу і вярнуцца як мага хутчэй.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// Метад `Err()` кантэксту вяртае памылку
		// якая тлумачыць, чаму канал `Done()` быў
		// закрыты.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// Як і раней, мы рэгіструем наш апрацоўшчык на маршруце "/hello"
	// і пачынаем абслугоўванне.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
