// Go забяспечвае ўбудаваную падтрымку [base64
// кадзіраванне/дэкадзіраванне](https://en.wikipedia.org/wiki/Base64).

package main

// Тут імпартуецца пакет `encoding/base64` з назвай `b64`
// замест стандартнай назвы `base64`. Гэта зэканомць месца ніжэй.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// Вось `string`, які мы будзем кадзіраваць/дэкадзіраваць.
	data := "abc123!?$*&()'-=@~"

	// Go падтрымлівае як стандартныя, так і URL-суладны
	// база64. Вось як кадаваць стандарнай кадзіроўкайю
	// Для кадзіравання патрабуецца `[]byte`, таму мы
	// пераўтварыць наш `string` у гэты тып.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Дэкадзіраванне можа вярнуць памылку, якую можна праверыць
	// калі не ведома ці маюць уваходныя данныя правільныую форму
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// Тут кадзіруецца/дэкадзіруецца ў URL-суладным фармаце
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
