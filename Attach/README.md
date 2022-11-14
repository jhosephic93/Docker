# Usando ATTACH
Recupera el proceso y lo ve como logs pero si se presion '**ctrl + c**', este proceso se cerrara y por ende el container pasara a **exited**.

```bash
$ docker run -d nginx:alpine
$ docker inspect <container-id>|grep ddr
$ curl -I <ip-container>
$ docker logs -f <container-id>
$ docker attach <container-id>
```