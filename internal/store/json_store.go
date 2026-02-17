package store

import (
	"os"
	"encoding/json"
	"errors"
    
    "task/internal/utils"
	"task/models"
)

type JsonStore struct {
}

func NewJsonStore() *JsonStore {
    return &JsonStore{}
}

func (j *JsonStore) Save(tasks []models.Task) error {
   
    path, err := utils.GeneratePath()
    if err != nil {
        return errors.New("Failed to generate path file")
    }

    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(path, data, 0644)
}


func (j *JsonStore) AllList() ([]models.Task, error)  {
    path, errPath := utils.GeneratePath()
    if errPath != nil {
        return []models.Task{}, errors.New("Generate path failed")
    }

    content, err := os.ReadFile(path)
    if err != nil {
        return []models.Task{}, errors.New("Echec read file")
    }

    var data []models.Task
    if err := json.Unmarshal(content, &data); err != nil {
        return []models.Task{}, errors.New("Error décodage JSON:")
    }

    return data, nil
}
