package service

import (
	"fmt"
	"strings"
	"errors"
	"time"
	"os"
	
	"task/models"
	"task/internal/utils"
	"task/internal/store"

	"github.com/olekukonko/tablewriter"
)

func Add(title string) error {

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
		return errors.New("save failed")
	}

	s := fmt.Sprintf(
		"%-12s: %v\n%-12s: %v\n%-12s: %v\n%-12s: %v",
		"ID", task.ID,
		"Title", task.Title,
		"Done", task.Done,
		"CreatedAt", task.CreatedAt.Format("2006-01-02 15:04"),
	)

	utils.Vlog(s)
	return nil
}

func Delete(title string) error {
	tasks, err := store.AllList()
	if err != nil {
		return errors.New("Cannot delete empty task")
	}

	var task []models.Task
	for _, t := range tasks {
		if t.Title == title {
			continue ;
		} else {
			task = append(task, t) 
		}
	}
	return store.Save(task)	
}

func Update(id int, markDone *bool, newTitle *string) error {
	tasks, err := store.AllList()
	if err != nil {
		return errors.New("Cannot update empty task")
	}
	
	if newTitle == nil && markDone == nil {
        return fmt.Errorf("No changes requested")
    }

	var task []models.Task
	for _, t := range tasks {
		if t.ID == id {
			if newTitle != nil && *newTitle != "" && *newTitle != t.Title {
				t.Title = *newTitle
			}
			if markDone != nil && t.Done != *markDone {
				t.Done = *markDone
			}
		}
		task = append(task, t)
	}
	return store.Save(task)	

}


func List(opt string, sorting string) error {
	tasks, err := store.AllList()
	if err != nil {
		return errors.New("Failed display tasks")
	}
	
	if sorting != "" {
		if sorting != "date" && sorting != "title" && sorting != "status" {
			return errors.New("Unknown option")
		}
		utils.SortingTask(sorting, tasks)
	}
	
	table := tablewriter.NewWriter(os.Stdout)
    table.Header([]string{"ID", "Title", "Status", "Created At"})
    for _, task := range tasks {
        status := "not-done"
        if task.Done {
            status = "done"
        }
		if opt != "" && opt != status{
			continue ;
		}
        table.Append([]string{
            fmt.Sprintf("%v", task.ID),
            task.Title,
            status,
            task.CreatedAt.Format("2006-01-02 15:04"),
        })
    }
    table.Render()
	return nil
}