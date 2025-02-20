# TO DO List Management with Go/Gin

# Team
-   Thierno Sadou Barry

## Overview

This is a simple to do list management application built with Go and Gin. It allows you to create, read, update and delete to do list items.

## Features

- Create a new to do list item
- Read all to do list items
- Update a to do list item
- Delete a to do list item

## Usage

```bash
go run cmd/api/main.go
```

## API Endpoints



# 📝 TODO-API - Tests Unitaires ✅

Ce projet implémente une API de gestion de tâches en Go avec Gin. Ce guide décrit comment exécuter les tests unitaires.

## 📌 1. Prérequis
Avant de lancer les tests, assure-toi d'avoir :
- Go installé (version 1.18+ recommandée)
- Les dépendances installées :

  ```sh
  go mod tidy
  ```

## 🚀 2. Exécuter les tests
Pour exécuter tous les tests unitaires du projet, utilise la commande :

```sh
go test ./...
```

Si tu veux exécuter les tests d'un package spécifique :

```sh
go test ./pkg/handlers
go test ./pkg/server
```

## 🧪 3. Tests réalisés

| Test | Description |
|------|------------|
| **TestSetupRouter** | Vérifie que le routeur est bien initialisé. |
| **TestGetTasksHandler** | Vérifie que `GET /tasks` renvoie bien une liste de tâches. |
| **TestCreateTaskHandler** | Vérifie qu'on peut créer une tâche via `POST /tasks`. |
| **TestCreateTaskHandlerInvalidJSON** | Vérifie que l'API gère les JSON invalides avec `400 Bad Request`. |

## ⏳ 4. Test à compléter
- [ ] **TestCreateTaskAndCheckList** : Vérifier qu'une tâche ajoutée est bien récupérable via `GET /tasks`. *(À implémenter !)*

## 🛠 5. Débogage
Si un test échoue, ajoute l'option `-v` pour voir les détails :

```sh
go test -v ./...
```

Pour relancer un test spécifique :

```sh
go test -run TestCreateTaskHandler ./pkg/handlers
```

## 🎯 6. Prochaine amélioration
- Ajouter les tests pour **la mise à jour et la suppression des tâches** (`PUT /tasks/:id` et `DELETE /tasks/:id`).

---

