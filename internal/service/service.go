package service

import (
	"fmt"
	"strings"
	"errors"
	"time"
	"sort"
	"math"
	
	"github.com/PassinThomas/task-tracker/models"
	"github.com/PassinThomas/task-tracker/internal/utils"
)

var ErrIDLimitReached = errors.New("task ID limit reached: cannot add more tasks")

type TaskStore interface {
	Save(task []models.Task) error
	AllList()([]models.Task, error)
}

type TaskService struct {
	store TaskStore
}

func NewTaskService(store TaskStore) *TaskService {
	return &TaskService{store: store}
}

func (s *TaskService) Add(title string) (*models.Task, error) {
	errParse := utils.ParseStr(title)
	if errParse != nil {
		return &models.Task{}, fmt.Errorf("Bad format: %w", errParse)
	}

	if len(title) < 1 {
		return &models.Task{}, errors.New("Title must not be empty")
	}

	tasks, _ := s.store.AllList()
    newID := 1
    for _, t := range tasks {
		if t.ID == math.MaxInt {
            return nil, ErrIDLimitReached
        }
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
	
	if err := s.store.Save(tasks); err != nil {
		return &models.Task{}, fmt.Errorf("failed to save changes to disk: %w", err)
	}

	return &task, nil
}

func (s *TaskService) Delete(id int) (*models.Task, error) {
	if id < 1 {
		return &models.Task{}, errors.New("Id does not exist")
	}
	
	tasks, err := s.store.AllList()
	if err != nil {
		return &models.Task{}, fmt.Errorf("Cannot delete empty task %w", err)
	}

	var task []models.Task
	var delTask models.Task
	isChanged := false

	for _, t := range tasks {
		if t.ID == id {
			delTask = t
			isChanged = true
			continue ;
		} else {
			task = append(task, t) 
		}
	}

	if !isChanged {
		return &models.Task{}, errors.New("ID does not match any task")
	}

	err = s.store.Save(task)
	if err != nil {
		return &models.Task{}, fmt.Errorf("failed to save changes to disk: %w", err)
	}

	return &delTask, nil
}

func (s *TaskService) Update(id int, flg models.FlgUpdate) (*models.Task, error) {
	if flg.NewTitle == "" && !flg.Done && !flg.NotDone {
		return &models.Task{}, errors.New("Bad flags input")
	}

	if flg.NewTitle != "" {
		err := utils.ParseStr(flg.NewTitle)
		if err != nil {
			return &models.Task{}, fmt.Errorf("Update title failed: %w", err)
		}
	}
	
	if id < 1 {
		return &models.Task{}, fmt.Errorf("ID task does not exist")
	}
	
	tasks, err := s.store.AllList()
	if err != nil {
		return &models.Task{}, fmt.Errorf("Cannot update empty task %w", err)
	}
	
	if flg.NewTitle == "" && flg.Done == false && flg.NotDone == false {
        return &models.Task{}, errors.New("No changes requested")
    }

	var task []models.Task
	var taskUpt models.Task
	found := false

	for _, t := range tasks {
		
		if t.ID == id {

			found = true
			if flg.NewTitle != "" && flg.NewTitle != t.Title {
				t.Title = flg.NewTitle
				t.UpdatedAt = time.Now()
				taskUpt = t
			}

			if flg.Done {
				t.Done = flg.Done
				t.UpdatedAt = time.Now()
				taskUpt = t
			}

			if flg.NotDone {
				t.Done = false
				t.UpdatedAt = time.Now()
				taskUpt = t
			}
		}

		task = append(task, t)
	}

	if !found {
        return &models.Task{}, fmt.Errorf("task with ID %d not found", id)
    }
	
	err = s.store.Save(task)
	if err != nil {
		return &models.Task{}, fmt.Errorf("failed to save changes to disk: %w", err)
	}
	
	return &taskUpt, nil
}


func (s *TaskService) List() ([]models.Task, error) {
	tasks, err := s.store.AllList()
	if err != nil {
		return []models.Task{}, fmt.Errorf("Failed get All tasks %w", err)
	}

	return tasks, nil
}

func (s *TaskService) SortList(tasks []models.Task, sorting map[string]string) ([]models.Task, error) {
	task := make([]models.Task, len(tasks))
	copy(task, tasks)

	field := ""
	order := "asc"

	for k, v := range sorting {
		field = k

		if err := utils.CheckField(field); err != nil {
			return nil, fmt.Errorf("flag option undefined: %w", err)
		}

		if v != "" && v != "asc" && v != "desc" {
			return nil, errors.New("bad order flag or undefined")
		}

		if v != "" {
			order = v
		}
	}

	sort.Slice(task, func(i, j int) bool {
		switch field {

		case "title":
			if order == "asc" {
				return strings.ToLower(task[i].Title) < strings.ToLower(task[j].Title)
			}
			return strings.ToLower(task[i].Title) > strings.ToLower(task[j].Title)

		case "created":
			if task[i].CreatedAt.Equal(task[j].CreatedAt) {
				if order == "asc" {
					return task[i].ID < task[j].ID
				}
				return task[i].ID > task[j].ID
			}

			if order == "asc" {
				return task[i].CreatedAt.Before(task[j].CreatedAt)
			}
			return task[i].CreatedAt.After(task[j].CreatedAt)

		case "updated":
			if task[i].UpdatedAt.Equal(task[j].UpdatedAt) {
				if order == "asc" {
					return task[i].ID < task[j].ID
				}
				return task[i].ID > task[j].ID
			}

			if order == "asc" {
				return task[i].UpdatedAt.Before(task[j].UpdatedAt)
			}
			return task[i].UpdatedAt.After(task[j].UpdatedAt)

		default:
			return false
		}
	})
	
	return task, nil
}


func (s *TaskService) ListWithOptions(opts models.ListOptions) ([]models.Task, error) {
	tasks, err := s.store.AllList()
	if err != nil {
		return nil, err
	}

	if opts.Sort != "" {
		sortOpt := map[string]string{
			opts.Sort: opts.Order,
		}

		tasks, err = s.SortList(tasks, sortOpt)
		
		if err != nil {
			return nil, err
		}
	}

	tmp := tasks
	tasks = utils.Filter(tasks, opts.Filter)
	if len(tasks) == 0 {
		tasks = tmp
	}

	return tasks, nil
}