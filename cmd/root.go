package cmd

import (
	"errors"
	"fmt"
	"os"
	"project-app-todo-list-cli-bayufirmansyah/model"
	"project-app-todo-list-cli-bayufirmansyah/services"
	"slices"

	"github.com/spf13/cobra"
)

var (
	svc   *services.Service
	flags Flags
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Mini Project CLI Todo List",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runTodo()
	},
}

func Execute(service *services.Service) {
	svc = service
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}

func init() {
	// add
	rootCmd.Flags().StringVar(&flags.Add, "add", "", "Tambah Task Baru")
	// list
	rootCmd.Flags().BoolVar(&flags.List, "list", false, "Menampilkan Semua Task")

	// update
	rootCmd.Flags().IntVar(&flags.UpdateId, "update", -1, "Update Task by ID")
	rootCmd.Flags().IntVar(&flags.DoneID, "done", -1, "Update Done Status by ID")
	rootCmd.Flags().StringVar(&flags.Status, "status", "", "Update Status pada Task")
	rootCmd.Flags().StringVar(&flags.Priority, "priority", "", "Update Priority pada Task")

	// del
	rootCmd.Flags().IntVar(&flags.DelId, "delete", -1, "Delete Task by ID")

	// search
	rootCmd.Flags().StringVar(&flags.Search, "search", "", "Search Task by Title or Task Name")

}

func runTodo() error {
	switch {
	case flags.Add != "":
		return handleAdd()
	case flags.List:
		return handleList()
	case flags.UpdateId > 0:
		return handleUpdate()
	case flags.DelId > 0:
		return handleDelete()
	case flags.Search != "":
		return handleSearch()
	case flags.DoneID > 0:
		return handleDoneStatusById()
	}

	return errors.New("flags invalid. use --help instead")
}

func handleDoneStatusById() error {
	status := "done"
	svc.UpdateTask(model.UpdateTask{
		ID:     flags.DoneID,
		Status: &status,
	})
	fmt.Printf("Task ID: %d is %s", flags.DoneID, status)
	return nil
}

func handleAdd() error {
	priority := flags.Priority

	if priority == "" {
		priority = "low"
	}

	// invalid priority value
	valid := []string{"low", "medium", "high"}
	if !slices.Contains(valid, priority) {
		return errors.New("error: invalid priority (low/medium/high)")
	}

	task := model.Task{
		Title:    flags.Add,
		Status:   "new",
		Priority: priority,
	}

	if err := svc.AddTask(task); err != nil {
		return err
	}

	fmt.Printf("Task added: %s | Priority: %s", flags.Add, priority)
	return nil
}

func handleUpdate() error {
	var priorityPtr *string
	var statusPtr *string

	// if user change priority
	if flags.Priority != "" {
		validPriority := []string{"low", "medium", "high"}

		if !slices.Contains(validPriority, flags.Priority) {
			return errors.New("invalid priority (low/medium/high)")
		}

		priorityPtr = &flags.Priority
	}

	if flags.Status != "" {
		validStatus := []string{"done", "progress", "pending"}

		if !slices.Contains(validStatus, flags.Status) {
			return errors.New("invalid status (done/progress/pending)")
		}

		statusPtr = &flags.Status
	}

	if statusPtr == nil && priorityPtr == nil {
		return errors.New("no update value in status or priority")
	}

	if err := svc.UpdateTask(model.UpdateTask{
		ID:       flags.UpdateId,
		Status:   statusPtr,
		Priority: priorityPtr,
	}); err != nil {
		return err
	}

	fmt.Printf("Task Updated ID= %d", flags.UpdateId)
	return nil
}

func handleList() error {
	if err := svc.ShowAllTask(); err != nil {
		return err
	}
	return nil
}

func handleDelete() error {
	if err := svc.DeleteTask(flags.DelId); err != nil {
		return err
	}

	fmt.Printf("Task Deleted ID= %d", flags.DelId)
	return nil
}

func handleSearch() error {
	if err := svc.FindTaskByTitle(flags.Search); err != nil {
		return err
	}
	return nil
}
