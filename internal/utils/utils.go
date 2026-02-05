package utils

import (
	"fmt"
	"strings"
	"errors"
	"sort"
	"task/models"
)

const (
	date = "date"
	title = "title"
	status = "status"
)

var (
	Verbose bool
)

func Vlog(v bool, s string) {
	if v == true {
		fmt.Println(s)
	}
	return
}

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

func SortingTask(sorting string, tasks []models.Task) {

	var task []models.Task
	for _, t := range tasks {
		task = append(task, t)
	}
	sort.Slice(task, func(i, j int)bool {
		if sorting == "title" {
			return task[i].Title < task[j].Title
		}
		if sorting == "date" {
			return task[i].CreatedAt.String() > task[j].CreatedAt.String()
		} else if sorting == "status" {
    		statusI := "not-done"
    		statusJ := "not-done"
    		if task[i].Done {
    		    statusI = "done"
    		}
    		if task[j].Done {
    		    statusJ = "done"
    		}
    		return statusI < statusJ
		}
		return false
	})
	copy(tasks, task)
} 