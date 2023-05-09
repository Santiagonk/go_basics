package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre", tarea.nombre)
		fmt.Println("Descripción", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados() {
	for _, tarea := range t.tasks {
		if tarea.completado {
			fmt.Println("Nombre", tarea.nombre)
			fmt.Println("Descripción", tarea.descripcion)
		}
	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	// t := &task{
	// 	nombre:      "Completar mi curso de Go",
	// 	descripcion: "Completar mi curso de Go en Platzi en esta semana",
	// }
	// fmt.Printf("%+v\n", t)
	// t.marcarCompleta()
	// t.actualizarNombre("Finalizar mi curso de Go")
	// t.actualizarDescripcion("Finalizar mi curso cuanto antes")
	// fmt.Printf("%+v\n", t)
	t1 := &task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python en Platzi en esta semana",
	}
	t2 := &task{
		nombre:      "Completar mi curso de Js",
		descripcion: "Completar mi curso de Js en Platzi en esta semana",
	}
	t3 := &task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go en Platzi en esta semana",
	}
	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	fmt.Println(lista.tasks[0])
	lista.agregarALista(t3)
	// fmt.Println(len(lista.tasks))
	// lista.eliminarDeLista(1)
	// fmt.Println(len(lista.tasks))

	// for i := 0; i < len(lista.tasks); i++ {
	// 	fmt.Println("Index", i, "nombre", lista.tasks[i].nombre)
	// }

	// for index, tarea := range lista.tasks {
	// 	fmt.Println("Index", index, "nombre", tarea.nombre)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println("Index", i)
	// }
	// lista.imprimirLista()
	lista.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Santiago"] = lista

	t4 := &task{
		nombre:      "Completar mi curso de Java",
		descripcion: "Completar mi curso de Java en Platzi en esta semana",
	}
	t5 := &task{
		nombre:      "Completar mi curso de C#",
		descripcion: "Completar mi curso de C# en Platzi en esta semana",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Africa"] = lista2
	fmt.Println("Tareas Santiago")
	mapaTareas["Santiago"].imprimirLista()
	fmt.Println("Tareas Africa")
	mapaTareas["Africa"].imprimirLista()
}
