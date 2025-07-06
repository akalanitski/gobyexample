// Go прапануе ўбудаваную падтрымку кадавання і дэкадавання JSON,
// у тым ліку пераход да ўбудаваных і сваіх тыпаў дадзеных.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Мы выкарыстаем гэтыя дзве структуры для дэманстрацыі кадавання і
// дэкадавання тыпаў ніжэй.
type response1 struct {
	Page   int
	Fruits []string
}

// Толькі экспартаваныя палі будуць закадзіраваны/дэкадаваны ў JSON.
// Палі павінны пачынацца з вялікіх літар для экспарту.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// Спачатку мы разгледзім кадаванне асноўных тыпаў дадзеных у
	// JSON-тэкст. Вось некалькі прыкладаў для прымітыўных
	// значэнняў.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// А вось некаторыя для слайсаў і мап, якія кадуюць
	// у масівы і аб'екты JSON, як і варта было чакаць.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Пакет JSON можа аўтаматычна кадзіраваць вашы
	// тыпы дадзеных. Ён будзе ўключаць толькі экспартаваныя
	// палі ў закадаваны вынік і па змаўчанні будзе
	// выкарыстоўваць гэтыя назвы ў якасці ключоў JSON.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Вы можаце выкарыстоўваць тэгі ў аб'явах палёў структуры
	// для налады закадаваных імёнаў ключоў JSON. Праверце
	// вызначэнне `response2` вышэй, каб убачыць прыклад
	// такіх тэгаў.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Цяпер давайце разгледзім дэкадаванне JSON-а ў тыпы вашай
	// праграамы Go. Ніжэй прыклад агульнай структуры дадзеных
	//.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Нам трэба падаць зменную, куды JSON
	// пакет можа змясціць дэкадаваныя даныя. Гэта
	// `map[string]interface{}` будзе ўтрымліваць мапу радкоў
	// для адвольных тыпаў даных.
	var dat map[string]interface{}

	// Вось уласна дэкадаванне і праверка на наяўнасць звязаных з ім памылак.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Каб выкарыстоўваць значэнні з дэкадаванай карты,
	// нам трэба будзе пераўтварыць іх у адпаведны тып.
	// Напрыклад, тут мы пераўтвараем значэнне `num` у
	// тып `float64`.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Доступ да ўкладзеных дадзеных патрабуе серыі
	// пераўтварэнняў.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Мы таксама можам дэкадаваць JSON у карыстальніцкія тыпы дадзеных.
	// Гэта мае перавагі, бо дадае дадатковую
	// бяспеку у нашы праграмы і ліквідуе
	// неабходнасць у сцвярджэннях тыпаў пры доступе да дэкадаваных
	// дадзеных.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// У прыведзеных вышэй прыкладах мы заўсёды выкарыстоўвалі байты і
	// тэкст у якасці прамежкавых злучэнняў паміж дадзенымі і
	// прадстаўленнем JSON на стандартным вывадзе. Мы таксама можам
	// перадаваць кадоўкі JSON непасрэдна ў `os.Writer`, напрыклад,
	// `os.Stdout` або нават у цела адказаў HTTP.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// Струменевае чытанне з `os.Reader`, напрыклад, `os.Stdin`
	// або целаў HTTP-запытаў выконваецца з дапамогай `json.Decoder`.
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}
