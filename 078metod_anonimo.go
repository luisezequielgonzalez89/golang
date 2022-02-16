package main
import "fmt"

type Usuario struct {
	nombre string
	clave string
}
type Administrador struct {
	Usuario
	privilegio int
}
func(user *Usuario) cargaUsuario(nombre, clave string) {
	user.nombre = nombre
	user.clave = clave
}
func(user Usuario) imprimir() {
	fmt.Println("El usuario es: ", user.nombre, "-- Su clave: ", user.clave)
}
func(admin *Administrador) cargar(nombre, clave string, privilegio int) {
	admin.Usuario.cargaUsuario(nombre, clave)
	admin.privilegio = privilegio
}
func (admin Administrador) imprimirPrivilegios() {
	admin.imprimir()
	fmt.Println("Los privilegios asignados son de nivel: ", admin.privilegio)
}
func main(){
	var administrador1 Administrador
	administrador1.cargar("mongoaurelio", "password1234", 1)
	administrador1.imprimirPrivilegios()
}