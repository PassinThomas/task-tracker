package service

import (
	"fmt"
	"strings"
	"errors"
	"time"
	
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
			}
			if markDone != nil && t.Done != *markDone {
				t.Done = *markDone
				taskUpt = t
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


func List(opt string, sorting string) ([]models.Task, error) {
	tasks, err := store.AllList()
	if err != nil {
		return []models.Task{}, fmt.Errorf("Failed get All tasks %w", err)
	}
	
	if sorting != "" {
		if sorting != "date" && sorting != "title" && sorting != "status" {
			return []models.Task{}, errors.New("Unknown option")
		}
		utils.SortingTask(sorting, tasks)
	}

	var t []models.Task
    for _, task := range tasks {
        status := "not-done"
        if task.Done {
            status = "done"
        }
		if opt != "" && opt != status{
			continue ;
		}
		t = append(t, task)
	}

	return t, nil
}