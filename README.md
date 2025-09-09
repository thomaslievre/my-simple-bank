# my-simple-bank

my-simple-bank est une application web de gestion bancaire simple, développée avec React pour le frontend et Go pour le backend.

## Fonctionnalités

- Création et gestion de comptes bancaires
- Consultation du solde
- Historique des transactions
- Dépôts et retraits
- Authentification des utilisateurs

## Stack technique

- **Frontend** : React, TypeScript
- **Backend** : Go (Golang), REST API
- **Base de données** : PostgreSQL
- **Gestion des états** : (ex: Redux, Context API)
- **Tests** : Go test
- **Conteneurisation** : Docker

## Installation

### Prérequis

- Node.js (>= 18)
- Go (>= 1.20)
- (Optionnel : Docker)

## Commandes Makefile

Voici les principales commandes disponibles pour gérer le projet :

| Commande            | Description                                           |
| ------------------- | ----------------------------------------------------- |
| `make build`        | Compile l'application Go et génère le binaire `main`  |
| `make run`          | Lance l'application Go                                |
| `make test`         | Exécute les tests unitaires Go avec couverture        |
| `make clean`        | Supprime le binaire compilé                           |
| `make migrateup`    | Applique toutes les migrations de base de données     |
| `make migrateup1`   | Applique la dernière migration appliquée              |
| `make migratedown`  | Annule toutes les migrations de base de données       |
| `make migratedown1` | Annule la dernière migration appliquée                |
| `make sqlc`         | Génère le code Go à partir des requêtes SQL avec sqlc |
| `make mock`         | Génère les mocks pour les tests                       |
| `make server`       | Lance le serveur Go (équivalent à `make run`)         |

Exemple pour lancer le projet :

```sh
make run
```
