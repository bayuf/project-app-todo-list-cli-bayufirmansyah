package utils

import (
	"errors"
	"project-app-todo-list-cli-bayufirmansyah/model"
	"strings"
)

func CheckInput(m model.Task) error {

	if strings.TrimSpace(m.Title) == "" {
		return errors.New("title is empty")
	}
	if strings.TrimSpace(m.Status) == "" {
		return errors.New("status is empty")
	}
	if strings.TrimSpace(m.Priority) == "" {
		return errors.New("priority is empty")
	}

	return nil
}

func IsLenghtValid(word string) bool {
	return len(word) >= 5
}
