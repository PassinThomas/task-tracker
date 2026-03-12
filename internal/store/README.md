# 💾 Store Layer - Abstraction du Stockage

## Responsabilité

Le store définit **comment** les données sont persistées.

## Interface TaskStore

```go
type TaskStore interface {
    Save(tasks []models.Task) error
    AllList() ([]models.Task, error)
}
```

Toute implémentation (JSON, SQL, Redis) doit respecter cette interface.

## Implémentation actuelle : JsonStore

**Fichier :** `json_store.go`

**Localisation :** `~/.config/.mycli/todo.json`

Exemple de structure JSON :

```json
[
  {
    "id": 1,
    "title": "Acheter du lait",
    "done": false,
    "created_at": "2026-02-19T10:30:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  }
]
```

## Ajouter une nouvelle implémentation

### Exemple : PostgreSQL

```go
// internal/store/postgres_store.go
package store

import "database/sql"

type PostgresStore struct {
    db *sql.DB
}

func NewPostgresStore(connString string) (*PostgresStore, error) {
    db, err := sql.Open("postgres", connString)
    // ...
    return &PostgresStore{db: db}, nil
}

func (ps *PostgresStore) Save(tasks []models.Task) error {
    // Implémentation PostgreSQL
    tx, _ := ps.db.Begin()
    for _, task := range tasks {
        tx.Exec("INSERT INTO tasks (id, title, done) VALUES (?, ?, ?)",
            task.ID, task.Title, task.Done)
    }
    tx.Commit()
    return nil
}

func (ps *PostgresStore) AllList() ([]models.Task, error) {
    rows, _ := ps.db.Query("SELECT id, title, done, created_at, updated_at FROM tasks")
    // ...
}
```

### Adapter cmd/root.go

```go
// cmd/root.go
var taskService *service.TaskService

func init() {
    // Option 1 : JSON (défaut)
    store := store.NewJsonStore()
    
    // Option 2 : PostgreSQL
    // store, _ := store.NewPostgresStore("postgres://...")
    
    taskService = service.NewTaskService(store)
}
```

## Migrations

Pour passer de JSON à PostgreSQL, prévoir une étape de migration :

```go
// tools/migrate.go
func MigrateJsonToPostgres(jsonPath, pgConnString string) error {
    // Lire depuis JSON
    jsonStore := store.NewJsonStore(jsonPath)
    tasks, _ := jsonStore.AllList()
    
    // Écrire vers PostgreSQL
    pgStore, _ := store.NewPostgresStore(pgConnString)
    return pgStore.Save(tasks)
}
```
