package todo

import (
	"ToDo/scanner"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

func HandleCommand(comMap map[string]string, com [5]string) {
	fmt.Println("Its the commands you can use:")
	for i := range com {
		fmt.Printf("--%s\n", com[i])
	}

	input := scanner.TextScaner()
	cmd, arg := parseCommand(input)

	switch cmd {
	case "help":
		ShowHelp(comMap)
	case "add":
		if arg == "" {
			fmt.Println("Enter the title")
			return
		}
		AddNewTask(arg)
		fmt.Printf("Task %s is added\n", arg)
		time.Sleep(3 * time.Second)
	case "list":
		List()
	case "complete":
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Enter the valid id")
		}

		if CompleteTask(id) {
			fmt.Printf("The task with %d id is Complete!\n", id)
		} else {
			fmt.Printf("There is no task with this id\n")
		}
		time.Sleep(3 * time.Second)
	case "del":
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Please provide valid task ID")
			return
		}
		if DelTask(id) {
			fmt.Printf("Task %d deleted!\n", id)
		} else {
			fmt.Printf("Task %d not found\n", id)
		}
		time.Sleep(3 * time.Second)
	case "exite":
		os.Exit(0)
	default:
		fmt.Println("Unknow command, use 'help' to know which commands you can use")
	}
}

func ShowHelp(comMap map[string]string) {
	w := tabwriter.NewWriter(os.Stdout, 14, 2, 4, ' ', 0)
	w.Write([]byte("Command\tDescritpion\n"))
	for k, v := range comMap {
		w.Write([]byte("--" + k + "\t" + v + "\n"))
		w.Flush()
	}
	fmt.Println()
	fmt.Println("input format [command] 'Text'")
}

func parseCommand(input string) (cmd, arg string) {
	if len(input) == 0 {
		return "", ""
	}

	for i, c := range input {
		if c == ' ' {
			return input[:i], input[i+1:]
		}
	}
	return input, ""
}
