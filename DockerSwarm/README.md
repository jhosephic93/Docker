# DOCKER SWARM

1. Iniciar cluster.
```console
$ docker swarm init
```

2. Listar nodos/workers a partir del Master/Manager/Leader

```console
$ docker node ls
```
3. Correr script en un segundo Server para unirse a un cluster como Worker:
```console
$ docker swarm join --token 
SWMTKN-1-5og2ckcrfsmseoc9mgjwbwl443twtzs1tjppmkn828tjzvrg5v-5wwsul2ayb6xcpvqperxzha3u 192.168.161.171:2377
```
4. Retirarse de un cluster:
```console
$ docker swarm leave --force
```
## Tolerancia a fallos en Docker Swarm:
- https://docs.docker.com/engine/swarm/admin_guide/

## Algoritmo Raft:
- https://docs.docker.com/engine/swarm/raft/ 


# En el Master Leader para visualizar:
## Crear servicio de VISUALIZER
```console
$ docker service create --name=viz --publish=80:8080/tcp --constraint=node.role==manager --mount=type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock dockersamples/visualizer
## unsupported platform on 1 node | No soportado en arm64
```
## Listar servicios a partir del Master
```console
$ docker service ls
```
1. En el browser http://51.68.161.171/ 

## Crear nuevo Servicio:
```console
$ docker service create --name mystack_serviceweb --publish 8081:80 --replicas=5 nginx:alpine
## No funciona en arm64
```
2. Probar en browser: http://51.68.161.171:8081/

### En cada Worker:
```console
$ docker ps | grep nginx
$ docker logs -f <containter-id>
```

### Scaling del Service:
```console
$ docker service scale <service-id>=15 # o tambien
$ docker service scale mystack_serviceweb=15
```

### Desplegar el Stack del App docker-stack.yaml:
```console
$ docker stack deploy --compose-file docker-stack.yml vote   
```
### Listar los Stack de Docker Swarm
```console 
$ docker stack ls
$ docker stack services vote
```
### Eliminar un Stack de Docker Swarm
```console 
$ docker stack rm vote
```
### Revisar en el visualizador http://51.68.161.171/ 
