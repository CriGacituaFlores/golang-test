package main

import "fmt"

type User struct {
	edad int
	nombre, apellido string
}

func (usuario User) nombre_completo() string {
	return usuario.nombre + " " + usuario.apellido
}


// ek * es para hacer referencia al puntero
func (this *User) set_name(n string) {
	this.nombre = n
}

func main() {
	cristian := new(User)

	cristian.nombre = "cristian"
	cristian.apellido = "gacitua"
	cristian.edad = 24

	cristian.set_name("marcos")

	fmt.Println(cristian.nombre_completo())

}