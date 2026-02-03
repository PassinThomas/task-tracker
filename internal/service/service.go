package service

import (
	"fmt"
	"errors"
	"time"
	
	"task/models"
	"task/internal/util"
	"task/internal/store"

	"github.com/google/uuid"
)

func Add(title string) error {

	task := models.Task{
		ID:			uuid.New(),	
		Title:		title,
		Done:		false,
		CreatedAt:	time.Now(),
	}

	s := fmt.Sprintf(
    	"%-12s: %v\n%-12s: %v\n%-12s: %v\n%-12s: %v",
    	"ID", task.ID,
    	"Title", task.Title,
    	"Done", task.Done,
    	"CreatedAt", task.CreatedAt.Format("2006-01-02 15:04:05"),
	)

	if err := store.Save(task); err != nil {
		return errors.New("save failed")
	} 
	util.Vlog(util.Verbose, s)
	return nil
}

// func Delete(title string) error {
	
// }

// func update(title string) error {

// }

// func list(val ...any) {

// }