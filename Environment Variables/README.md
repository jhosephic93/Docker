# Environment Variables | Variables de Entorno

1.  Set environment variables
```bash
$ docker run -e MIVAR=foobar -ti alpine sh
```
Ver varibles
```console
# env
# echo $MIVAR
```
2. Otro ejemplo:
```console
$ docker run --rm -e MIVAR=foobar -e VAR2=bar alpine env
```