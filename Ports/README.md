# Forward de puertos:

1. Container usa puerto 80 y sale por el https://ip-público:8080
```bash
$ docker run --name mynginx -p 8080:80 -d nginx:alpine
$ docker logs -f mynginx
```
En otro terminal
```bash
$ curl -I localhost:8080
```
2. Container usa puerto 80 y sale por el https://ip-público:puerto-aleatorio
```bash
$ docker run -p 80 -d nginx:alpine
```
3. Listar todos los **PORTS** usados por el container.
```bash
$ docker port mynginx
```