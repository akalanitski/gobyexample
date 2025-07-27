// _Defer_ выкарыстоўваецца для таго, каб выканаць код пасля выканання функцыі
// ў якой эн быў напісаны. Гэта зручна для мэтаў ачысткі.
// `defer` часта выкарыстоўваецца там, дзе, у іншых мовах выкарыстоўваюцца
// `ensure` і `finally`

package main

import (
	"fmt"
	"os"
)

// Дапусцім, мы хочам стварыць файл, запісаць у яго дадзеныя,
// а потым закрыць яго пасля завяршэння. Вось як мы маглі б
// зрабіць гэта з дапамогай `defer`.
func main() {

	// Адразу пасля атрымання аб'екта файла з дапамогай
	// `createFile`, мы адкладаем закрыццё гэтага файла
	// з дапамогай `closeFile`. Гэта будзе выканана ў канцы
	// функцыі, якая ахоплівае файл (`main`), пасля
	// завяршэння `writeFile`.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// Важна правяраць наяўнасць памылак пры закрыцці
	// файла, нават у адкладзенай функцыі.
	if err != nil {
		panic(err)
	}
}
