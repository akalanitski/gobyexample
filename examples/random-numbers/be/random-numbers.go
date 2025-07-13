// Пакет `math/rand/v2` у Go забяспечвае генерацыю
// [псеўдавыпадковых лічбаў](https://en.wikipedia.org/wiki/Pseudorandom_number_generator)

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// Напрыклад, `rand.IntN` вяртае выпадковае `цэлае` n,
	// `0 <= n < 100`.
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// `rand.Float64` вяртае `float64` `f`,
	// `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// Гэта можна выкарыстоўваць для генерацыі выпадковых лікаў з плаваючай кропкай у
	// іншых дыяпазонах, напрыклад, `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Калі патрэбна вядомае пачатковае значэнне, стварыце новае
	// `rand.Source` і перадайце яго ў канструктар `New`
	// `NewPCG` стварае новую
	// [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator)
	// крыніцу, якая патрабуе пачатковага значэння з двух лікаў `uint64`.
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
