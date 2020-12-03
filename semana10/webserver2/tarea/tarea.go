package tarea

type Tarea struct {
	Nombre string
	Estado string
}

type AdminTareas struct {
	Tareas []Tarea
}

func (tareas *AdminTareas) Agregar(tarea Tarea) {
	tareas.Tareas = append(tareas.Tareas, tarea)
}

func (tareas *AdminTareas) String() string {
	var html string
	for _, tarea := range tareas.Tareas {
		html += "<tr>" + "<td>" + tarea.Nombre + "</td>" + "<td>" + tarea.Estado + "</td>" + "</tr>"
	}
	return html
}
