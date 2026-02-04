package service

import (
	"fmt"
	"errors"
	"time"
	"os"
	
	"task/models"
	"task/internal/util"
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
		Title:		title,
		Done:		false,
		CreatedAt:	time.Now(),
	}

	
	if err := store.Save(task); err != nil {
		return errors.New("save failed")
	}

	s := fmt.Sprintf(
		"%-12s: %v\n%-12s: %v\n%-12s: %v\n%-12s: %v",
		"ID", task.ID,
		"Title", task.Title,
		"Done", task.Done,
		"CreatedAt", task.CreatedAt.Format("2006-01-02 15:04"),
	)

	util.Vlog(util.Verbose, s)
	return nil
}

// func Delete(title string) error {
	
// }

// func update(title string) error {

// }

func List(opt *string) error {
	tasks, err := store.AllList()
	if err != nil {
		return errors.New("Failed display tasks")
	}
	
	table := tablewriter.NewWriter(os.Stdout)
    table.Header("ID", "Title", "Status", "Created At")
    for _, task := range tasks {
        status := "not-done"
        if task.Done {
            status = "done"
        }
		if opt != nil && *opt != status{
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