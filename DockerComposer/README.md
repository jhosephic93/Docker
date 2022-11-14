# DOCKER COMPOSE

## Install Docker Compose
```bash
$ sudo apt install python3 python3-pip -y
$ sudo pip3 install -U docker-compose
$ docker-compose --version
```

1. Crear archivo docker-compose.yml con contenido:
```bash
$ nano docker-compose.yml
```
```yaml
version: '2'
services:
 web:
   image: nginx:alpine
   ports:
     - 8080:80

 alpine:
   image: alpine
   command: ["/bin/sh", "-c", "sleep 3600"]

```
2. Levantar
```bash
$ docker-compose up -d
#Dato: Si modificamos o eliminamos el archivo docker-compose.yml al hacer "docker-compose ps" ya no aparecera, pero si en "docker ps"
```

3. Verificar
```bash
$ docker-compose ps
$ docker-compose logs -f web
$ curl -I localhost:8080
## Es Similiar a -> $ docker run -p 8080:80 -d nginx:alpine
$ docker-compose exec alpine sh
# ping web
```

4. Eliminar
```bash
$ docker-compose down
```