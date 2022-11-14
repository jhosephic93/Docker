# EJEMPLO ENTRYPOINT

```bash
$ mkdir myentrypoint && cd myentrypoint/
```
1. Crear archivo entrypoint.sh con contenido:
```console
$ nano entrypoint.sh
```
```bash
#!/bin/sh
echo "my entrypoint"
exec $@ # ->> ejecuta el comando que yo le pase como parámetro
```
2. Crear archivo cmd.sh
```console
$ nano cmd.sh
```
```bash
#!/bin/sh
echo "primer comando" && echo "segundo comando"
```
3. Seteamos permisos de ejecucion:
```console
$ chmod +x *.sh
```
4. Crear archivo Dockerfile con contenido:
```console
$ nano Dockerfile
```
```Dockerfile
FROM alpine:latest
COPY entrypoint.sh /docker-entrypoint.sh
COPY cmd.sh /cmd.sh
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["/cmd.sh"]
```
5. Scripts
```console
$ docker build -t myimage ./
```
6. Ejecutar y Analizar salida:
```console
$ docker run --rm myimage id
$ docker run --rm myimage
```
7. Sobre escribir el entrypoint:
```console
$ docker run --rm -ti --entrypoint=id myimage
$ docker history myimage
$ docker run --rm -ti --entrypoint=/bin/sh myimage
# ps aux
# exit
```