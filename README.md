# 📋 Task Tracker CLI

> **Un gestionnaire de tâches en ligne de commande moderne et performant, développé en Go**

Task Tracker est une application CLI (Command Line Interface) complète permettant de gérer vos tâches quotidiennes directement depuis votre terminal. Construit avec Go et suivant les principes de **Clean Architecture** avec injection de dépendances, ce projet est conçu pour être scalable, maintenable et facile à étendre.

---

## 📑 Table des matières

- [✨ Fonctionnalités](#-fonctionnalités)
- [🏗️ Architecture du projet](#️-architecture-du-projet)
- [🚀 Installation et configuration](#-installation-et-configuration)
  - [Installer Go](#1-installer-go)
  - [Cloner le projet](#2-cloner-le-projet)
  - [Installation des dépendances](#3-installation-des-dépendances)
- [⚙️ Compilation avec Makefile](#️-compilation-avec-makefile)
- [📖 Utilisation de la CLI](#-utilisation-de-la-cli)
  - [Commandes disponibles](#commandes-disponibles)
  - [Options globales](#options-globales)
- [🔍 Guide détaillé des commandes](#-guide-détaillé-des-commandes)
- [💾 Stockage des données](#-stockage-des-données)
- [🧪 Tests](#-tests)
- [🛠️ Technologies utilisées](#️-technologies-utilisées)

---

## ✨ Fonctionnalités

- ✅ **Ajouter** des tâches avec un titre descriptif
- ❌ **Supprimer** des tâches par ID
- 🔄 **Mettre à jour** des tâches (titre, statut de complétion)
- 📋 **Lister** toutes les tâches avec un affichage formaté en tableau
- 🔍 **Filtrer** les tâches (terminées, non terminées, modifiées, non modifiées)
- 🔀 **Trier** les tâches par titre, date de création, date de modification ou statut
- 📊 Affichage élégant avec `tablewriter`
- 🐛 **Mode debug** pour afficher les logs détaillés
- 💾 Persistance des données en **JSON local**
- 🏗️ Architecture clean avec **injection de dépendances**

---

## 🏗️ Architecture du projet

Le projet suit la **Standard Layout Architecture** pour Go :

```
task-tracker/
├── cmd/                      # Point d'entrée de l'application
│   ├── task/
│   │   └── main.go          # Main principal
│   ├── root.go              # Commande racine Cobra
│   ├── add.go               # Commande 'add'
│   ├── delete.go            # Commande 'delete'
│   ├── update.go            # Commande 'update'
│   └── list.go              # Commande 'list'
├── internal/                 # Code interne non exportable
│   ├── service/             # Couche métier (logique)
│   │   ├── service.go
│   │   └── service_test.go
│   ├── store/               # Couche persistance (stockage)
│   │   └── json_store.go
│   └── utils/               # Utilitaires
│       ├── utils.go
│       └── logger.go
├── models/                   # Structures de données
│   └── task.go
├── bin/                      # Binaires compilés
├── Makefile                  # Automatisation de la compilation
├── go.mod                    # Gestion des dépendances
└── go.sum                    # Hashes des dépendances
```

### Principes architecturaux

1. **Separation of Concerns** : chaque couche a une responsabilité unique
   - `cmd/` : interface CLI (Cobra)
   - `internal/service/` : logique métier
   - `internal/store/` : accès aux données (abstraction avec interface)
   - `models/` : structures de données partagées

2. **Injection de dépendances** : le `TaskService` reçoit un `TaskStore` via interface
   ```go
   type TaskStore interface {
       Save(task []models.Task) error
       AllList()([]models.Task, error)
   }
   ```
   Cela permet de changer facilement le mode de stockage (JSON → SQL, Redis, etc.)

3. **Testabilité** : utilisation de mocks pour tester la logique métier indépendamment du stockage

---

## 🚀 Installation et configuration

### 1. Installer Go

#### Sur Linux (Debian/Ubuntu)
```bash
# Télécharger la dernière version de Go (exemple : Go 1.23.0)
wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz

# Supprimer l'ancienne installation (si elle existe)
sudo rm -rf /usr/local/go

# Extraire l'archive
sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz

# Ajouter Go au PATH (ajouter à ~/.bashrc ou ~/.zshrc)
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Recharger le shell
source ~/.bashrc

# Vérifier l'installation
go version
```

#### Sur macOS
```bash
# Avec Homebrew
brew install go

# Vérifier l'installation
go version
```

#### Sur Windows
1. Téléchargez l'installateur MSI depuis [go.dev/dl](https://go.dev/dl/)
2. Suivez l'assistant d'installation
3. Vérifiez avec `go version` dans PowerShell

### 2. Cloner le projet

```bash
# Cloner le dépôt
git clone https://github.com/PassinThomas/task-tracker.git

# Accéder au dossier
cd task-tracker
```

### 3. Installation des dépendances

```bash
# Télécharger les modules Go
go mod download

# Vérifier les dépendances
go mod tidy
```

**Dépendances principales** :
- [`spf13/cobra`](https://github.com/spf13/cobra) v1.10.2 : Framework CLI puissant
- [`olekukonko/tablewriter`](https://github.com/olekukonko/tablewriter) v1.1.3 : Affichage de tableaux ASCII

---

## ⚙️ Compilation avec Makefile

Le projet inclut un `Makefile` pour automatiser les tâches courantes.

### Commandes disponibles

| Commande       | Description                                           | Détails                                      |
|----------------|-------------------------------------------------------|----------------------------------------------|
| `make build`   | Compile le projet et génère le binaire               | Crée `bin/task` exécutable                   |
| `make test`    | Lance tous les tests avec verbose                    | Exécute `go test -v ./...`                   |
| `make run`     | Compile puis exécute `task list`                     | Raccourci pour tester rapidement             |
| `make clean`   | Supprime le dossier `bin/`                           | Nettoie les binaires compilés                |

### Exemples d'utilisation

```bash
# 1. Compiler le projet
make build
# Résultat : bin/task est créé

# 2. Exécuter le binaire
./bin/task --help

# 3. Lancer les tests
make test

# 4. Nettoyer les binaires
make clean
```

### Installation globale (optionnel)

Pour utiliser `task` n'importe où dans votre terminal :

```bash
# Installer le binaire dans $GOPATH/bin
go install ./cmd/task

# Vérifier que $GOPATH/bin est dans votre PATH
echo $PATH | grep $GOPATH/bin

# Si absent, ajouter à ~/.bashrc ou ~/.zshrc
export PATH=$PATH:$GOPATH/bin

# Utiliser la commande globalement
task --help
```

---

## 📖 Utilisation de la CLI

### Syntaxe générale

```bash
task [commande] [arguments] [flags]
```

### Commandes disponibles

| Commande   | Description                                    | Arguments requis         |
|------------|------------------------------------------------|--------------------------|
| `add`      | Ajouter une nouvelle tâche                    | `<titre>`                |
| `delete`   | Supprimer une tâche par son ID                | `<id>`                   |
| `update`   | Mettre à jour une tâche (titre ou statut)     | `<id>` + flags           |
| `list`     | Afficher la liste des tâches                  | aucun                    |
| `help`     | Afficher l'aide générale ou d'une commande    | `[commande]` (optionnel) |

### Options globales

| Flag             | Raccourci | Description                          | Valeur par défaut |
|------------------|-----------|--------------------------------------|-------------------|
| `--debug`        | `-d`      | Active le mode debug (logs détaillés)| `false`           |

**Exemple** :
```bash
task list --debug  # Affiche les logs internes
```

---

## 🔍 Guide détaillé des commandes

### 1️⃣ `task add` - Ajouter une tâche

Créer une nouvelle tâche avec un titre descriptif.

#### Syntaxe
```bash
task add "<titre de la tâche>"
```

#### Exemples
```bash
# Ajouter une tâche simple
task add "Acheter du lait"

# Ajouter une tâche avec des espaces
task add "Préparer la présentation pour lundi"

# Avec mode debug
task add "Finir le projet Go" --debug
```

#### Sortie
```
✓ Task 1 created successfully
```

#### Détails techniques
- L'ID est auto-incrémenté (calcul du maximum des IDs existants + 1)
- `CreatedAt` est défini automatiquement à l'instant présent
- `UpdatedAt` est initialisé à `zero time` (non modifié)
- `Done` est `false` par défaut

---

### 2️⃣ `task delete` - Supprimer une tâche

Supprimer définitivement une tâche par son ID.

#### Syntaxe
```bash
task delete <id>
```

#### Exemples
```bash
# Supprimer la tâche n°3
task delete 3

# Avec debug
task delete 5 --debug
```

#### Sortie
```
✓ Task 3 deleted successfully
```

#### Gestion d'erreurs
- Si l'ID n'existe pas : `Error: task ID 999 not found`
- Si l'ID n'est pas un nombre : `Error: Fail conversion ID of deleteCmd`

---

### 3️⃣ `task update` - Mettre à jour une tâche

Modifier le titre et/ou le statut de complétion d'une tâche.

#### Syntaxe
```bash
task update <id> [--title "<nouveau titre>"] [--complete | --incomplete]
```

#### Flags disponibles

| Flag              | Raccourci | Description                          | Mutually exclusive avec |
|-------------------|-----------|--------------------------------------|-------------------------|
| `--title`         | -         | Modifier le titre de la tâche        | -                       |
| `--complete`      | `-c`      | Marquer la tâche comme terminée      | `--incomplete`          |
| `--incomplete`    | `-i`      | Marquer la tâche comme non terminée  | `--complete`            |

#### Exemples

```bash
# Marquer la tâche 2 comme terminée
task update 2 --complete
task update 2 -c

# Changer le titre de la tâche 1
task update 1 --title "Acheter du pain et du lait"

# Changer titre ET statut
task update 3 --title "Terminer le README" --complete

# Réouvrir une tâche terminée
task update 3 --incomplete
task update 3 -i

# Avec debug
task update 4 --complete --debug
```

#### Sortie
```
✓ Task 2 updated successfully
```

#### Comportement interne
- Si `--complete` ou `--incomplete` est utilisé, `UpdatedAt` est mis à jour
- Si `--title` est fourni, `UpdatedAt` est également mis à jour
- Les flags `--complete` et `--incomplete` sont mutuellement exclusifs (Cobra gère l'erreur)

---

### 4️⃣ `task list` - Afficher les tâches

Lister toutes les tâches sous forme de tableau formaté, avec possibilité de **filtrage** et **tri**.

#### Syntaxe
```bash
task list [--done | --undone] [--updated | --not-updated] [--sort <champ>] [--order <asc|desc>]
```

#### Flags de filtrage

| Flag              | Description                                  | Mutually exclusive avec |
|-------------------|----------------------------------------------|-------------------------|
| `--done`          | Afficher uniquement les tâches terminées     | `--undone`              |
| `--undone`        | Afficher uniquement les tâches non terminées | `--done`                |
| `--updated`       | Afficher uniquement les tâches modifiées     | `--not-updated`         |
| `--not-updated`   | Afficher uniquement les tâches non modifiées | `--updated`             |

#### Flags de tri

| Flag              | Valeurs possibles                      | Description                               |
|-------------------|----------------------------------------|-------------------------------------------|
| `--sort`          | `title`, `created`, `updated`, `status`| Champ de tri                              |
| `--order`         | `asc` (ascendant), `desc` (descendant) | Ordre de tri (défaut : `asc`)             |

#### Exemples

```bash
# Afficher toutes les tâches
task list

# Afficher uniquement les tâches terminées
task list --done

# Afficher uniquement les tâches non terminées
task list --undone

# Afficher les tâches modifiées
task list --updated

# Trier par titre (ordre alphabétique)
task list --sort title

# Trier par titre (ordre inverse)
task list --sort title --order desc

# Trier par date de création (les plus récentes en premier)
task list --sort created --order desc

# Trier par date de modification
task list --sort updated --order asc

# Combiner filtres et tri
task list --done --sort created --order desc

# Trier par statut (done avant undone)
task list --sort status

# Avec debug
task list --debug
```

#### Sortie exemple

```
+----+---------------------------+--------+------------------+------------------+
| ID |           TITLE           | STATUS |    CREATED AT    |    UPDATED AT    |
+----+---------------------------+--------+------------------+------------------+
|  1 | Acheter du lait           | undone | 2026-02-19 10:30 | none             |
|  2 | Préparer présentation     | done   | 2026-02-19 11:00 | 2026-02-19 14:22 |
|  3 | Finir le projet Go        | done   | 2026-02-19 12:15 | 2026-02-19 15:45 |
+----+---------------------------+--------+------------------+------------------+
```

#### Détails techniques sur le tri

Le tri est géré par la fonction `SortList` dans `internal/service/service.go` :

- **`title`** : tri alphabétique (case-insensitive)
- **`created`** : tri par `CreatedAt` (datetime)
- **`updated`** : tri par `UpdatedAt` (datetime)
- **`status`** : tri par statut (`done` avant `undone` en ordre croissant)

Si deux tâches ont la même valeur pour le champ de tri (ex: même date), elles sont triées par ID en ordre croissant.

#### Filtrage combiné

Les filtres sont appliqués **après** le tri. Logique :

```go
// Pseudo-code
tasks = GetAllTasks()
if sort_provided:
    tasks = Sort(tasks, sort_field, order)
tasks = Filter(tasks, filter_options)
```

Si aucun filtre ne correspond, **toutes** les tâches sont retournées.

---

### 5️⃣ `task help` - Aide

Afficher l'aide générale ou l'aide d'une commande spécifique.

#### Syntaxe
```bash
task help [commande]
```

#### Exemples
```bash
# Aide générale
task help

# Aide pour la commande list
task help list

# Aide pour la commande update
task help update
```

---

## 💾 Stockage des données

### Emplacement du fichier JSON

Les tâches sont stockées dans un fichier JSON local :

```
~/.config/.mycli/todo.json
```

Où `~` est votre répertoire utilisateur :
- **Linux/macOS** : `/home/username/.config/.mycli/todo.json`
- **Windows** : `C:\Users\username\.config\.mycli\todo.json`

### Format JSON

```json
[
  {
    "id": 1,
    "title": "Acheter du lait",
    "done": false,
    "created_at": "2026-02-19T10:30:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  },
  {
    "id": 2,
    "title": "Préparer présentation",
    "done": true,
    "created_at": "2026-02-19T11:00:00Z",
    "updated_at": "2026-02-19T14:22:00Z"
  }
]
```

### Architecture de stockage

Le projet utilise une **abstraction par interface** :

```go
type TaskStore interface {
    Save(task []models.Task) error
    AllList()([]models.Task, error)
}
```

Implémentation actuelle : `JsonStore` (fichier JSON local)

**Extensibilité** : pour passer à une base de données, il suffit de :
1. Créer une nouvelle struct implémentant `TaskStore` (ex: `PostgresStore`)
2. Modifier l'injection de dépendances dans `cmd/root.go`

```go
// Exemple pour PostgreSQL
postgresStore := store.NewPostgresStore(connectionString)
taskService = service.NewTaskService(postgresStore)
```

---

## 🧪 Tests

Le projet inclut des tests unitaires pour la couche service.

### Lancer les tests

```bash
# Tous les tests avec output verbose
make test

# Ou directement avec Go
go test -v ./...

# Tests d'un package spécifique
go test -v ./internal/service
```

### Exemple de test (Mock Store)

Le fichier `internal/service/service_test.go` utilise un **mock store** :

```go
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
```

Cela permet de tester la logique métier **sans dépendre du système de fichiers**.

### Cas de tests couverts

- ✅ Ajout de tâche avec auto-incrémentation d'ID
- ✅ Suppression de tâche
- ✅ Mise à jour de tâche (titre, statut)
- ✅ Tri par différents champs
- ✅ Gestion des erreurs (ID inexistant, etc.)

---

## 🛠️ Technologies utilisées

### Langage et runtime
- **Go 1.23.0** : langage compilé, performant, avec gestion native de la concurrence

### Bibliothèques externes

| Dépendance                | Version  | Rôle                                           |
|---------------------------|----------|------------------------------------------------|
| `spf13/cobra`             | v1.10.2  | Framework CLI (commandes, flags, help)         |
| `olekukonko/tablewriter`  | v1.1.3   | Affichage de tableaux ASCII formatés           |
| `log/slog`                | (stdlib) | Logging structuré (mode debug)                 |

### Outils de développement
- **Make** : automatisation de la compilation et des tests
- **go test** : framework de tests intégré
- **go mod** : gestion des dépendances

---

## 🎯 Cas d'usage complets

### Scénario 1 : Workflow quotidien

```bash
# Ajouter des tâches du matin
task add "Répondre aux emails"
task add "Meeting équipe à 10h"
task add "Réviser le code PR #42"

# Afficher la liste
task list

# Marquer une tâche comme terminée
task update 1 --complete

# Afficher seulement les tâches non terminées
task list --undone

# Terminer et modifier une tâche
task update 2 --title "Meeting équipe (reporté à 14h)" --complete

# Afficher les tâches terminées, triées par date
task list --done --sort created --order desc
```

### Scénario 2 : Debug d'un problème

```bash
# Activer le mode debug pour voir les logs
task add "Test debug" --debug

# Voir les détails du tri
task list --sort updated --order desc --debug
```

### Scénario 3 : Nettoyage de tâches anciennes

```bash
# Afficher les tâches modifiées
task list --updated

# Supprimer une tâche obsolète
task delete 5

# Vérifier
task list
```

---

## 🌟 Points forts du projet

1. **Architecture propre** : séparation claire des responsabilités (CLI / Service / Store)
2. **Testable** : injection de dépendances avec interfaces
3. **Extensible** : facile d'ajouter de nouveaux modes de stockage
4. **User-friendly** : messages clairs, tableaux formatés, aide intégrée
5. **Robuste** : gestion d'erreurs complète, validation des entrées
6. **Performant** : binaire compilé, démarrage instantané

