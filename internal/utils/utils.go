package utils

import (
	"os"
	"fmt"
	"strings"
	"errors"

	"path/filepath"
	
	"task/models"
)

const (
	date = "date"
	title = "title"
	status = "status"
    pathCli = "/.config/.mycli"
    fileStore = "/todo.json"
)

var (
	DebugVar bool
	OptSort = make(map[string]string)
)

func ParseStr(s string) error {
	if strings.TrimSpace(s) == "" {
		return errors.New("Input cannot be empty.")
	}
	if len([]rune(s)) > 255 {
        return errors.New("Title too long")
    }
	if strings.ContainsAny(s, "/\\") {
        return errors.New("Title cannot contain slashes")
    }
	return nil
}

// func SortingTask(sorting string, tasks []models.Task) {

// 	var task []models.Task
// 	for _, t := range tasks {
// 		task = append(task, t)
// 	}
// 	sort.Slice(task, func(i, j int)bool {
// 		if sorting == title {
// 			return strings.ToLower(task[i].Title) < strings.ToLower(task[j].Title)
// 		}
// 		if sorting == date {
// 			return task[i].CreatedAt.String() > task[j].CreatedAt.String()
// 		} else if sorting == status {
//     		statusI := "not-done"
//     		statusJ := "not-done"
//     		if task[i].Done {
//     		    statusI = "done"
//     		}
//     		if task[j].Done {
//     		    statusJ = "done"
//     		}
//     		return statusI < statusJ
// 		}
// 		return false
// 	})
// 	copy(tasks, task)
// } 

func CheckField(field string) error {

	validFields := map[string]bool{
		"title": true,
		"created": true,
		"updated": true,
	}
	
	if !validFields[field] {
		return errors.New("Unknown sort field")
	}
	return nil
}


func GeneratePath() (string, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }

    configPath := filepath.Join(home, pathCli)
    if err := os.MkdirAll(configPath, 0750); err != nil {
        return "", err
    }

    path := filepath.Join(configPath, fileStore)
    return path, nil
}

func UpadtedFormat(task models.Task) string {
	if task.CreatedAt.Format("2006-01-02 15:04") > task.UpdatedAt.Format("2006-01-02 15:04") {
		return fmt.Sprintf("none")
	} else {
		return task.UpdatedAt.Format("2006-01-02 15:04")
	}
}