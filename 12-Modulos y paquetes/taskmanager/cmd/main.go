package main

import (
	"fmt"

	"github.com/ricardocuellar/taskmanager/internal/tasks"
)

func main() {
	store := tasks.NewStore()
	store.Add("Aprender Go Modules")
	store.Add("Hacer proyecto del curso de Python")
	store.Add("FastAPI desde cero")
	store.Add("DUPLICADO")

	fmt.Println("===MIS TAREAS===")
	for _, task := range store.List() {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}

		fmt.Printf("%s %d. %s\n", status, task.ID, task.Title)
	}

	fmt.Println("Completar tarea:")
	store.Complete(1)
	store.Complete(0)

	store.Delete(3)

	for _, task := range store.List() {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}

		fmt.Printf("%s %d. %s\n", status, task.ID, task.Title)
	}

}
