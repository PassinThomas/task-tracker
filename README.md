# 📋 Task Tracker CLI

> Un gestionnaire de tâches en ligne de commande moderne et performant, développé en Go

**Table des matières**
- [Quick Start](#quick-start)
- [Documentation complète](./docs/INSTALLATION.md)
- [Guide des commandes](./cmd/README.md)
- [Architecture](./docs/ARCHITECTURE.md)

## Quick Start

### Installation rapide
\`\`\`bash
git clone https://github.com/PassinThomas/task-tracker.git
cd task-tracker
make build
./bin/task --help
\`\`\`

### 5 commandes de base
\`\`\`bash
task add "Ma première tâche"
task list
task update 1 --complete
task list --done
task delete 1
\`\`\`

## 📚 Documentation complète
- **[Installation détaillée](./docs/INSTALLATION.md)** : Go, dépendances, Makefile
- **[Guide des commandes](./cmd/README.md)** : toutes les commandes et flags
- **[Architecture](./docs/ARCHITECTURE.md)** : structure du projet, design patterns
- **[Développement interne](./internal/README.md)** : service, store, modèles

## ✨ Fonctionnalités principales
- ✅ Ajouter/Supprimer/Mettre à jour des tâches
- 📋 Lister avec filtrage et tri
- 💾 Persistance JSON locale
- 🏗️ Architecture clean (injection de dépendances)
- 🧪 Tests unitaires

## 🌟 Points forts
- Architecture modulaire et testable
- Facile à étendre (ajouter une BD, une API, etc.)
- Interface CLI intuitive et bien documentée
