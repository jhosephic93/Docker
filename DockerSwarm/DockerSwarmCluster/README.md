# DOCKER SWARM CLUSTER

## Pre-requisitos
- Tener instalado Docker en los nodos (Master/Worker)

## Configurar Master/Worker
- NODO1 = Master
```console
$ docker swarm init --advertise-addr=<ip-privada-nodo1>
## Debe mostrar el comando para ingresar al cluster como Worker:
```

- NODO2 = Worker
```bash
$ docker swarm join --token <token-swarm> <ip-privada-nodo1>:2377
```

- NODO1 = Master
1. Listar todos los nodos/workers a partir del Master
```console
$ docker node ls
```

## CREAR SERVICIO-01

1. Crear un servicio que brinde una interfaz grafica en el nodo master=manager para VISUALIZAR los containers en el master y/o nodos
```console
$ docker service create --name=viz --publish=80:8080/tcp --constraint=node.role==manager --mount=type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock dockersamples/visualizer
$ docker service ls
```
- En el browser veremos una interfaza grafica de los containers http://ip-publica-nodo1/ 

3. Crear nuevo Servicio:
```console
$ docker service create --name myserviceweb --publish 8081:80 --replicas=3 nginx:alpine
```
- Probar en browser: http://ip-publica-nodo1:8081/
4. Crear servicio solo en Worker:
```console
$ docker service create --name myserviceweb2 --constraint=node.role==worker --publish 8082:80 --replicas=3 nginx:alpine
```
- Probar en browser: http://ip-publica-nodo1:8082/
- En el browser visualizar containers: http://<ip-publica-nodo1>/ 
- Probar la resiliencia de un docker service:
```bash
$ docker container ls
$ docker rm -f <container id>
$ docker service ls
## Debe aparecer eliminado (2/3) y luego debe recuperarse 3/3 por ejemplo.
```
## Promover un Nodo Worker como Manager:
### En el nodo1:
```console
$ docker node ls
$ docker node promote <node-id>
$ docker node ls
```
### En el nodo2:
```console
$ docker node ls
$ docker service ls
```

## Services Scaling:
```console
$ docker service scale myserviceweb=5
$ docker service scale myserviceweb=0
```
## Eliminar un Service:
```console
$ docker service rm myserviceweb
```
## Inspeccionar un Service:
```console
$ docker service inspect myserviceweb2
```
### En nodo1:
```bash
#Si nos fijamos el service myserviceweb2 esta en 0/3, debido a que no existen nodos de tipo worker, entonces cambiaremos el rol al nodo2 como worker.
```
```console
$ docker node ls
$ docker node demote <node-id>
$ docker service ls
```
### Probar en nodo2:
```console
$ docker node ls
```
### Listar containers de los services:
```console
$ docker service ps myserviceweb2
```
### Eliminar todos los services:
```console
$ docker service rm viz
$ docker service rm myserviceweb2
```