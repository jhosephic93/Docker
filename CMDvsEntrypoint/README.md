# CMD vs Entrypoint:
## Examples - 01
```bash
$ docker run composer:latest version
# /docker-entrypoint.sh: exec: line 24: version: not found
$ docker run composer:latest --version
# Aqui debe mostrar la version del binario composer
```
## Examples - 02
```bash
$ docker run nginx:alpine version
# /docker-entrypoint.sh: exec: line 47: version: not found
$ docker run nginx:alpine id
# Aqui debe mostrar id.
```
