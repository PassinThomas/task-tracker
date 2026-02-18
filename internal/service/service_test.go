package service

import (
    "testing"
    "task/models"
)

//simulation stockage
type MockStore struct {
    tasks []models.Task
}

func (m *MockStore) Save(tasks []models.Task) error {
    m.tasks = tasks
    return nil
}

func (m *MockStore) AllList() ([]models.Task, error) {
    return m.tasks, nil
}

func TestAdd(t *testing.T) {
    mock := &MockStore{tasks: []models.Task{
        {ID: 1, Title: "task1"}, 
        {ID: 2, Title: "task2"},
    }}
    service := NewTaskService(mock)

    task, err := service.Add("task3")
    
    // Verif 1 : Pas d'erreur technique
    if err != nil {
        t.Errorf("Add() failed unexpectedly: %v", err)
    }

    // Verif 2 : Logique metier (ID doit etre 3)
    expectedID := 3
    if task.ID != expectedID {
        t.Errorf("Expected ID %d, got %d", expectedID, task.ID)
    }
    
    if task.Title != "task3" {
         t.Errorf("Expected Title 'task3', got '%s'", task.Title)
    }
}

func TestDeleteNonExistent(t *testing.T) {
    mock := &MockStore{tasks: []models.Task{{ID: 1, Title: "pain"}}}
    service := NewTaskService(mock)

    _, err := service.Delete(99)

    // Verif : On veut une erreur technique
    if err == nil {
        // si aucune erreur -> BUG
        t.Errorf("Test failed: expected an error when deleting ID 99, but got nil (success)")
    } else {
        //debug, erreur reçue
        t.Logf("Success: got expected error -> %v", err)
    }
}

/* TEST UPDATE A IMPLEMENTE */