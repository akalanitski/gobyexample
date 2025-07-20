// Некаторыя праграммы з тэкставым інтэрфейсам, такія як `go` або `git`
// маюць шмат *падкаманд*, кожная з якіх мае свой уласны набор
// параметраў. Напрыклад, `go build` і `go get` — гэта дзве
// розныя падкаманды асноўнай праграмы `go`.
// Пакет `flag` дазваляе нам лёгка вызначаць
// падкаманды, якія маюць свае ўласныя параметры.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Мы аб'яўляем падкаманду з функцыяй `NewFlagSet`
	// і працягваем вызначаць новыя параметры, для гэтай падкаманды.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Для розных падкаманд мы можам вызначыць розныя
	// параметры.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Падкаманда павінна быць першым аргументам
	// праграмы.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Праверка таго, якая падкаманда выклікаецца.
	switch os.Args[1] {

	// Для кожнай падкаманды мы аналізуем яе параметры і
	// маем доступ да канчатковых пазіцыйных аргументаў.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
