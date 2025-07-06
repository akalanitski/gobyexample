// Часам нашым праграмам на Go трэба ствараць іншыя працэсы, не звязаныя з Go.

package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// Мы пачнем з простай каманды, якая не прымае ніякіх
	// аргументаў або ўваходных дадзеных і проста выводзіць нешта на
	// стандартны вывад. Памочнік `exec.Command` стварае аб'ект
	// для прадстаўлення гэтага знешняга працэсу.
	dateCmd := exec.Command("date")

	// Метад `Output` выконвае каманду, чакае яе завяршэння
	// і збірае стандартны вывад.
	// Калі памылак не было, `dateOut` будзе захоўваць байты
	// з інфармацыяй пра дату.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` і іншыя метады `Command` вернуць
	// `*exec.Error`, калі ўзнікла праблема пры выкананні
	// каманды (напрыклад, няправільны шлях), і `*exec.ExitError`
	// калі каманда выканалася, але завяршылася з ненулявым кода.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// Далей мы разгледзім крыху больш складаны выпадак
	// дзе мы перадаем дадзеныя знешняму працэсу па яго
	// `stdin` і збіраем вынікі з яго `stdout`.
	grepCmd := exec.Command("grep", "hello")

	// Тут мы відавочна захопліваем уваходныя/выхадныя каналы, запускаем
	// працэс, запісваем у яго ўваходныя дадзеныя, чытаем
	// вынік і, нарэшце, чакаем выхаду працэсу.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// У прыведзеным вышэй прыкладзе мы прапусцілі праверкі на памылкі, але
	// можна выкарыстоўваць звычайны шаблон `if err != nil` для
	// ўсіх іх. Мы таксама збіраем толькі вынікі `StdoutPipe`
	// але можна збіраць `StderrPipe` акім жа чынам.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Звярніце ўвагу, што пры стварэнні каманд нам трэба
	// падаць выразна акрэсленую каманду і
	// масіў аргументаў, замест таго, каб проста перадаць адзін
	// радок каманднага радка. Калі вы хочаце стварыць поўную
	// каманду з радком, вы можаце выкарыстоўваць опцыю `-c` у `bash`:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
