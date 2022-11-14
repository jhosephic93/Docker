# DOCKER SWARM | SECRETS

```console
$ printf "This is a secreto" | docker secret create my_secret_data -
$ docker secret ls
```

## Usando el secret en un service:
```console
$ docker service create --name redis --secret my_secret_data redis:alpine
$ docker ps
$ docker exec -ti <redis-containerid> sh
# cat /run/secrets/my_secret_data
# exit
```

## Eliminar Secret:
```console
$ docker secret rm my_secret_data
```