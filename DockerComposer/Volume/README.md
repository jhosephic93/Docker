# VOLUMEN DE TIPO BIND MOUNT:
- Volume de tipo Bind Mount, Podemos montar tanto directorios como ficheros. De esta manera conseguimos:
  - Compartir ficheros entre el host y los containers.
  - Que otras aplicaciones que no sean docker tengan acceso a esos ficheros, ya sean código, ficheros etc…
- Volúmenes docker
  - Persistencia de datos.

1. Crear carpeta.
```console
$ mkdir html
$ echo "Hello from docker compose" > html/index.html
```
2. Editar el docker-compose.yml como sigue.
```console
$ nano docker-compose.yml
```
```yaml
version: '2'
services:
  web:
    image: nginx:alpine
    volumes:
    - "./html/:/usr/share/nginx/html/"
    - "temporal:/mnt/temporal/"
    ports:
    - 8080:80

  alpine:
    image: alpine:latest
    command: ["/bin/sh", "-c", "sleep 3600"]
    volumes:
    - "temporal:/mnt/temporal/"

volumes:
  temporal:
```
3. Levantar.

```console
$ docker-compose up -d
$ curl -I localhost:8080
```
3. Verificar que el VOLUME este sincronizado.
```console
$ docker-compose exec web sh
# echo "xD" > /mnt/temporal/hi.txt
# exit
$ docker-compose exec alpine sh
# cat /mnt/temporal/hi.txt
# exit
```

4. Ispeccionar el nuevo Volume.
```console
$ docker volume inspect clase2_temporal
```
