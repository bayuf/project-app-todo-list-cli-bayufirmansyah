package services

import (
	"errors"
	"project-app-todo-list-cli-bayufirmansyah/model"
	"project-app-todo-list-cli-bayufirmansyah/utils"
	"strings"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) AddTask(task model.Task) error {
	// Validasi input
	if err := utils.CheckInput(task); err != nil {
		return err
	}

	// decoder json lama ambil slice
	tasks, err := utils.DecoderTask()
	if err != nil {
		return err
	}

	//validasi double title
	for i := range tasks {
		if utils.LowerString(tasks[i].Title) == utils.LowerString(task.Title) {
			return errors.New("title already used")
		}
	}

	// index generator
	task.ID = len(tasks) + 1

	tasks = append(tasks, task)

	// tulis tugas baru ke json
	if err := utils.EncoderTask(tasks); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateTask(updateTask model.UpdateTask) error {
	tasks, err := utils.DecoderTask()
	if err != nil {
		return err
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == updateTask.ID {
			// menemukan index yg sama
			found = true

			if updateTask.Status != nil {
				tasks[i].Status = *updateTask.Status
			}
			if updateTask.Priority != nil {
				tasks[i].Priority = *updateTask.Priority
			}

			// berhenti ketika sudah di temukan
			break
		}
	}

	if !found {
		return errors.New("task not found")
	}

	// update ke json
	if err := utils.EncoderTask(tasks); err != nil {
		return err
	}

	return nil
}

func (s *Service) ShowAllTask() error {
	tasks, err := utils.DecoderTask()
	if err != nil {
		return err
	}

	utils.PrintTable(tasks)

	return nil
}

func (s *Service) DeleteTask(id int) error {
	// decode json
	tasks, err := utils.DecoderTask()
	if err != nil {
		return err
	}

	index := -1
	// cari id
	for i, task := range tasks {
		if task.ID == id {
			// update index value
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("id not found")
	}

	// hapus task dari slice tasks
	tasks = append(tasks[:index], tasks[index+1:]...)

	// encode kembali slice ke json
	if err := utils.EncoderTask(tasks); err != nil {
		return err
	}

	return nil
}

func (s *Service) FindTaskByTitle(title string) error {
	tasks, err := utils.DecoderTask()
	if err != nil {
		return err
	}

	// slice kosong untuk menyimpan daftar slice
	taskList := []model.Task{}

	for i := range tasks {
		if !utils.IsLenghtValid(title) {
			break
		}

		if strings.Contains(strings.ToLower(tasks[i].Title), strings.ToLower(title)) {
			taskList = append(taskList, tasks[i])
		}
	}

	if len(taskList) == 0 {
		return errors.New("task not found")
	}

	utils.PrintTable(taskList)

	return nil
}
