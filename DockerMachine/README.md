# DOCKER MACHINE

## Install VirtualBox
```bash
$ sudo apt install -y virtualbox
```

## Install Docker Machine on Linux
1. On console
```console
$ base=https://github.com/docker/machine/releases/download/v0.16.0
$ curl -L $base/docker-machine-$(uname -s)-$(uname -m) >/tmp/docker-machine
$ chmod +x /tmp/docker-machine
$ sudo mv /tmp/docker-machine /usr/local/bin/docker-machine
$ docker-machine --version
```

## Crear un nuevo Docker Server:
```console
$ docker-machine create -d virtualbox demo
```

## Listar las vm:
```console
$ docker-machine ls
```

## Conectarse al vm:
```console
$ eval "$(docker-machine env demo)"
```

## Eliminar una vm:
```console
$ docker-machine rm -f demo
```

## Probar lanzando un container:
```console
$ docker info
$ docker run -ti alpine sh
# ps aux
$ docker ps -a
```

## Entrando por ssh a la vm:
```console
$ docker-machine ssh demo
$ docker ps -a
$ docker images
$ cat /etc/*release
```

### More info
- https://www.slideshare.net/mario21ic/docker-ecosystem-part-iii-machine 
### Crear Docker Swarm desde Docker Machine (revisar si funciona)
- https://www.adictosaltrabajo.com/2015/12/03/docker-compose-machine-y-swarm/  