# DOCKER NETWORK EN DOCKER COMPOSE:

1. Crear carpeta.
```console
$ mkdir html
$ echo "Hello from docker compose" > html/index.html
```
2. Crear docker-compose.yml como sigue:

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
    - "8080:80"
    networks:
    - front

  nginx:
    image: nginx:alpine
    ports:
    - "8081:80"
    networks:
    - front
    - back

  alpine:
    image: alpine:latest
    command: ["/bin/sh", "-c", "sleep 3600"]
    volumes:
    - "temporal:/mnt/temporal/"
    networks:
    - back

volumes:
  temporal:

networks:
  front:
  back:
```

3. Probar si alpine puede ver a nginx y no a web:
```console
$ docker-compose exec alpine sh
# ping nginx # Alpine no puede pinear a web pero si a web
# ping web
```

4. Probar que nginx si pueda hacer ping a web y alpine: 
```console
$ docker-compose exec nginx sh
# ping alpine
# ping web
# ifconfig
```
5. Debemos ver que tiene dos interfaces de red (eth0 y eth1)