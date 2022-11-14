# GUARDAR Y CARGAR DOCKER IMAGES

```bash
$ docker save nginx:alpine -o saveimage.tar
$ docker rmi nginx:v1
$ docker load < saveimage.tar
```