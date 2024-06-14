package main

import "fmt"

type Tarea struct {
	nombre      string
	descripcion string
	responsable string
	estado      string
}

func newTask(nombre string, descripcion string, responsable string) *Tarea {
	task := Tarea{
		nombre:      nombre,
		descripcion: descripcion,
		responsable: responsable,
		estado:      "pendiente",
	}
	return &task
}

func (e Tarea) changeWorker(responsable string) *Tarea {
	e.responsable = responsable
	return &e
}

func (e Tarea) changeStatus(estado string) *Tarea {
	e.estado = estado
	return &e
}

func displayTask(task Tarea) {
	fmt.Printf("Nombre de la tarea: %s \n", task.nombre)
	fmt.Printf("Descripcion de la tarea: %s \n", task.descripcion)
	fmt.Printf("Responsable de la tarea: %s \n", task.responsable)
	fmt.Printf("Estado de la tarea: %s \n", task.estado)
}

func displayPending(arrTask []Tarea) {
	for _, tarea := range arrTask {
		if tarea.estado == "pendiente" {
			fmt.Println("------------------------------")
			displayTask(tarea)
		}
	}
}

func main() {
	tareas := []Tarea{}
	tareas = append(tareas, *newTask("comprar", "lorem ipsum", "jose"))
	tareas = append(tareas, *newTask("vender", "lorem ipsum", "juan"))
	tareas = append(tareas, *newTask("mostrar", "lorem ipsum", "pepe"))
	displayPending(tareas)
}
