# CREANDO NUESTRO PROPIO NEGINX - DOCKERFILE

1. Crear carpeta y entrar.
```bash
$ mkdir mynginx && cd mynginx
$ nano default.conf
```
2. Crear archivo default.conf con contenido:
```nginx
server {
 listen 80 default_server;
 server_name _;
 root /apps;
 location / {
   try_files $uri $uri/ =404;
 }
 location = /404.html {
   internal;
 }
}
```
3. Crear archido DOCKERFILE
```bash
$ nano Dockerfile
```
```Dockerfile
FROM nginx:alpine
ENV APP_DIR=/apps
RUN apk update && apk add nginx && mkdir -p /run/nginx/
COPY default.conf /etc/nginx/conf.d/default.conf
VOLUME $APP_DIR
WORKDIR $APP_DIR
EXPOSE 80
CMD nginx -g 'daemon off;'
```
4. Scripts
```bash
## Crear archivo index.html:
$ echo "Hellou" > index.html
## Generar el docker image:
$ docker build -t mynginx:latest ./
## Volver a generar el docker image:
$ docker build -t mynginx:0.1 ./
## Generar docker image sin cache:
$ docker build --no-cache -t mynginx:0.2 ./
## No se hizo la prueba de Docker run, solo es un ejemplo de recetario de Dockerfile.
```
5. Excluyendo archivos para Docker Build Context: | .dockerignore
```bash
$ nano .dockerignore
soypesado.txt
vendor/
```