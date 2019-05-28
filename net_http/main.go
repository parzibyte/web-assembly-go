// +build js,wasm
/*
	Petición HTTP GET con WebAssembly
	¿XMLHttpRequest o fetch? no, mejor WebAssembly con Go

	@author parzibyte
	Visita: parzibyte.me/blog
*/
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"syscall/js"
)

// Una simple función que hace una petición GET HTTP
// para demostrar cómo podemos usar net/http de GO
// en el navegador web. Sí, has leído bien
// Mira: https://parzibyte.me/blog/2019/05/21/peticion-post-get-put-delete-go-net-http/
func peticionHttp() (string, error) {
	clienteHttp := &http.Client{}
	// Si quieres agregar parámetros a la URL simplemente haz una
	// concatenación :)
	url := "https://httpbin.org/get"
	peticion, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	// Podemos agregar encabezados
	peticion.Header.Add("Content-Type", "application/json")
	peticion.Header.Add("X-Hola-Mundo", "Ejemplo")
	respuesta, err := clienteHttp.Do(peticion)
	if err != nil {
		// Maneja el error de acuerdo a tu situación
		return "", err
	}
	// No olvides cerrar el cuerpo al terminar
	defer respuesta.Body.Close()

	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		return "", err
	}

	log.Printf("Código de respuesta: %d", respuesta.StatusCode)
	log.Printf("Encabezados: '%q'", respuesta.Header)
	contentType := respuesta.Header.Get("Content-Type")
	log.Printf("El tipo de contenido: '%s'", contentType)
	// Aquí puedes decodificar la respuesta si es un JSON, o convertirla a cadena

	respuestaString := string(cuerpoRespuesta)
	log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)
	return respuestaString, nil
}

func main() {
	// Declaración de elementos del DOM

	// Equivalente a document.querySelector("#hacerPeticion")
	botonHacerPeticion := js.Global().Get("document").Call("querySelector", "#hacerPeticion")

	// Equivalente a document.querySelector("#resultado")
	resultado := js.Global().Get("document").Call("querySelector", "#resultado")

	// Un canal para eso del código asíncrono y las rutinas que envuelve las llamadas HTTP
	canal := make(chan struct{})
	// botonHacerPeticion.addEventListener("click", () => {})
	botonHacerPeticion.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resultado.Set("innerHTML", "Cargando...")
		log.Println("Haciendo petición")
		go func() {

			respuestaString, err := peticionHttp()
			log.Printf("El error: '%v'", err)
			if err != nil {
				resultado.Set("innerHTML", "Error haciendo petición")
			}
			// resultado.innerHTML = respuestaString;
			resultado.Set("innerHTML", respuestaString)
		}()
		return nil
	}))

	<-canal
}
