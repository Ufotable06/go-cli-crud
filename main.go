package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/ufotable06/go-cli-crud/tasks"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		messaggeUse()
		return
	}

	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)

	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Escribe el nombre de la tarea")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.AddTask(tasks, name)
		task.SaveTasks(file, tasks)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Debes proporcionar ID de tarea a eliminar")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El id debe ser un numero")
			return
		}

		tasks = task.DeleteTask(tasks, id)
		task.SaveTasks(file, tasks)

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Debes proporcionar ID de tarea a actualizar tarea")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El id debe ser un numero")
			return
		}

		tasks = task.CompleteTask(tasks, id)
		task.SaveTasks(file, tasks)

	default:
		messaggeUse()
	}
}

func messaggeUse() {
	fmt.Println("Uso: go-cli-crud [list|add|complete|delete]")
}
