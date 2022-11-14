# DOCKER SWARM | LIMITAR RECURSOS CPU & MEMORY.
1. Clonar repo
```console
$ git clone https://github.com/docker/example-voting-app
```
2. Modificar docker-stack.yml

```consolee
$ nano docker-stack.yml
```
3. Dejar el servicio redis como sigue:
```yaml
redis:
    image: redis:alpine
    networks:
      - frontend
    deploy:
      replicas: 1
      update_config:
        parallelism: 2
        delay: 10s
      restart_policy:
        condition: on-failure
      resources:
        reservations:
          cpus: '0.25'
          memory: '256M'
        limits:
          cpus: '0.50'
          memory: '512M'
```

4. Correr Script
```console
$ docker stack deploy --compose-file docker-stack.yml vote   
$ docker service ls
$ docker services inspect vote_redis
```