package main

import (
	"fmt"
	"net/http"
)

var taskItems = []string{}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.HandleFunc("/add-task", addTask)

	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	// Set the Content-Type to HTML
	writer.Header().Set("Content-Type", "text/html")

	// Create the HTML response
	var greeting = "Hello user. Welcome to our Todolist App!<br>"
	greeting += "<form action='/add-task' method='POST'>" +
		"<input type='text' name='task' placeholder='Enter a new task'/>" +
		"<input type='submit' value='Add Task'/>" +
		"</form>"

	// Write the HTML response
	fmt.Fprintln(writer, greeting)
}


func showTasks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	for _, task := range taskItems {
		fmt.Fprintf(writer, "%s <br>", task)
	}
	fmt.Fprintln(writer, "<a href='/'>Add Another Task</a>")
}

// Handle form submission to add a task
func addTask(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		writer.Header().Set("Content-Type", "text/html")
		// Parse the form data
		request.ParseForm()
		newTask := request.FormValue("task")

		// Add the new task to the list if it's not empty
		if newTask != "" {
			taskItems = append(taskItems, newTask)
			fmt.Fprintf(writer, "Task '%s' added successfully!<br>", newTask)
		} else {
			fmt.Fprintln(writer, "Task cannot be empty.<br>")
		}

		// Show the updated list of tasks
		fmt.Fprintln(writer, "<a href='/show-tasks'>View all tasks</a>")
	} else {
		fmt.Fprintln(writer, "Only POST method is supported for adding tasks.")
	}
}