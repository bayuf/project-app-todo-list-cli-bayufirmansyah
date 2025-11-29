package utils

import (
	"encoding/json"
	"errors"
	"os"

	"project-app-todo-list-cli-bayufirmansyah/model"
)

func fileCheckExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func getJsonFileName() (string, error) {
	fileName := "data/tasks.json"
	if !fileCheckExist(fileName) {
		// buat folder data jika belum ada
		os.MkdirAll("data", 0755)

		// buat file
		err := os.WriteFile(fileName, []byte("[]"), 0644)
		if err != nil {
			return "", errors.New("gagal membuat file json")
		}
	}

	return fileName, nil
}

// Decode File Json menjadi slice
func DecoderTask() ([]model.Task, error) {

	fileJson, err := getJsonFileName()
	if err != nil {
		return nil, err
	}

	jsonData, err := os.ReadFile(fileJson)
	if err != nil {
		return nil, errors.New("gagal membaca file json")
	}

	var tasks []model.Task

	if err := json.Unmarshal(jsonData, &tasks); err != nil {
		return nil, errors.New("decode json gagal")
	}

	return tasks, nil

}

// Streaming Json File dan menulis
func EncoderTask(tasks []model.Task) error {
	fileJson, err := getJsonFileName()
	if err != nil {
		return err
	}
	file, err := os.Create(fileJson)
	if err != nil {
		return errors.New("cant create file json")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(tasks)

	return nil

}
