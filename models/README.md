# 📊 Models - Structures de Données

## Task

La structure principale du projet.

\`\`\`go
type Task struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Done      bool      `json:"done"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
\`\`\`

### Champs

| Champ | Type | Description |
|-------|------|-------------|
| `ID` | `int` | Identifiant unique, auto-incrémenté |
| `Title` | `string` | Description de la tâche |
| `Done` | `bool` | État de complétion (false = à faire) |
| `CreatedAt` | `time.Time` | Date de création (immutable) |
| `UpdatedAt` | `time.Time` | Date de la dernière modification |

### Remarques sur UpdatedAt

- Initialisé à **zéro time** lors de la création : `0001-01-01T00:00:00Z`
- Affiché comme `"none"` dans la CLI si zéro
- Mis à jour à chaque modification

### JSON Marshaling

Les tags `json:` permettent de contrôler la sérialisation :

\`\`\`json
{
  "id": 1,
  "title": "Acheter du lait",
  "done": false,
  "created_at": "2026-02-19T10:30:00Z",
  "updated_at": "0001-01-01T00:00:00Z"
}
\`\`\`