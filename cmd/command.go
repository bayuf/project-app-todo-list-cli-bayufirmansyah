package cmd

import (
	"fmt"
	"project-app-todo-list-cli-bayufirmansyah/model"
	"project-app-todo-list-cli-bayufirmansyah/services"
)

type Command struct {
	Services *services.Service
}

func NewCommand(svc *services.Service) *Command {
	return &Command{
		Services: svc,
	}
}

func (c *Command) AddTask() {
	var title string
	var priority string

	c.Services.AddTask(model.Task{
		Title:    title,
		Status:   "New",
		Priority: priority,
	})
}

func (c *Command) ShowAllTask() {
	err := c.Services.ShowAllTask()
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Command) UpdateTask() {
	status := "Done"
	priority := "Low"

	c.Services.UpdateTask(model.UpdateTask{
		ID:       1,
		Status:   &status,
		Priority: &priority,
	})
}

func (c *Command) DeleteTask() {

}

func (c *Command) FindTaskByTitle() {
	if err := c.Services.FindTaskByTitle(""); err != nil {
		fmt.Println(err)
	}
}
