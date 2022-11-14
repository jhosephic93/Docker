# EJERCIOS

## DEMO-01
RETO: Agregar mynewweb.com para resolver '**localmente**':

```bash
$ sudo nano /etc/hosts
# 127.0.0.1 mynewweb.com
```
1. Conseguir archivo de configuracion:
```bash
$ docker run --name mycontainer -d nginx:alpine
$ docker cp mycontainer:/etc/nginx/conf.d/default.conf ./
```
2. Editar archivo de local '**default.conf**' reemplazando localhost por mynewweb.com y eliminar container mycontainer.
2. Finalmente montar mediante volume la nueva configuracion y probar:
```bash
$ docker run -p 80:80 -d -v $PWD/default.conf:/etc/nginx/conf.d/default.conf nginx:alpine
$ curl -I mynewweb.com
```
4. En browser cargar  https://ip-publica

### CON DOCKERFILE   ******************************************

RETO: Agregar mynewweb.com para resolver '**localmente**':

```bash
$ sudo nano /etc/hosts
# 127.0.0.1 mynewweb.com
```
1. Conseguir archivo de configuracion:
```bash
$ docker run --name mycontainer -d nginx:alpine
$ docker cp mycontainer:/etc/nginx/conf.d/default.conf ./
```
2. Crear carpeta y acceder.

```bash
$ mkdir mynginx/ && cd mynginx/
$ cp ../default.conf ./
```
3. Crear archivo '**Dockerfile**' con contenido:
```Dockerfile
FROM nginx:alpine
COPY default.conf /etc/nginx/conf.d/default.conf
```
4. Generar docker image:
```bash
$ docker build -t mynginx:v1 ./ # Opcion solo local
$ docker build -t <usuario-docker-hub>/nginx:v1 ./ # Opcion para poder subir al dockerhub
```
5. Ejecutar nueva image:
```bash
$ docker run -d -p 80:80 <usuario-docker-hub>/nginx:v1
$ curl -I mynewweb.com
```