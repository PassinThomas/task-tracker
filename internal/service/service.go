package service

import (
	"fmt"
	"strings"
	"errors"
	"time"
	"sort"
	
	"task/models"
	"task/internal/utils"
	"task/internal/store"

)

func Add(title string) (*models.Task, error) {

	tasks, _ := store.AllList()
    newID := 1
    for _, t := range tasks {
        if t.ID >= newID {
            newID = t.ID + 1
        }
    }

	task := models.Task{
		ID:			newID,	
		Title:		strings.ToLower(title),
		Done:		false,
		CreatedAt:	time.Now(),
	}

	tasks = append(tasks, task)
	
	if err := store.Save(tasks); err != nil {
		return &models.Task{}, fmt.Errorf("failed to save changes to disk: %w", err)
	}

	return &task, nil
}

func Delete(id int) (*models.Task, error) {
	tasks, err := store.AllList()
	if err != nil {
		return &models.Task{}, fmt.Errorf("Cannot delete empty task %w", err)
	}

	var task []models.Task
	var delTask models.Task
	for _, t := range tasks {
		if t.ID == id {
			delTask = t
			continue ;
		} else {
			task = append(task, t) 
		}
	}
	err = store.Save(task)
	if err != nil {
		return &models.Task{}, fmt.Errorf("failed to save changes to disk: %w", err)
	}

	return &delTask, nil
}

func Update(id int, markDone *bool, newTitle *string) (*models.Task, error) {
	tasks, err := store.AllList()
	if err != nil {
		return &models.Task{}, fmt.Errorf("Cannot update empty task %w", err)
	}
	
	if newTitle == nil && markDone == nil {
        return &models.Task{}, errors.New("No changes requested")
    }

	var task []models.Task
	var taskUpt models.Task
	for _, t := range tasks {
		if t.ID == id {
			if newTitle != nil && *newTitle != "" && *newTitle != t.Title {
				t.Title = *newTitle
				taskUpt = t
				t.UpdatedAt = time.Now()
			}
			if markDone != nil && t.Done != *markDone {
				t.Done = *markDone
				taskUpt = t
				t.UpdatedAt = time.Now()
			}
			
		}
		task = append(task, t)
	}
	
	err = store.Save(task)
	if err != nil {
		return &models.Task{}, fmt.Errorf("failed to save changes to disk: %w", err)
	}
	return &taskUpt, nil
}


func List() ([]models.Task, error) {
	tasks, err := store.AllList()
	if err != nil {
		return []models.Task{}, fmt.Errorf("Failed get All tasks %w", err)
	}

	return tasks, nil
}

func SortList(tasks []models.Task, sorting map[string]string) ([]models.Task, error) {
	
	task, field := []models.Task{}, ""
	for _, t := range tasks {
		task = append(task, t)
	}
	for k, v := range sorting {
		field = k
		err := utils.CheckField(field)
		if err != nil {
			return []models.Task{}, fmt.Errorf("Flag option undefined: %w", err)
		}
		if v != "asc" && v != "" && v != "desc" {
			return []models.Task{}, errors.New("Bad order flag or undefined")
		}
	}
	
	sort.Slice(task, func(i, j int)bool {
		switch field {
			case "title":
				if sorting[field] == "asc" || sorting[field] == "" {
					return strings.ToLower(task[i].Title) < strings.ToLower(task[j].Title)	
				} else {
					return strings.ToLower(task[i].Title) > strings.ToLower(task[j].Title)	
				}
			
			case "created":
				if sorting[field] == "asc" || sorting[field] == "" {
					if task[i].CreatedAt.Format("2006-01-02 15:04") == task[j].CreatedAt.Format("2006-01-02 15:04") {
						return task[i].ID < task[j].ID
					}
					return task[i].CreatedAt.Format("2006-01-02 15:04") < task[j].CreatedAt.Format("2006-01-02 15:04")
				} else {
					if task[i].CreatedAt.Format("2006-01-02 15:04") == task[j].CreatedAt.Format("2006-01-02 15:04") {
						return task[i].ID > task[j].ID
					}
					return task[i].CreatedAt.Format("2006-01-02 15:04") > task[j].CreatedAt.Format("2006-01-02 15:04")
				}

			case "updated":
				if sorting[field] == "desc" || sorting[field] == "" {
					if task[i].UpdatedAt.Format("2006-01-02 15:04") == task[j].UpdatedAt.Format("2006-01-02 15:04") {
						return task[i].ID < task[j].ID
					}
					return task[i].UpdatedAt.Format("2006-01-02 15:04") > task[j].UpdatedAt.Format("2006-01-02 15:04")
				} else {
					if task[i].UpdatedAt.Format("2006-01-02 15:04") == task[j].UpdatedAt.Format("2006-01-02 15:04") {
						return task[i].ID > task[j].ID
					}
					return task[i].UpdatedAt.Format("2006-01-02 15:04") < task[j].UpdatedAt.Format("2006-01-02 15:04")
				}
			
			default:
				return false
		}
	})
	copy(tasks, task)

	return tasks, nil
}