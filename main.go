package main

import (
	"project-app-todo-list-cli-bayufirmansyah/cmd"
	"project-app-todo-list-cli-bayufirmansyah/services"
)

func main() {
	// init service cobra
	svc := services.NewService()
	cmd.Execute(svc)
}
