# Crear DOCKER IMAGE a partir de un Container

1. En un primer terminal:
```bash
$ docker run -ti --name mycontainer alpine sh
# apk add htop
# htop
```
2. En un segundo terminal:
```console
$ docker container commit mycontainer <usuario-docker-hub>/myhtop:latest
$ docker images
$ docker run -ti <usuario-docker-hub>/myhtop htop
```