# DOCKER SWARM | DESPLEGAR STACK APP VOTING

1. Iniciar cluster.
```console
$ git clone https://github.com/docker/example-voting-app
$ cd example-voting-app/
$ docker stack deploy --compose-file docker-stack.yml vote   
$ docker stack ls
$ docker stack services vote
```

![Architecture diagram](img/1.jpg)

- Revisar en el visualizador http://ip-publica-nodo1:8080/
- Votar por gatitos o perritos eh http://ip-publica-nodo1:5000/
- Ver resultados eh http://ip-publica-nodo1:5001/
