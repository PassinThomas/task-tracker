package store

import (
	"fmt"
	"os"
	"encoding/json"
	"errors"

	"task/models"
)

func checkFile(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		return false
	}
	return true
}

func Save(task models.Task) error {
    home, err := os.UserHomeDir()
    if err != nil {
        return err
    }

    configPath := home + "/.config/.mycli"
    if err := os.MkdirAll(configPath, 0750); err != nil {
        return err
    }

    filePath := configPath + "/todo.json"

    var tasks []models.Task

    if checkFile(filePath) {
        read, err := os.ReadFile(filePath)
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

    return os.WriteFile(filePath, data, 0644)
}

func List() {
    filePath := "/home/tpassin/.config/.mycli/todo.json"
    content, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println("Echec read file:", err)
        return
    }

    var data models.Task
    if err := json.Unmarshal(content, &data); err != nil {
        fmt.Println("Error décodage JSON:", err)
        return
    }

    fmt.Println(data)
}
