# ARG

## Compilar Docker image de distintas versiones de base image:
1. Crear Dockerfile con contenido:
```console
$ nano Dockerfile
```
```Dockerfile
ARG UBUNTU_VERSION=18.04
FROM ubuntu:${UBUNTU_VERSION}
RUN cat /etc/*release > /version.txt
```
2. Compilar para ubuntu 20.04:
```console
$ docker build --build-arg UBUNTU_VERSION=20.04 -t michaeltinga/ubuntu20.04 ./
$ docker run --rm michaeltinga/ubuntu20.04 cat /version.txt
```