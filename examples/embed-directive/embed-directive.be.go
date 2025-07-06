// `//go:embed` — гэта [дырэктыва кампілятара](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives),
// якая дазваляе ўключаць у праграму змест файла або дырыкторыі падчас зборкі.
// Больш падрабязна пра дырэктыву embed чытайце [тут](https://pkg.go.dev/embed).
package main

// Імпартуйце пакет `embed`; калі вы не выкарыстоўваеце ніякіх экспартаваных
// ідэнтыфікатараў з гэтага пакета, вы можаце зрабіць пусты імпарт праз `_ "embed"`.
import (
	"embed"
)

// дырэктывы `embed` прымаюць шляхі адносна каталога, які змяшчае
// файл зыходнага кода Go. Гэтая дырэктыва ўбудоўвае змесціва файла ў
// зменную `fileString`, якая ідзе адразу за ім.
//
//go:embed folder/single_file.txt
var fileString string

// Або ўбудаваць змесціва файла ў `[]байт`.
//
//go:embed folder/single_file.txt
var fileByte []byte

// Можам убудоўваць некалькі файлаў ці нават каталог праз падстаноўныя
// знакі. Для гэтага выкарыстоўваецца зменная тыпу
// [embed.FS](https://pkg.go.dev/embed#FS), якая
// рэалізуе простую віртуальную файлавую сістэму.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// Вывесці змесціва `single_file.txt`.
	print(fileString)
	print(string(fileByte))

	// Атрымайце некалькі файлаў з убудаванага каталога.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
