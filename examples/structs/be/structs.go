// _Структуры_ у Go — гэта тыпізаваныя калекцыі палёў.
// Яны карысныя для групавання дадзеных разам

package main

import "fmt"

// Тып `person`, гэта структура якая мае палі `name` і `age`.
type person struct {
	name string
	age  int
}

// `newPerson` стварае экзэмпляр структуры `person` і зададзе імя.
func newPerson(name string) *person {
	// Go — гэта мова праграмавання са зборшчыкам смецця; вы можаце бяспечна
	// вярнуць указальнік на лакальную зменную — яна будзе ачышчана толькі
	// зборшчыкам смецця, калі на яе не будзе актыўных спасылак.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// Гэты сінтаксіс стварае экзэмпляр структуру і задае значэнні яе палям
	fmt.Println(person{"Bob", 20})

	// Вы можаце выкрысоўваць імёны палей пры заданні іх значэнняў.
	fmt.Println(person{name: "Alice", age: 30})

	// Прапушчаныя палі будуць мець нулявыя значэнні.
	fmt.Println(person{name: "Fred"})

	// Прэфікс `&` дае ўказальнік на структуру.
	fmt.Println(&person{name: "Ann", age: 40})

	// Ідыяматычна інкапсуляваць стварэнне новай структуры ў функцыі-канструктары
	fmt.Println(newPerson("Jon"))

	// Паля структур даступны праз кропу.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Можна выкарыстоўваць кропку і з указальнікамі на структуры.
	// Пры звароце да поля структуры, указальнік аўтаматычна разіменоўваецца
	sp := &s
	fmt.Println(sp.age)

	// Структуры зменлівыя (mutable).
	sp.age = 51
	fmt.Println(sp.age)

	// Калі тып структуры выкарыстоўваецца толькі для аднаго значэння,
	// нам не абавязкова даваць яму назву. Значэнне можа мець ананімны
	// тып структуры. Гэты метад звычайна выкарыстоўваецца для
	// [таблічна-кіраваных тэстаў](testing-and-benchmarking).
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
