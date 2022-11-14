# Docker

## Installation on Linux

1. On console

```bash
$ sudo su
# id
# echo $USER
# apt update && apt install -y curl
# curl -sSL https://get.docker.com | sh
# exit
$ sudo usermod -aG docker $USER
$ id
$ exit
```

2. Volver a logearse.

```bash
$ id
$ docker ps
```

## Uninstall Docker

```console
$ sudo apt list --installed | grep docker
$ sudo apt remove docker*
$ sudo iptables -L
$ sudo reboot
$ sudo iptables -L #Revisamos que las reglas de IPTables generadas por Docker desapareceran.
```


## Validar Client - Server

```bash
$ docker --version
$ sudo service docker status
$ docker info
$ docker context ls
$ docker context use default
$ ps aux | grep docker
```

## Visualizar Containers Layers

```bash
$ sudo ls -la /var/lib/docker/containers/
```

## Listar el socket de Docker daemon

```bash
$ sudo ls -la /run/containerd/containerd.sock
```

## Ver Kernel de Linux (vemos que comparten el mismo Kernel tanto el container como mi local)

```bash
$ uname -a
$ uname -r
## Validar Kernel de Linux con Kernel de Container
$ docker run --rm ubuntu:18.04 uname -r
```

## Basic commands

### Containers

1.  List containers

```bash
$ docker ps -a
$ docker container ls -a
```

2.  Run containers

```bash
$ docker run -ti ubuntu bash
$ docker run -ti centos:8 sh
$ docker run --name myalpine alpine:latest echo "Hello from container"
```

3.   Remove a container wiith status exited

```bash
$ docker rm <container-id>
# Force the removal of a running container
$ docker rm -f <container-id>
# Eliminar TODOS los containers (exited, running)
$ docker rm -f $(docker ps -aq)
```

4. Purga containers que no estan corriendo.

```bash
$ docker system prune -f
```

5. Ver metadata de Container

```bash
$ docker inspect <name-container>
$ docker container inspect <name-container>
$ docker inspect <name_container> | grep ddr
```

6. Ver logs de Container

```bash
$ docker logs -f <name-container>
$ docker logs <name-container>
```

7. Ver metricas de todos los Containers

```bash
$ docker stats
```

8. Ingresar al Container para run commands

```bash
$ docker exec -ti <name-container> sh
```

9. Ejecutar comandos en los containers sin ingresar

```bash
$ docker exec <container-id> ps
$ docker exec <container-id> apk add htop
$ docker exec -ti <container-id> htop
```

### Images

1.  List images

```bash
$ docker images -a
```

2.  Remove images

```bash
$ docker rmi <image-id>
```

3.  Search Image on Registry

```bash
$ docker search <name-image>:<tag>
$ docker search <name-image> --limit 5
$ docker search --help
```

4.  Pull images

```bash
$ docker pull <name-image>:<tag>
$ docker pull <username-docker-hub>/<name-image>:<tag>

```
5.  Taggear images

```bash
$ docker tag ubuntu:latest <username-docker-hub>/<name-image>:<tag>
```

6.  Revisar history de comands de un DOCKER IMAGE

```bash
$ docker history <image-id>:<tag> --no-trunc
```

## Docker Hub

1. Login DockerHub
```bash
$ docker login
$ docker logout
```

2. Push image a DockerHub

```bash
$ docker push <username_docker_hub>/ubuntu:latest
```

3. Validar usuario logeado en DockerHub

```console
$ cat $HOME/.docker/config.json
```

## ESCANEAR VULNERABILIDADES DE DOCKER IMAGES

```console
$ sudo su
# curl -sSfL https://raw.githubusercontent.com/anchore/grype/main/install.sh | sh -s -- -b /usr/local/bin
# grype version
# grype michaeltinga/myhtop
# grype <name_image>
```

- Escaneo a base de vulenrabilidades en mitre.org
- Clonar y seguir readme.md https://github.com/mario21ic/anchore-demos

## DOCKER API

```console
$ sudo ps aux|grep dockerd
$ sudo apt install jq -y
```

1. Listar containers haciendo el curl es lo mismo que docker ps -a

```console
$ curl -s --unix-socket /var/run/docker.sock http://1.24/containers/json | jq .[].Names # = Es igual que hacer "docker ps"
$ docker container ls
```

2. Listar Images haciendo curl es lo mismo que docker images
```console
$ curl -s --unix-socket /var/run/docker.sock http://v1.24/images/json | jq .[].RepoTags # = Es igual que hacer "docker images"
$ docker images
```

## CREAR TU PROPIO DOCKER BASE IMAGE.

```console
$ sudo apt-get install debootstrap
$ sudo debootstrap xenial xenial > /dev/null
$ sudo tar -C xenial -c . | docker import - xenial
$ docker run xenial cat /etc/lsb-release
```
- Mas info https://docs.docker.com/develop/develop-images/baseimages/
