package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"strings"
)

var (
	red    = color.New(color.FgRed).PrintfFunc()
	green  = color.New(color.FgGreen).PrintfFunc()
	yellow = color.New(color.FgYellow).PrintfFunc()
	cyan   = color.New(color.FgCyan).PrintlnFunc()
)

type Task struct {
	ID     int
	Title  string
	IsDone bool
}

var tasks []Task
var nextID int = 1

const (
	StatusPending = "[-]"
	StatusDone    = "[+]"
)

func main() {
	cyan("\n__ Welcome to TODO CLI mini-app __")
	for {
		fmt.Println("1 - Create task.")
		fmt.Println("2 - Show tasks.")
		fmt.Println("3 - Show task by ID.")
		fmt.Println("4 - Update task.")
		fmt.Println("5 - Mark as completed.")
		fmt.Println("6 - Delete task.")
		fmt.Println("7 - Exit.")
		fmt.Println()
		fmt.Print("Choose a command: ")

		var userChoice int
		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			createTask()
		case 2:
			getTasks()
		case 3:
			getTaskByID()
		case 4:
			updateTask()
		case 5:
			completeTask()
		case 6:
			deleteTask()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid command.")
			fmt.Println()
			continue
		}
	}
}

func createTask() {
	fmt.Print("Enter the task title: ")
	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	title = strings.TrimSpace(title)
	if title == "" {
		red("Error: Task title cannot be empty\n")
		return
	}
	tasks = append(tasks, Task{ID: nextID, Title: title, IsDone: false})
	nextID++

	green("Task created successfully!\n")
	fmt.Println()
}

func getTasks() {
	for i := range tasks {
		status := color.YellowString(StatusPending)
		if tasks[i].IsDone {
			status = color.GreenString(StatusDone)
		}
		fmt.Printf("%d. %s %s\n", tasks[i].ID, tasks[i].Title, status)
	}

	if len(tasks) == 0 {
		yellow("No task to show\n")
	}
	fmt.Println()
}

func getTaskByID() {
	fmt.Print("Enter the ID of the task you want to view: ")
	var taskID int
	fmt.Scan(&taskID)
	for i := range tasks {
		if tasks[i].ID == taskID {
			status := color.YellowString(StatusPending)
			if tasks[i].IsDone {
				status = color.GreenString(StatusDone)
			}
			fmt.Printf("%d. %s %s\n", tasks[i].ID, tasks[i].Title, status)
			fmt.Println()
			return
		}
	}

	red("Task not found\n")
	fmt.Println()
}

func updateTask() {
	fmt.Print("Enter the ID of the task you want to update: ")
	var taskID int
	fmt.Scan(&taskID)
	for i := range tasks {
		if tasks[i].ID == taskID {
			fmt.Print("Enter a new title: ")
			reader := bufio.NewReader(os.Stdin)
			title, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			title = strings.TrimSpace(title)
			tasks[i].Title = title
			green("Task updated successfully!\n")
			fmt.Println()
			return
		}
	}

	red("Task not found\n")
	fmt.Println()
}

func completeTask() {
	fmt.Print("Enter the ID of the task you want to mark: ")
	var taskID int
	fmt.Scan(&taskID)
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks[i].IsDone = true
			green("Task marked as completed!\n")
			fmt.Println()
			return
		}
	}

	red("Task not found\n")
	fmt.Println()
}

func deleteTask() {
	fmt.Print("Enter the ID of the task you want to delete: ")
	var taskID int
	fmt.Scan(&taskID)
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			green("Task deleted successfully!\n")
			fmt.Println()
			return
		}
	}

	red("Task not found\n")
	fmt.Println()
}
