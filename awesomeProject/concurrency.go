package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	go mi_nombre_lentooo("Cristian eduardo gacitua flores")
	fmt.Println("Que aburridooo")
	var wait string
	fmt.Scanln(&wait)
}

func mi_nombre_lentooo(name string) {
	letras := strings.Split(name,"")

	for _, letra := range(letras) {
		time.Sleep(400 * time.Millisecond)
		fmt.Println(letra)
	}
}