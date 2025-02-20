# API REST en Go - Gestion de Tâches (TODO)

## Objectif

Créer une API REST simple en Go qui gère une liste de tâches (TODO). Cette API devra être extensible pour les prochaines étapes du projet.

## Instructions

### 1. Initialiser un projet Go

- Créez un dossier pour le projet.
- Initialisez un module Go avec la commande suivante :

  ```sh
  go mod init votre_nom_de_module
  ```

### 2. Installer les dépendances

- Utilisez le package `github.com/gin-gonic/gin` pour créer l’API.
- Installez-le avec la commande :

  ```sh
  go get github.com/gin-gonic/gin
  ```

### 3. Développer un premier endpoint

- Implémentez un serveur HTTP avec un seul endpoint :

  - `GET /tasks` : Retourne une liste vide de tâches sous format JSON.

### 4. Tester l’API

- Lancez le serveur et vérifiez que `GET /tasks` fonctionne correctement.

## Exécution

Pour exécuter l’API, utilisez la commande suivante :

```sh
go run main.go
```

## Vérification

Une fois le serveur lancé, utilisez un navigateur ou un outil comme `curl` ou Postman pour tester l'endpoint :

```sh
curl http://localhost:8080/tasks
```


**Auteurs :** [Souleymane SALL, ]  
**Date :** [20 Février 2025 ]

