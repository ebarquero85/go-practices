// Reference: https://www.youtube.com/watch?v=z75DKfOfDA4
package main

import (
	"fmt"
	"time"
)

// Es como un pipe
// Se bloquea sólo cuando el buffer channel está full
// Puede ser un buffer channel

func main() {

	userch := make(chan string, 2)

	userch <- "Bob"
	userch <- "Alice"
	userch <- "Edgard" // Bloquea porque esta esperando por el consumo del Channel

	user := <-userch

	fmt.Println(user)

}

func example1() { // No Funciona
	userch := make(chan string) // Solo cabe una cookie // Si se coloca un 2 funcionaría
	userch <- "Bob"             // Se bloquea y espera que alguien tome una cookie
	user := <-userch            // Continua bloqueado
	fmt.Println(user)           // fatal error: all goroutines are asleep - deadlock!
}

func example2() { // Funciona porque espera a que alguien escriba

	userch := make(chan string)

	go func() {
		// Las go routines nunca se bloquean
		time.Sleep(2 * time.Second)
		userch <- "Bob" //
	}()

	user := <-userch // Espera hasta que alguien escriba en el channel
	fmt.Println(user)

}

func example3() { // Funciona

	userch := make(chan string, 2) // [" "," "]

	userch <- "Bob" // ["Bob"," "]

	user := <-userch // [" "," "] Espera hasta que alguien escriba en el channel
	fmt.Println(user)

}

func example4() { // Funciona

	userch := make(chan string, 2)

	userch <- "Bob"
	userch <- "Alice"
	// ["Bob", "Alice"]
	user := <-userch
	// ["Alice", ""]

	fmt.Println(user)

}

func example5() { // No Funciona

	userch := make(chan string, 2)

	userch <- "Bob"
	userch <- "Alice"
	userch <- "Edgard" // Bloquea porque esta esperando por el consumo del Channel

	user := <-userch

	fmt.Println(user)

}
