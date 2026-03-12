# 📖 Guide des Commandes

## Syntaxe générale
```bash
task [commande] [arguments] [flags]
```

## Vue d'ensemble

| Commande | Description |
|----------|-------------|
| `add` | Ajouter une tâche |
| `delete` | Supprimer une tâche |
| `update` | Mettre à jour une tâche |
| `list` | Lister les tâches |
| `help` | Afficher l'aide |

---

## 1️⃣ `task add` - Ajouter une tâche

### Syntaxe
```bash
task add "<titre de la tâche>"
```

### Exemples
```bash
task add "Acheter du lait"
task add "Préparer la présentation" --debug
```

### Détails techniques
- ID auto-incrémenté
- `CreatedAt` = maintenant
- `UpdatedAt` = zéro
- `Done` = false

[Continuer avec les autres commandes...]

---

## Options globales

| Flag | Raccourci | Description |
|------|-----------|-------------|
| `--debug` | `-d` | Mode debug (logs détaillés) |

---

## Cas d'usage complets

### Workflow quotidien
```bash
task add "Répondre aux emails"
task list
task update 1 --complete
task list --undone
```

### Debug
```bash
task add "Test" --debug
task list --sort updated --debug
\`\`\`
