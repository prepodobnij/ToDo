package todo

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Task struct {
	Id        int
	Title     string
	CreatedAt time.Time
	Done      bool
}

var tasks []Task
var lastId int = 0

func AddNewTask(title string) {
	lastId++
	task := Task{
		Id:        lastId,
		Title:     title,
		CreatedAt: time.Now(),
		Done:      false,
	}

	tasks = append(tasks, task)
}

func List() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTitle\tStatus\tCreated At")
	for _, v := range tasks {
		status := "panding"
		if v.Done {
			status = "Done"
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n",
			v.Id,
			v.Title,
			status,
			v.CreatedAt.Format("2006-01-02 15:04"))
	}
	w.Flush()

	fmt.Print("Enter '..' to return to the main menu: ")
	var ex string
	fmt.Scan(&ex)
	if ex == ".." {
		return
	}

}

func CompleteTask(id int) bool {
	for i, v := range tasks {
		if id == v.Id {
			tasks[i].Done = true
			return true
		}
	}
	return false
}

func DelTask(id int) bool {
	for i, t := range tasks {
		if t.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
