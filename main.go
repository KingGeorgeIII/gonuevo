package main

// This is the package we created in the previous example"
import "github.com/gonuevo/ejercicios"

func main() {
	/*estado, texto := variables.ConviertoaTexto(1258)*/

	numero, texto := ejercicios.ReturnValores2("1258")
	println(numero)
	println(texto)

}
