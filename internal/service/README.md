# 🧠 Service Layer - Logique Métier

## Responsabilité

Le service (`TaskService`) contient toute la logique métier :
- Validation des données
- Calcul des IDs
- Tri et filtrage
- Gestion des dates

**Le service N'A PAS le droit de** :
- Lire/écrire directement dans les fichiers
- Connaître les détails de stockage
- Gérer l'interface utilisateur

## Architecture

\`\`\`go
type TaskService struct {
    store TaskStore  // Injection de dépendance
}

type TaskStore interface {
    Save(tasks []models.Task) error
    AllList() ([]models.Task, error)
}
\`\`\`

## Principales fonctions

### `AddTask(title string) error`
- Crée une tâche avec un ID auto-incrémenté
- Valide le titre (non vide)
- Sauvegarde via le store

### `DeleteTask(id int) error`
- Supprime par ID
- Retourne erreur si ID inexistant

### `UpdateTask(id int, ...) error`
- Modifie titre et/ou statut
- Met à jour `UpdatedAt` si modification

### `GetTasks(filters...) []Task`
- Récupère les tâches
- Applique filtres (done, updated, etc.)
- Applique tri (title, created, updated, status)
- Applique ordre (asc, desc)

### `SortList(tasks []Task, field string, order string) []Task`
Tri par :
- `title` : alphabétique (case-insensitive)
- `created` : date de création
- `updated` : date de modification
- `status` : done avant undone

## Tests

Le service est testé avec un `MockStore` :

\`\`\`go
type MockStore struct {
    tasks []models.Task
}

// Les tests sont indépendants du stockage réel
func TestAddTask(t *testing.T) {
    store := &MockStore{}
    service := NewTaskService(store)
    // ...
}
\`\`\`

## Extensibilité

Pour ajouter une nouvelle fonctionnalité :

1. Ajouter la fonction au `TaskService`
2. Écrire les tests avec le `MockStore`
3. Créer une nouvelle commande dans `cmd/`

Exemple : ajouter la recherche par titre
\`\`\`go
func (ts *TaskService) SearchByTitle(query string) []models.Task {
    // Logique indépendante du stockage
}
\`\`\`