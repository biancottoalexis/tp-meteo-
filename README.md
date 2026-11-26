# TP Météo Go

Ce projet consiste à créer un site météo en Go. Il permet de rechercher une commune française et de consulter ses prévisions météo pour les 7 prochains jours.

J'ai utilisé deux API : l'API Géo pour rechercher les communes et récupérer leurs coordonnées, puis Open-Meteo pour obtenir les prévisions météo.

## Fonctionnalités

- Rechercher une commune avec son nom
- Choisir une commune parmi les résultats
- Consulter la météo des 7 prochains jours
- Afficher les dates en français, les températures et les icônes météo
- Voir la commune sur une carte OpenStreetMap
- Comparer les températures avec des barres
- Afficher un message d'erreur si une recherche ou une requête ne fonctionne pas

## Technologies utilisées

Le projet utilise Go pour le serveur et les appels aux API, HTML pour les pages et CSS pour la mise en forme.

Les pages sont générées avec les templates Go.

## Installation et lancement

Il faut avoir Go installé sur son ordinateur.

Ouvrir le projet dans VS Code, puis ouvrir un terminal à la racine du projet.

Lancer le serveur avec :

```bash
go run .
```

Ouvrir ensuite l'adresse locale correspondant au port configuré dans `main.go`.

Pour arrêter le serveur, faire `Ctrl + C` dans le terminal.

## Organisation du projet

- `api` : appels aux API Géo et Open-Meteo
- `models` : structures utilisées pour les données
- `services` : traitement des prévisions et calcul des barres de températures
- `controllers` : gestion des requêtes et affichage des pages
- `routes` : déclaration des routes du serveur
- `helpers` : fonctions pour les dates, températures, icônes et la carte
- `templates` : pages HTML
- `assets` : fichiers statiques comme le CSS
- `main.go` : démarrage du serveur

## Routes

| Route | Description |
|---|---|
| `/` | Page d'accueil |
| `/ville?nom=Marseille` | Recherche d'une commune |
| `/meteo?code=13055` | Prévisions météo de la commune sélectionnée |

## Gestion des erreurs

Le projet prévoit plusieurs statuts HTTP :

- 400 : recherche invalide ou code manquant
- 404 : commune ou page introuvable
- 500 : erreur interne du serveur
- 503 : API ou prévisions indisponibles

Les requêtes vers les API utilisent un timeout de 5 secondes et un User-Agent personnalisé.

## Tests

J'ai vérifié la compilation du projet, le lancement du serveur, la recherche d'une commune et l'affichage de la météo.

J'ai aussi testé les erreurs 400 et 404 avec Postman.

Les tests des erreurs 500 et 503 restent