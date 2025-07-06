// URL-адрасы забяспечваюць
// [уніфікаваны спосаб пошуку рэсурсаў](https://be.wikipedia.org/wiki/URL).
// Вось як разабраць URL-адрасы ў Go.

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// Мы разбярэм гэты прыклад URL-адраса, які ўключае ў сябе
	// схему, інфармацыю для аўтэнтыфікацыі, host, port, шлях,
	// параметры запыту і фрагмент запыту.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Разбіраем URL і пераконваемся ў адсутнасці памылак.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Доступ да схемы просты.
	fmt.Println(u.Scheme)

	// `User` змяшчае ўсю інфармацыю для аўтэнтыфікацыі; выклічце
	// `Username` і `Password` для асобных значэнняў.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` змяшчае як імя хаста, так і порт,
	// калі ёсць. Выкарыстоўвайце `SplitHostPort` для іх здабывання.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Тут мы здабываем `шлях` і фрагмент пасля сімвала `#`.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// Выкарыстоўвайце `RawQuery`, каб атрымаць параметры запыту ў
	// выглядзе тэкставых пар `k=v`.
	// Вы таксама можаце разабраць параметры запыту
	// ў мапу. Разабраныя мапа параметраў запыту дазваляе атрымаць
	// Значэнне па ключу. Але павяртае значенне параметра як слайс,
	// таму да яго трэба звяртацца праз жорсткія дужкі `[` і `]`
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
