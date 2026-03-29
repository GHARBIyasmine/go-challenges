package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

// The backtick tags tell Go: "when converting to/from JSON, use id not ID"
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	dir, _ := os.Getwd()
	path := dir + "\\25-cli\\tasks.json"
	var tasks []Task

	// 1. Read — handle the case where file doesn't exist yet
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			tasks = []Task{} // start fresh
		} else {
			log.Fatal(err)
		}
	} else {
		// file exists, unmarshal it
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			log.Fatal(err)
		}
	}

	add := flag.String("add", "", "add a new task")
	list := flag.Bool("list", false, "list all tasks")
	done := flag.Int("done", 0, "mark task as done")
	delete := flag.Int("delete", 0, "delete a task")

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	if *add != "" {
		tasks = append(tasks, Task{ID: len(tasks) + 1, Title: *add, Done: false})
		fmt.Printf("Added task: %s\n", *add)
	}

	if *list {
		for i, task := range tasks {
			if task.Done {
				fmt.Printf("%d. [x] %s\n", i+1, task.Title)
			} else {
				fmt.Printf("%d. [ ] %s\n", i+1, task.Title)
			}
		}
	}

	if *done > 0 && *done <= len(tasks) {
		tasks[*done-1].Done = true
		fmt.Printf("Marked task %d as done\n", *done)
	}

	if *delete > 0 && *delete <= len(tasks) {
		tasks = append(tasks[:*delete-1], tasks[*delete:]...)
		fmt.Printf("Deleted task %d\n", *delete)
	}

	// 3. Write back (creates file if not exist, overwrites if it does)
	data, err = json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(path, data, 0666)
	if err != nil {
		log.Fatal(err)
	}

}
