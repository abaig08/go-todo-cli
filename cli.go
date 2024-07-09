package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const todoFile = "todo.json"

type Todo struct {
    ID     int    `json:"id"`
    Task   string `json:"task"`
    Status string `json:"status"`
}

var todos []Todo

func loadTodos() {
    file, err := os.Open(todoFile)
    if err != nil {
        if os.IsNotExist(err) {
            todos = []Todo{}
            return
        }
        fmt.Println("Error loading TODOs:", err)
        return
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&todos); err != nil {
        fmt.Println("Error decoding TODOs:", err)
    }
}

func saveTodos() {
    file, err := os.Create(todoFile)
    if err != nil {
        fmt.Println("Error saving TODOs:", err)
        return
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    if err := encoder.Encode(todos); err != nil {
        fmt.Println("Error encoding TODOs:", err)
    }
}

func addTodo(task string) {
    id := len(todos) + 1
    todo := Todo{ID: id, Task: task, Status: "pending"}
    todos = append(todos, todo)
    saveTodos()
    fmt.Println("Added TODO:", task)
}

func listTodos() {
    for _, todo := range todos {
        fmt.Printf("%d. %s [%s]\n", todo.ID, todo.Task, todo.Status)
    }
}

func deleteTodo(id int) {
    for i, todo := range todos {
        if todo.ID == id {
            todos = append(todos[:i], todos[i+1:]...)
            saveTodos()
            fmt.Println("Deleted TODO:", todo.Task)
            return
        }
    }
    fmt.Println("TODO not found with ID:", id)
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: todo <command> [arguments]")
        fmt.Println("Commands: add, list, delete")
        return
    }

    command := os.Args[1]
    loadTodos()

    switch command {
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Usage: todo add <task>")
            return
        }
        task := strings.Join(os.Args[2:], " ")
        addTodo(task)
    case "list":
        listTodos()
    case "delete":
        if len(os.Args) < 3 {
            fmt.Println("Usage: todo delete <id>")
            return
        }
        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Invalid ID:", os.Args[2])
            return
        }
        deleteTodo(id)
    default:
        fmt.Println("Unknown command:", command)
        fmt.Println("Commands: add, list, delete")
    }
}
