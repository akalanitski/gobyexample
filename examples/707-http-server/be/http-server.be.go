// Напісаць просты HTTP-сервер лёгка з дапамогай пакета
// `net/http`.
package main

import (
	"fmt"
	"net/http"
)

// Фундаментальная канцэпцыя сервераў `net/http` — гэта
// *апрацоўшчыкі*. Апрацоўшчык — гэта аб'ект, які рэалізуе
// інтэрфейс `http.Handler`. Распаўсюджаны спосаб запісу
// апрацоўшчыка — выкарыстанне адаптара `http.HandlerFunc`
// на функцыях з адпаведнай сігнатурай.
func hello(w http.ResponseWriter, req *http.Request) {

	// Функцыі, якія служаць апрацоўшчыкамі, прымаюць
	// `http.ResponseWriter` і `http.Request` як
	// аргументы. `ResponseWriter` выкарыстоўваецца для запаўнення
	// HTTP-адказу. Тут наш просты адказ — гэта проста
	// "hello\n".
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// Гэты апрацоўшчык робіць нешта больш складанае,
	// чытаючы ўсе загалоўкі HTTP-запыту і перадаючы іх у цела адказу.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// Мы рэгіструем нашы апрацоўшчыкі на маршрутах сервера з дапамогай
	// функцыі `http.HandleFunc`. Яна ўсталёўвае
	// *маршрутызатар па змаўчанні* ў пакеце `net/http` і
	// прымае функцыю ў якасці аргумента.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Нарэшце, мы выклікаем `ListenAndServe` з портам
	// і апрацоўшчыкам. `nil` паведамляе яму выкарыстоўваць маршрутызатар па змаўчанні
	// які мы толькі што наладзілі.
	http.ListenAndServe(":8090", nil)
}
