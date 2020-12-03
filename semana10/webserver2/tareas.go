package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"./tarea"
)

var misTareas tarea.AdminTareas

func cargarHTML(a string) string {
	html, _ := ioutil.ReadFile(a)

	return string(html)
}

func tareas(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHTML("./webserver2/tabla.html"),
			misTareas.String(),
		)
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		fmt.Println(req.PostForm)
		tarea := tarea.Tarea{Nombre: req.FormValue("tarea"), Estado: req.FormValue("estado")}
		misTareas.Agregar(tarea)
		fmt.Println(misTareas)
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHTML("./webserver2/respuesta.html"),
			tarea.Nombre,
		)
	}
}

func form(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHTML("./webserver2/form.html"),
	)
}

func main() {
	http.HandleFunc("/form", form)
	http.HandleFunc("/tareas", tareas)
	fmt.Println("Corriendo servidor de tareas...")
	http.ListenAndServe(":9000", nil)
}
