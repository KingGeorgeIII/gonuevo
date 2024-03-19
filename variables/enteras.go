// Package variables provides ...
package variables

import (
	"fmt"
)

func MuestroEnteros() {
	var intcomun int
	intde32 := int32(10)
	intcomun = 10
	intde32 = 20
	fmt.Println(intcomun)
	fmt.Println(intde32)
	fmt.Println(intcomun + int(intde32))

}
