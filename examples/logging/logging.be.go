// Стандартная бібліятэка Go забяспечвае простыя
// інструменты для вываду журналаў з праграм Go, з
// пакетам [log](https://pkg.go.dev/log) для
// вываду ў вольнай форме і
// пакетам [log/slog](https://pkg.go.dev/log/slog) для
// структураванага вываду.
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// Проста выклік функцый, такіх як `Println` з пакета
	// `log`, выкарыстоўвае _standard_ logger, які
	// ужо папярэдне настроены для разумнага рэгістравання
	// вываду ў `os.Stderr`. Дадатковыя метады, такія як
	// `Fatal*` або `Panic*`, прывядуць да выхаду з праграмы пасля
	// рэгістравання.
	log.Println("standard logger")

	// Рэгістратары можна наладзіць з дапамогай _flags_ для ўсталёўкі
	// фармату іх вываду. Па змаўчанні стандартны
	// рэгістратар мае ўсталяваныя сцягі `log.Ldate` і `log.Ltime`
	// і яны сабраны ў `log.LstdFlags`.
	// Мы можам змяніць яго сцягі, каб яны выпраменьвалі час з
	// дакладнасцю да мікрасекунд, напрыклад.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// Таксама падтрымліваецца вывад назвы файла і
	// радка, з якога выклікаецца функцыя `log`.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// Можа быць карысна стварыць уласны логер і
	// перадаць яго далей. Пры стварэнні новага логера мы
	// можам усталяваць _прэфікс_, каб адрозніць яго вынік
	// ад іншых логераў.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// Мы можам усталяваць прэфікс
	// на існуючых логерах (у тым ліку на стандартным)
	// з дапамогай метаду `SetPrefix`.
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// Рэгістратары могуць мець уласныя мэты вываду;
	// любы `io.Writer` працуе.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// Гэты выклік запісвае вывад журнала ў `buf`.
	buflog.Println("hello")

	// Гэта сапраўды пакажа гэта на стандартным вывадзе.
	fmt.Print("from buflog:", buf.String())

	// Пакет `slog` забяспечвае
	// _структураваны_ вывад журнала. Напрыклад, рэгістрацыя
	// у фармаце JSON з'яўляецца простай.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// Акрамя паведамлення, вывад `slog` можа
	// утрымліваць адвольную колькасць пар ключ=значэнне
	//.
	myslog.Info("hello again", "key", "val", "age", 25)
}
