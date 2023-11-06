package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/StillLearnSVN/go-todo"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	del := flag.Int("del", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		handleError(err)
	}

	switch {
	case *add:
		task := getInput(flag.Args()...)
		todos.Add(task)
		handleError(todos.Store(todoFile))
	case *complete > 0:
		handleError(todos.Complete(*complete))
		handleError(todos.Store(todoFile))
	case *del > 0:
		handleError(todos.Delete(*del))
		handleError(todos.Store(todoFile))
	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func getInput(args ...string) string {
	input := strings.Join(args, " ")
	if input == "" {
		fmt.Print("Enter todo: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		} else {
			handleError(scanner.Err())
		}
	}
	return input
}