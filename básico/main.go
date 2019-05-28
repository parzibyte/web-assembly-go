// +build js,wasm
/*
	Un acercamiento a WebAssembly con Go

	@author parzibyte
	Visita: parzibyte.me/blog
*/
package main

import (
	"log"
)

func main() {
	log.Printf("Hola WebAssembly. 5 + 5 = %d", 5+5)
}
