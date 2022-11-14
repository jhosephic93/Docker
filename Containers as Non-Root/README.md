# CONTAINERS as NON-ROOT
## EJEMPLO DE VULNERACION DE SEGURIDAD
1. Editar un archivo con PERMISOS ROOT desde CONTAINER: PELIGRO
```bash
$ docker run -ti -v /etc/passwd:/etc/passwd alpine sh
# echo "Agregar ultima linea" >> /etc/passwd
# exit
$ tail /etc/passwd
```

## Crear containers como Non-Root
1. Crear un file.txt con el user Ubuuntu mediante touch.
```bash
$ docker run --rm -u $(id -u):$(id -g) -v $PWD/:/tmp/ alpine touch /tmp/file.txt
$ docker run -ti -u $(id -u):$(id -g) -v $PWD/:/tmp/ alpine sh
$ id
$ echo "Its me" > /tmp/hello.txt
$ exit
## Extra:
## Probando a reconocer usuario en el container:
$ docker run --rm -u $(id -u):$(id -g) -v $PWD/:/tmp/ -v /etc/passwd:/etc/passwd alpine id
```

2. Crear un file-root.txt desde container.
```bash
$ docker run --rm -v $PWD/:/tmp/ alpine touch /tmp/file-root.txt
$ ls -la ./*.txt
```

3. Nginx running as Non-Root.
```bash
$ docker run -p 80 -u $(id -u):$(id -g) -ti nginx:alpine
# (Debe salir error por el usuario no puede iniciar nginx)
```
```bash
$ git clone https://github.com/mario21ic/docker-security.git
$ cd docker-security/4-run-nginx
$ cat ./start.sh
$ ./start.sh
## En un segundo terminal ejecutar:
$ docker ps
$ docker inspect <container-id> | grep ddr
$ curl <container-ip>:8080
$ curl localhost:80
```

## EJECUTAR PHP COMPOSER
1. Crear archivo composer.json con contenido:
```bash
$ mkdir composer && cd composer && nano composer.json
```
```php
{"require": {"phpmailer/phpmailer": "^5.2"}}
```
### CONTAINER AS ROOT: Lanzar container con cmd install con root:
```bash
$ docker run --rm -ti -v $PWD:/app composer install
# Vemos que se crea la carpeta vendor siendo el propietario root
$ ls -la vendor/
```

### CONTAINER AS USER: Crear container con el usuario ubuntu que lo ejecuta:
```bash
$ docker run --rm -ti -v $PWD:/app -u $(id -u):$(id -g) composer install
# Vemos que se crea la carpeta vendor siendo el propietario ubuntu
$ ls -la vendor/
```