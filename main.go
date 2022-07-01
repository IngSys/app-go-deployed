package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

func conexionDB() (conexion *sql.DB) {
	//mysql://b88d0a8e947584:823ea48f@us-cdbr-east-06.cleardb.net/heroku_8c0816ad800ce51?reconnect=true
	Driver := "mysql"
	Usuario := "b88d0a8e947584"
	Contrasenia := "823ea48f"
	Nombre := "heroku_8c0816ad800ce51"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(us-cdbr-east-06.cleardb.net)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var templ = template.Must(template.ParseGlob("template/*"))

func main() {

	http.HandleFunc("/", Start)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/actualizar", Actualizar)

	log.Println("Corriendo servidor...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"

	}
	http.ListenAndServe(":"+port, nil)

}

type Empleado struct {
	Id     int
	Nombre string
	Correo string
}

func Start(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hola mundo")

	conexionEstablecida := conexionDB()
	dataEmpleado, err := conexionEstablecida.Query("SELECT id,nombre,correo FROM empleados")

	if err != nil {
		panic(err.Error())
	}

	empleado := Empleado{}
	listEmpleado := []Empleado{}

	for dataEmpleado.Next() {
		var id int
		var nombre string
		var correo string
		err = dataEmpleado.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}

		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

		listEmpleado = append(listEmpleado, empleado)
	}
	fmt.Println(listEmpleado)

	templ.ExecuteTemplate(w, "index", listEmpleado)

}

func Crear(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "crear", nil)

}

func Insertar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		conexionEstablecida := conexionDB()
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre,correo) VALUES(?,?)")

		if err != nil {
			panic(err.Error())
		}

		insertarRegistros.Exec(nombre, correo)

		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	idEmpleado := r.URL.Query().Get("id")
	//fmt.Println(idEmpleado)
	conexionEstablecida := conexionDB()
	deleteRegistros, err := conexionEstablecida.Prepare("DELETE FROM empleados WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	deleteRegistros.Exec(idEmpleado)

	http.Redirect(w, r, "/", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	//fmt.Println(idEmpleado)
	conexionEstablecida := conexionDB()
	data, err := conexionEstablecida.Query("SELECT * FROM empleados WHERE id=?", idEmpleado)
	if err != nil {
		panic(err.Error())
	}

	empleado := Empleado{}
	for data.Next() {
		var id int
		var nombre string
		var correo string
		err = data.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}

		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

	}
	///	fmt.Println(empleado)
	templ.ExecuteTemplate(w, "editar", empleado)
}

func Actualizar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		conexionEstablecida := conexionDB()
		editarRegistros, err := conexionEstablecida.Prepare("UPDATE empleados SET nombre=?,correo=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		editarRegistros.Exec(nombre, correo, id)

		http.Redirect(w, r, "/", 301)
	}

}
