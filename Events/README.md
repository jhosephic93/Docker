# DOCKER EVENTS
Obtiene los eventos en tiempo real del Server, si hago un docker run aparecera en docker events.
1. En primer terminal.
```bash
$ docker events
```
2. En segundo terminal.
```bash
$ docker run nginx:alpine
## Veremos en el primer terminal los events (image pull, container create, network connect, container start)
```