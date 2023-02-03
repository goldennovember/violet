package main

import (
	"flag"
	"fmt"
	"github.com/goldennovember/control/pcag/todo"
	"github.com/goldennovember/control/pcag/wc"
	"os"
)

const todoFileName = ".todo.json"

func main() {

	// Defining a boolean flag to count words/lines/bytes
	count := flag.Bool("count", false, "Count words/lines/bytes")
	words := flag.Bool("w", false, "Count words")
	bytes := flag.Bool("b", false, "Count bytes")

	// Parsing command line flags for todo

	task := flag.Bool("task", false, "Todo list")
	add := flag.Bool("add", false, "Add task to the ToDo list")
	delete := flag.Int("delete", 0, "Item to be deleted")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Control CLI tool by Hieu Le")
		flag.PrintDefaults()
	}

	// Parsing the flags provided by the user
	flag.Parse()

	switch {

	// If the user provided the -w or -b flag
	// call the wc package to count words/bytes
	case *count:
		fmt.Println(wc.Count(os.Stdin, *words, *bytes))
	case *task:
		// Define an items list
		l := &todo.List{}

		if err := l.Get(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		switch {
		case *list:
			// List current to do items
			fmt.Print(l)

		case *complete > 0:
			// Complete the given item
			if err := l.Complete(*complete); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			// Save the new list
			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		case *add:
			// When any arguments (excluding flags) are provided, they will be
			// used as the new task
			t, err := todo.GetTask(os.Stdin, flag.Args()...)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			l.Add(t)
			// Save the new list
			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		case *delete > 0:

			if err := l.Delete(*delete); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			// Save the new list
			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

		default:
			// Invalid flag provided for todo task
			fmt.Fprintln(os.Stderr, "Invalid option for todo")
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}
