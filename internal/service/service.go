package service

import (
	"fmt"
	// "errors"
	"time"
	
	"task/models"
	"task/internal/util"
	"task/internal/store"

	"github.com/google/uuid"
)

func Add(title string) error {

	err := util.ParseStr(title)
	if err != nil {
		util.Vlog(util.Verbose, fmt.Sprintf("%v", err))
		return err
	}

	task := &models.Task{
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

	if err = store.CreateTask(task); err != nil {
		return err
	} 
	
	util.Vlog(util.Verbose, s)
	return nil
}