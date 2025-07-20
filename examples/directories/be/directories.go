// Go мае некалькі карысных функцый для працы з
// *каталогамі* файлавай сістэмы.

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Стварыць падкаталог у працоўным каталогу.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Пры стварэнні тымчасовых каталогаў добра
	// практыкавацца ў тым, каб адкласці іх выдаленне.
	// `os.RemoveAll` выдаліць усё дрэва каталогаў (падобна да
	// `rm -rf`).
	defer os.RemoveAll("subdir")

	// Функцыя для стварэння пустога файла.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Мы можам стварыць іерархію каталогаў, у тым ліку
	// над-каталогаў з дапамогай `MkdirAll`. Гэта падобна
	// да каманды `mkdir -p`.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` выводзіць змест каталога, вяртаючы
	// слайс з элементамі `os.DirEntry`.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` дазваляе змяніць працоўны каталог,
	// падобна да каманды `cd`.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Цяпер мы ўбачым змесціва каталога `subdir/parent/child`
	// у які перайшлі вышэй
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `cd` назад да таго месца, адкуль мы пачалі.
	err = os.Chdir("../../..")
	check(err)

	// Мы таксама можам наведаць каталог *рэкурсіўна*,
	// уключаючы ўсе яго падкаталогі. `WalkDir` прымае
	// функцыю зваротнага выкліку для апрацоўкі кожнага файла або
	// наведанага каталога.
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

// `visit` выклікаецца для кожнага знойдзенага файла або каталога
// рэкурсіўна з дапамогай `filepath.WalkDir`.
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
