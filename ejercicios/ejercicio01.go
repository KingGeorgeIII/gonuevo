package ejercicios

import (
	"strconv"
)

func ReturnValores2(texto string) (int, string) {

	numero, err := strconv.Atoi(texto)
	if err != nil {
		panic(err)
	}
	var texto1 string
	if numero > 100 {
		texto1 = "Mayor a 100"
	} else {
		texto1 = "Menor a 100"
	}

	return numero, texto1

}
