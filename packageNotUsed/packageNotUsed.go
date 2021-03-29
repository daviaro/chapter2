package main

import (
	"fmt"
	_ "os"
)

//El paquete os no se usa por lo que el compilado presentara errores.
// para eso se le coloca el _ para evitar los errores en la compilacion.
func main() {
	fmt.Println("Hello there!")
}
