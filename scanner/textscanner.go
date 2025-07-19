package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TextScaner() string {
	scaner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the command: ")

	if ok := scaner.Scan(); !ok {
		return "Error"
	}

	text := scaner.Text()
	return strings.TrimSpace(text)
}
