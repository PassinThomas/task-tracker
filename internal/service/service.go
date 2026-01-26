package service

import (
	"fmt"
	"strings"
	"errors"

	"task/models"
)

func Add(title string) error {
	if strings.TrimSpace(title) == "" {
	    return errors.New("Input cannot be empty.")
	}
	m := &models.Task{
		Title: title,
	}
	fmt.Printf("%v", m.Title)
	return nil
}