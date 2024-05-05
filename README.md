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

Pour installer le projet, il faut executer quelque commandes.
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

### Configuration du domaine _spoofé_

Pour faire croire à notre ordinateur que le domaine reiffeisen.ch nous appartient,
il faut éditer le fichier host (`/etc/hosts` sur UNIX, 
`C:\Windows\system32\drivers\etc\hosts` sur Windows), et y ajouter la ligne 
suivante:

```hosts
127.0.0.1 reiffeisen.ch
```
