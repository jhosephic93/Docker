# DOCKER REGISTRY EN LOCALHOST

1. Se va a usar este docker image https://hub.docker.com/_/registry
```docker
$ docker run -v $PWD/data:/var/lib/registry -d -p 5000:5000 --name registry registry:2
$ curl http://localhost:5000/v2/_catalog
```
2. Tagear un docker image para localhost:5000
```docker
$ docker tag alpine:latest localhost:5000/alpine:latest
$ docker images
```
3. Enviar docker image al registry:
```docker
$ docker push localhost:5000/alpine:latest
```
4. Probando si ya tiene el nuevo docker image:
```bash
$ curl http://localhost:5000/v2/_catalog
```
5. Listando contenido del volume local:
```bash
$ ls -la data/docker/registry/v2/repositories/alpine/
```
6. Eliminar y descargar nuevamente:
```bash
$ docker rmi localhost:5000/alpine:latest
$ docker images localhost:5000/alpine:latest
$ docker pull localhost:5000/alpine:latest
$ docker images localhost:5000/alpine:latest
```
7. Listar tags del registry local:
```docker
$ curl http://localhost:5000/v2/alpine/tags/list
$ docker tag alpine localhost:5000/alpine
$ docker push localhost:5000/alpine
$ curl http://localhost:5000/v2/alpine/tags/list
```

# DOCKER REGISTRY SIN HTTPS

```docker
$ docker tag alpine:latest 192.168.161.179:5000/alpine:0.1
```
1. Crear el archivo /etc/docker/daemon.json con contenido:
```bash
$ sudo nano /etc/docker/daemon.json
```
```json
{
  "insecure-registries" : ["192.168.161.179:5000"]
}
```
```bash
$ sudo service docker restart
$ docker push 192.168.161.179:5000/alpine:0.1
```
2. Verificar que se subio la imagen.
```bash
$ curl http://192.168.161.179:5000/v2/alpine/tags/list
```
3. Resolver un dominio sin DNS Server.
```bash
## Agregar la ip y dominio en /etc/hosts
192.168.161.179 docker.colega.com
## Agregamos en /etc/docker/daemon.json el contenido:
{
  "insecure-registries" : ["192.168.161.179:5000", "docker.colega.com:5000"]
}
```
4. Reiniciamos el docker daemon.
```bash
$ sudo service docker restart
```
5. Tageamos y probamos:
```docker
$ docker tag alpine:latest docker.colega.com:5000/alpine:colega
$ docker push docker.colega.com:5000/alpine:colega
$ curl http://docker.colega.com:5000/v2/alpine/tags/list
```