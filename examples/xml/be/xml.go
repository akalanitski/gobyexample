// Go прапануе ўбудаваную падтрымку XML і XML-падобных фарматаў з дапамогай
// пакета `encoding/xml`.

package main

import (
	"encoding/xml"
	"fmt"
)

// Тып `Plant` будзе адлюстраваны ў XML. Падобна да прыкладаў
// JSON, тэгі палёў утрымліваюць дырэктывы для
// кадавальніка і дэкодэра. Тут мы выкарыстоўваем некаторыя спецыяльныя функцыі
// пакета XML: назва поля `XMLName` вызначае
// назву элемента XML, які прадстаўляе гэту структуру;
// `id,attr` азначае, што поле `Id` з'яўляецца XML
// _attribute_, а не ўкладзеным элементам.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Вывесці XML-файл, які прадстаўляе нашу расліну; выкарыстоўваючы
	// `MarshalIndent` для стварэння больш
	// чытэльнага вываду.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Каб дадаць агульны XML-загаловак да вываду, дадайце яго яўна.
	fmt.Println(xml.Header + string(out))

	// Выкарыстоўваць `Unmarshal` для разбору патоку байтаў з XML
	// у структуру дадзеных. Калі XML няправільны або
	// не можа быць разабраны у Plant, будзе вернута памылка.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// Тэг поля `parent>child>plant` паведамляе энкодэру
	// укласці ўсе `plant` у `<parent><child>...`
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
