package utils

import (
	"fmt"
	"os"
	"project-app-todo-list-cli-bayufirmansyah/model"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const (
	Blue   = "\033[34m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Reset  = "\033[0m"
)

func ColorChangeByStatus(status string) string {
	switch strings.ToLower(status) {
	case "done":
		return Green + status + Reset
	case "pending":
		return Yellow + status + Reset
	case "progress":
		return Cyan + status + Reset
	case "new":
		return Blue + status + Reset
	default:
		return status
	}
}

func PrintTable(tasks []model.Task) {
	fmt.Println("======================  List Tasks  ======================")
	table := tablewriter.NewWriter(os.Stdout)

	table.Header([]string{"No", "ID", "Task", "Status", "Priority"})

	for i, t := range tasks {
		row := []string{
			fmt.Sprintf("%d", i+1),
			fmt.Sprintf("%d", t.ID),
			t.Title,
			ColorChangeByStatus(t.Status),
			t.Priority,
		}

		table.Append(row)
	}
	table.Render()
}
