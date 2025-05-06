// _Пералічвальныя тыпы_ (enum) — гэта асобны выпадак
// [sum types](https://en.wikipedia.org/wiki/Algebraic_data_type).
// Enum — гэта тып, які мае фіксаваную колькасць магчымых
// значэнняў, кожнае з якіх мае асобную назву. Go не мае
// тыпу пералічэння як асобнай моўнай функцыі, але пералічэнні
// лёгка рэалізаваць з выкарыстаннем існуючых моўных сродкаў.

package main

import "fmt"

// Наш тып пералічэння `ServerState` мае базавы тып `int`.
type ServerState int

// Магчымыя значэнні для `ServerState` вызначаюцца як
// канстанты. Спецыяльнае ключавое слова [iota](https://go.dev/ref/spec#Iota)
// аўтаматычна генеруе паслядоўныя канстантныя значэнні; у гэтым
// выпадку 0, 1, 2 і г.д.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Рэалізуючы інтэрфейс [fmt.Stringer](https://pkg.go.dev/fmt#Stringer)
// значэнні `ServerState` можна вывесці або пераўтварыць
// у радкі.
//
// Гэта можа быць складана, калі магчымых значэнняў шмат. У такіх
// выпадках [інструмент stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// можна выкарыстоўваць разам з `go:generate` для аўтаматызацыі
// працэсу. Глядзіце [гэты пост](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)
// для больш падрабязнага тлумачэння.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// Калі ў нас ёсць значэнне тыпу `int`, мы не можам перадаць яго ў `transition`
	// — кампілятар паскардзіцца на неадпаведнасць тыпаў. Гэта забяспечвае
	// пэўную ступень бяспекі тыпаў падчас кампіляцыі для пералічэнняў.

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition эмулюе пераход стану для
// сервера; ён прымае існуючы стан і вяртае
// новы стан.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Уявім, мы правяраем тут некаторыя прэдыкаты, каб
		// вызначыць наступны стан...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
