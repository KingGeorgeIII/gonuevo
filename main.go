package main

import (
	"github.com/gonuevo/variables" // This is the package we created in the previous example"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1258)
	println(estado)
	println(texto)

}
