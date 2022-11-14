# SCALING DOCKER COMPOSE:

1. Crear carpeta.
```console
$ mkdir html
$ echo "Hello from docker compose" > html/index.html
```
2. Editar el docker-compose.yml como sigue.
```console
$ nano docker-compose.yml
```
```yaml
version: '2'
services:
  web:
    image: nginx:alpine
    volumes:
    - "./html/:/usr/share/nginx/html/"
    - "temporal:/mnt/temporal/"
    ports:
    - "8080:80"

  nginx:
    image: nginx:alpine
    ports:
    - "80"

  alpine:
    image: alpine:latest
    command: ["/bin/sh", "-c", "sleep 3600"]
    volumes:
    - "temporal:/mnt/temporal/"

volumes:
  temporal:

```
3. Levantar y escalar.

```console
$ docker-compose up -d
$ docker-compose scale nginx=5
$ docker-compose ps
```
4. Revisar que se escalo el servicio hasta 5 containers cada uno con su propio puerto.