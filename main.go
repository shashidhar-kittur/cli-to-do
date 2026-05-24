package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/AnshKumar200/cli-to-do/models"
	"github.com/AnshKumar200/cli-to-do/todo"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var exit bool = false
	var choice int

	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	var TodoData []models.TodoItem
	err = json.Unmarshal(data, &TodoData)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	for !exit {
		fmt.Println("What do you want to do?")
		fmt.Println("1 - View Tasks")
		fmt.Println("2 - Add")
		fmt.Println("3 - Mark As Complete")
		fmt.Println("4 - Delete")
		fmt.Println("5 - Exit")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			todo.ViewAllTask(TodoData)
		case 2:
			fmt.Print("Task: ")
			TaskName, _ := reader.ReadString('\n')
			TaskName = strings.TrimSpace(TaskName)
			todo.AddTask(TaskName, &TodoData)
		case 3:
			var index int
			fmt.Print("Enter the index of the task: ")
			fmt.Scan(&index)
			todo.MarkTask(index, &TodoData)
		case 4:
			var index int
			fmt.Print("Enter the index of the task: ")
			fmt.Scan(&index)
			todo.DeleteTask(index, &TodoData)
		case 5:
			os.Exit(1)
		}

		fmt.Println("=====================")
	}

}
