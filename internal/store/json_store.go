package store

import (
    // "fmt"
	"os"
    "path/filepath"
	"encoding/json"
	"errors"

	"task/models"
)

const (
    pathCli = "/.config/.mycli"
    fileStore = "/todo.json"
)

func checkFile(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		return false
	}
	return true
}

// func displayTask()

func Save(task models.Task) error {
    home, err := os.UserHomeDir()
    if err != nil {
        return err
    }

    configPath := filepath.Join(home, pathCli)
    if err := os.MkdirAll(configPath, 0750); err != nil {
        return err
    }

    path := filepath.Join(configPath, fileStore)

    var tasks []models.Task

    if checkFile(path) {
        read, err := os.ReadFile(path)
        if err != nil {
            return err
        }

        _ = json.Unmarshal(read, &tasks)
    }
    tasks = append(tasks, task)

    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(path, data, 0644)
}

func AllList() ([]models.Task, error)  {
    path := filepath.Join(os.Getenv("HOME"), pathCli, fileStore)
    content, err := os.ReadFile(path)
    if err != nil {
        return []models.Task{}, errors.New("Echec read file")
    }

    var data []models.Task
    if err := json.Unmarshal(content, &data); err != nil {
        return []models.Task{}, errors.New("Error décodage JSON:")
    }
    return data, nil
}
