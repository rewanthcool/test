package B

import (
	"fmt"

	a "github.com/rewanthcool/test/A"
)

func B() {
	fmt.Println("B")
	a.A()
}
