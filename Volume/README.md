# Docker Volumen
## VOLUME COMPARTIDO
1. Volume compartido de localhost con container
```bash
$ mkdir html && echo "Hello from container" > html/index.html
$ docker run --name mynginx -v $PWD/html/:/usr/share/nginx/html/ -p 8080:80 nginx:alpine
```
En browser cargar  https://ip-publica:8080

## VOLUMES
1. Crear Volume
```bash
$ docker volume create demovol
```
2. Listar Volume
```bash
$ docker volume ls
```
3. Informacion de Volume
```bash
$ docker volume inspect <name-volume>
```
4. Eliminar Volumes
```bash
$ docker volume rm <name-volume>
## Eliminar Volume no usados
$ docker volume prune
```
### Examples
1. En la primera consola
```bash
$ docker run --name mysh1 -ti -v demovol:/tmp/demovol alpine sh
# echo "xD" > /tmp/demovol/hi1.txt
# ls -la /tmp/demovol
# exit
$ docker run --name mysh2 -ti -v demovol:/tmp/demovol alpine sh
# ls -la /tmp/demovol
# cat /tmp/demovol/hi1.txt
```