# 🏗️ Architecture du Projet

## Vue d'ensemble

```
┌─────────────────────────────────────────┐
│ CLI Layer (Cobra)                       │
│ cmd/: add, delete, update, list        │
└────────────────┬────────────────────────┘
                 │
┌────────────────▼────────────────────────┐
│ Service Layer (Métier)                  │
│ internal/service/: logique métier       │
└────────────────┬────────────────────────┘
                 │
┌────────────────▼────────────────────────┐
│ Store Layer (Abstraction)               │
│ internal/store/: interface TaskStore    │
└────────────────┬────────────────────────┘
                 │
┌────────────────▼────────────────────────┐
│ Persistence Layer                       │
│ JSON, PostgreSQL, Redis, etc.          │
└─────────────────────────────────────────┘
```

## Principes de conception

### 1. Separation of Concerns
- **cmd/** : Interface utilisateur uniquement
- **internal/service/** : Logique métier (indépendante du stockage)
- **internal/store/** : Abstraction du stockage
- **models/** : Structures partagées

### 2. Injection de dépendances
```go
type TaskStore interface {
    Save(tasks []models.Task) error
    AllList() ([]models.Task, error)
}

type TaskService struct {
    store TaskStore
}
```

**Avantages :**
- Testabilité : utiliser un mock store
- Extensibilité : changer de BDD sans toucher à la logique

### 3. Testabilité
- Les tests unitaires utilisent un 
- Aucune dépendance au système de fichiers dans les tests métier

## Détails par couche

Voir les README spécifiques :
- [cmd/README.md](../../cmd/README.md) : Interface CLI
- [internal/service/README.md](../../internal/service/README.md) : Logique métier
- [internal/store/README.md](../../internal/store/README.md) : Abstraction stockage
- [models/README.md](../../models/README.md) : Structures de données

## Extensibilité future

### Ajouter une nouvelle persistance (PostgreSQL)

```go
// 1. Créer internal/store/postgres_store.go
type PostgresStore struct {
    db *sql.DB
}

func (p *PostgresStore) Save(tasks []models.Task) error {
    // Implémentation PostgreSQL
}

// 2. Modifier cmd/root.go pour injecter la bonne implémentation
store := store.NewPostgresStore(connString)
taskService := service.NewTaskService(store)
```

### Ajouter une nouvelle commande

Voir [cmd/README.md](../../cmd/README.md)
