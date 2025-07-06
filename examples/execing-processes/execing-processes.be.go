// У папярэднім прыкладзе мы разглядалі [стварэнне працэсаў](spawning-processes.html).
// Гэта зручка калі патрэбен працэс, даступны для запушчанага працэсу Go.
// Часам мы проста хочам цалкам замяніць працэс Go іншым (магчыма, не Go).
// Для гэтага мы будзем выкарыстоўваць рэалізацыю функцыі Go
// HTML:
// <a href="https://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// MD:
// [`exec`](https://en.wikipedia.org/wiki/Exec_(operating_system))
// функцыі.

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// У нашым прыкладзе выконваецца `ls`.
	// Go патрабуе абсалютны шлях да двайковага файла, які мы хочам выканаць,
	// таму мы будзем выкарыстоўваць `exec.LookPath`, каб знайсці яго (верагодна,
	// `/bin/ls`).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` патрабуе аргументаў у выглядзе слайсу (не адзін доўга радок).
	// Мы дамо `ls` некалькі распаўсюджаных аргументаў.
	// Звярніце ўвагу, што першы аргумент павінен быць назвай праграмы.
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` таксама патрабуе набор [зменных асяроддзя](environment-variables.html)
	// для выкарыстання. Тут мы проста падаем асяродзе без змен
	env := os.Environ()

	// Вось фактычны выклік `syscall.Exec`.
	// Паспяховы выклік прыпыніц выкананне гэтай праграммы і запусціць замест яе
	// новы новую (`/bin/ls -a -l -h`)
	// У выпадку памылкі, мы атрымаем памылку ў `execErr`.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
