# Docker Networking

1. Crear Network
```bash
$ docker network create -d bridge demonet
## Opcional, setear una subnet cambiando clase:
$ docker network create -d bridge --subnet=172.20.0.0/24 subnet1
```
2. Listar Network
```bash
$ docker network ls
```
3. Informacion de Network
```bash
$ docker network inspect demonet
```
4. Eliminar Network
```bash
$ docker network rm demonet
```
## Verificar la RED DE CONTAINER

```bash
$ docker run -ti alpine sh
# ip addr
Revisar que tenga dos redes lo y eth0@if30
```


### Examples
1. Crear container con NETWORK:
```bash
$ docker run -ti --network demonet alpine sh
# ifconfig
## (Revisar si esta en inet addr:172.18.0.x)
# exit
```
2. Quitar un Network de un Container:

```bash
$ docker run -ti --network demonet --name mysh alpine sh
## En un segundo terminal
$ docker network disconnect demonet mysh
## Regresar al terminal del container y verificar que no tenga red:
# ifconfig
## En el segundo terminal volver a conectar:
$ docker network connect demonet mysh
## Regresar al terminal del container y verificar que tenga red eth1:
```

2. Networking: Saber IP-Public desde Container
```bash
$ docker network ls 
$ docker run -ti alpine sh
# apk add curl
# curl ifconfig.me
```