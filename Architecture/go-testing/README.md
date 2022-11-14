# Multi stage:
1. Clonar:
```bash
$ git clone https://github.com/mario21ic/go-testing.git
```
2. Crear archivo Dockerfile
```Dockerfile
FROM golang:1.13 as builder
WORKDIR /app
COPY ./ /app/
RUN go build main.go
#CMD ["/app/main"]

FROM scratch
COPY --from=builder /app/main /main
CMD ["/main"]
```
```console
$ docker build -t goruntime ./
$ docker images goruntime
```
3. Debemos ver que la imagen pesa apenas 2mb
```console
$ docker run --rm goruntime
```
4. Revisar el historial del docker image:
```console
$ docker history goruntime
```

## Nota:
#### **Compilacion estatica**: todas sus dependencias son incluidas en el compilado. Por ello el compilado es de mayor peso.
#### **Compilacion dinamica**: todas las dependencias son referenciadas. Por eso el compilado es de menor peso.

# Compilar Docker image para otra arquitectura (Arm64)
```console 
$ docker buildx create --name myarm64
$ docker buildx use myarm64
$ docker buildx ls
$ docker buildx inspect --bootstrap
$ docker buildx build --platform linux/arm64,linux/amd64 -t <usuario-docker-hub>/goruntime ./ --push
```
- Se debe subir el docker image a su docker hub y mostrar ambas arquitecturas
