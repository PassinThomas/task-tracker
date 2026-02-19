package service

import (
    "testing"
    "github.com/PassinThomas/task-tracker/models"
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

func TestUpdateDone(t *testing.T) {
    mock := &MockStore{tasks: []models.Task{
        {ID: 1, Title: "task1"},
        {ID: 2, Title: "task2", Done: true},
    }}
    
    // On utilise directement la valeur
    flg := models.FlgUpdate{Done: true} 
    service := NewTaskService(mock)
    
    cp := &MockStore{
        tasks: make([]models.Task, len(mock.tasks)),
    }
    copy(cp.tasks, mock.tasks)
    
    task, err := service.Update(1, flg)
    // 1. ARRÊT IMMÉDIAT EN CAS D'ERREUR (Fatalf)
    if err != nil {
        t.Fatalf("Update failed: %v", err) // Stoppe le test ici
    }
    
    t.Logf("UPDATE TEST: %v", task.UpdatedAt.Format("2006-01-02 15:04"))
    
    // 2. Vérification de la valeur
    if task.Done == cp.tasks[0].Done {
        t.Errorf("Old value unchanged Done: %v", task.Done)
    } else {
        t.Logf("old value: %v -> new value: %v", cp.tasks[0].Done, task.Done)
    }
    
    // 3. Vérification de la date
    if task.UpdatedAt.Equal(cp.tasks[0].UpdatedAt) {
        t.Errorf("new UpdatedAt %v equal with old updated %v", task.UpdatedAt.Format("2006-01-02 15:04"), cp.tasks[0].UpdatedAt.Format("2006-01-02 15:04"))
    } else {
        t.Logf("Date updated successfully")
    }
}

func TestUpdateNotDone(t *testing.T) {
    mock := &MockStore{tasks: []models.Task{
        {ID: 1, Title: "task1"},
        {ID: 2, Title: "task2", Done: true}, 
    }}
    
    flg := models.FlgUpdate{NotDone: true}
    service := NewTaskService(mock)

    cp := &MockStore{
        tasks: make([]models.Task, len(mock.tasks)),
    }
    copy(cp.tasks, mock.tasks)

    task, err := service.Update(2, flg)
    
    // 1. ARRÊT IMMÉDIAT (Fatalf)
    if err != nil {
        t.Fatalf("Update failed: %v", err) // Stoppe le test ici
    }

    t.Logf("UPDATE NotDone TEST: %v", task.UpdatedAt.Format("2006-01-02 15:04"))

    // 2. Vérification de la valeur
    if task.Done == true {
        t.Errorf("Task should be Not Done (false), but got Done: %v", task.Done)
    } else if task.Done == cp.tasks[1].Done {
        t.Errorf("Value unchanged. Old: %v, New: %v", cp.tasks[1].Done, task.Done)
    } else {
        t.Logf("Success: Old value: %v -> New value: %v", cp.tasks[1].Done, task.Done)
    }

    // 3. Vérification de la date
    if task.UpdatedAt.Equal(cp.tasks[1].UpdatedAt) {
        t.Errorf("UpdatedAt did not change. Old: %v, New: %v", 
            cp.tasks[1].UpdatedAt, task.UpdatedAt)
    } else {
        t.Logf("Date updated successfully")
    }
}