# FAKE REIFFEISEN LOGIN

This project is used for educational purpose in the context of a conference
about fake news. Please do not use.

## Installation
### Prérequis
- git
- compilateur go
- npm
- node

### Installation et lancement

Pour installer le projet, il executer quelque commandes.
```sh
# Cloner le dépôt
git clone https://github.com/drafolin/fake-login-raif

# Compiler l'api de récupération
cd api
go build

# Lancer l'api 
go run .

```

Ensuite, dans un autre shell ou en lançant l'api en tant que _job_, il faut
lancer le serveur web. Pour ce faire, se rendre dans le dossier racine.

```sh
# Lancer le serveur
cd web
npm start
```
