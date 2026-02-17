package utils

import (
	"os"
	"fmt"
	"strings"
	"errors"
	"path/filepath"
	
	"task/models"

	"github.com/olekukonko/tablewriter"
)

const (
	date = "date"
	title = "title"
	status = "status"
    pathCli = "/.config/.mycli"
    fileStore = "/todo.json"
)

var DebugVar bool

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

func Filter(tasks []models.Task, filter models.FilterOptions) []models.Task {
	var filtered []models.Task

    for _, task := range tasks {
        if filter.Done && !task.Done {
            continue
        }
        if filter.Undone && task.Done {
            continue
        }
        if filter.Updated && task.UpdatedAt.IsZero() {
            continue
        }
        if filter.NotUpdated && !task.UpdatedAt.IsZero() {
            continue
        }
        filtered = append(filtered, task)
    }

    return filtered
}

func RenderTasks(tasks []models.Task) {

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Title", "Status", "Created At", "Updated At"})

	for _, task := range tasks {

		status := "undone"
		if task.Done {
			status = "done"
		}

		table.Append([]string{
			fmt.Sprintf("%v", task.ID),
			task.Title,
			status,
			task.CreatedAt.Format("2006-01-02 15:04"),
			UpadtedFormat(task),
		})
	}

	table.Render()
}
