// Запіс файлаў у Go выконваецца па падобных шаблонах,
// як і тыя, што мы бачылі раней для чытання.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Для пачатку, вось як запісаць тэкст або байты ў файл.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// Для паступовака запісу, спачатку адкрыйце файл для запісу.
	f, err := os.Create("/tmp/dat2")
	check(err)

	// Добрай практыка указаць адкладзеннае закрыцця файлу адрузу пасля адкрыцця.
	defer f.Close()

	// `Write` запісвае байтавыя фрагменты ў файл.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// Таксама даступны `WriteString`.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Выканайце каманду `Sync` запісу дадзеных на дыск.
	f.Sync()

	// `bufio` забяспечвае буферызаваныя запісы ў дадатак да буферызаваных
	// чытачоў, якія мы разглядалі раней.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Выкарыстоўвайце `Flush`, каб пераканацца, што ўсе аперацыі з
	// буферызацыяй былі ўжытыя.
	w.Flush()

}
