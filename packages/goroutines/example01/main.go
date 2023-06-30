package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	listaSaludar := []string{"Edgard", "Uriel", "Barquero", "Real"}
	listaDespedir := []string{"Luis", "Calors", "María", "Luisa"}

	go saludar(listaSaludar)
	go despedir(listaDespedir)

	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Texto digitado:", s)

}

func saludar(nombres []string) {

	for _, nombre := range nombres {
		fmt.Printf("Saludos %s\n", nombre)
		time.Sleep(time.Millisecond * 700)
	}

}

func despedir(nombres []string) {

	for _, nombre := range nombres {
		fmt.Printf("Adiós %s\n", nombre)
		time.Sleep(time.Millisecond * 400)
	}

}
