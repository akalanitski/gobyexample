// Пакет `filepath` прадастаўляе функцыі для разбору
// і пабудовы *шляхоў да файлаў* у выглядзе сумесным за рознымі сістэмамі;
// напрыклад, `dir/file` у Linux і `dir\file` у Windows.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` варта выкарыстоўваць для пабудовы шляхоў.
	// Ён прымае любую колькасць аргументаў і будуе з іх шлях.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Варта заўсёды выкарыстоўваць `Join` замест аб'ядноўваць тэксту з
	// `/` або `\`, Акрамя таго, `Join` таксама
	// нармалізуе шляхі, выдаляючы лішнія раздзяляльнікі
	// і змены каталогаў.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` і `Base` можна выкарыстоўваць для падзелу шляху да
	// каталога і файла.
	// Акрамя таго, `Split` верне абодва ў адным выкліку.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Можна праверыць ці з'яўляецца шлях абсалютным.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Некаторыя імёны файлаў маюць пашырэнні пасля кропкі. Мы
	// можам аддзяліць пашырэнне ад такіх імёнаў з дапамогай `Ext`.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Каб атрымаць назву файла без пашырэння,
	// выкарыстоўвайце `strings.TrimSuffix`.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` знаходзіць адносны шлях паміж *базавым* і
	// *мэтай*. Вяртае памылку, калі мэта не можа
	// быць адносным да базы.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
