package main

import (
	"fmt"
	"sort"
)

type Tasks struct {
	ID          int
	Description string
	Priority    int
}

var taskList []*Tasks

func main() {
	fmt.Println("To-Do List Manager")
	fmt.Println("1. Add Task")
	fmt.Println("2. View Task")
	fmt.Println("3. Delete Task")
	fmt.Println("4. Update Task")
	fmt.Println("5. Exit")
	fmt.Println("Choose an option: ")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		// continue
	}

	switch choice {
	case 1:
		addTask()
	case 2:
		viewTasks()
	case 3:
		deleteTask()
	case 4:
		updateTask()
	case 5:
		fmt.Println("Exiting... Goodbye!")
		return
	default:
		fmt.Println("Invalid choice, Please try again.")
	}
}

func addTask() {
	var description string
	var priority, id int

	fmt.Print("Enter task description: ")
	fmt.Scanln(&description)

	if description == "" {
		fmt.Println("Error: Description cannot be empty.")
		return
	}

	fmt.Print("Enter task prioty (1 = High, 2 = Medium, 3 = Low): ")
	fmt.Scanln(&priority)

	if priority < 1 || priority > 3 {
		fmt.Println("Error: priority must be between 1 and 3.")
		return
	}

	id = len(taskList) + 1
	taskList = append(taskList, &Tasks{ID: id, Description: description, Priority: priority})
	fmt.Printf("Task added successfully! ID: %d, Decription: %s, Priority: %d\n", id, description, priority)
}

func viewTasks() {
	if len(taskList) == 0 {
		fmt.Println("No tasks to display.")
		return
	}

	// sort tasks by priority (High to Low)
	sort.Slice(taskList, func(i, j int) bool {
		return taskList[i].Priority < taskList[j].Priority
	})

	fmt.Println("\nTasks (sorted by priority): ")
	for _, task := range taskList {
		priorityStr := priorityToString(task.Priority)
		fmt.Printf("ID: %d, Description: %s, Priority: %s\n", task.ID, task.Description, priorityStr)
	}
}

func updateTask() {
	var id, priority int
	var description string
	fmt.Print("Enter the ID of the task to update: ")
	fmt.Scan(&id)

	for i, t := range taskList {
		if t.ID == id {
			fmt.Print("Enter new decription (leave empty to keep current): ")
			fmt.Scanln(&description)
			if description != "" {
				taskList[i].Description = description
			}

			taskList[i].Description = description
			taskList[i].Priority = priority

			fmt.Print("Enter new priority (1 = High, 2 = Medium, 3 = Low): ")
			_, err := fmt.Scan(&priority)
			if err != nil && priority >= 1 && priority <= 3 {
				taskList[i].Priority = priority
			} else {
				fmt.Println("Invalid priority. Keeping current priority.")
			}
			fmt.Printf("Task with ID %d updated successfully.\n", id)
		}
		fmt.Printf("Error: Task with ID %d not found.\n", id)
	}

}

func deleteTask() {
	var id int
	fmt.Println("Enter the ID of the task to delete: ")
	fmt.Scan(&id)

	for i, t := range taskList {
		if t.ID == id {
			taskList = append(taskList[:i], taskList[i+1:]...)
			fmt.Printf("Task with ID %d deleted successfully.\n", id)
		}
	}

	// deleting id which does not existin the list.
	fmt.Printf("Error: Task with ID %d not found\n", id)
}

func priorityToString(priority int) string {
	switch priority {
	case 1:
		return "High"
	case 2:
		return "Medium"
	case 3:
		return "Low"
	default:
		return "Unkwon"
	}
}
