# APLICACION DE VOTOS

![Architecture diagram](img/1.jpg)
1. Se va desplegar una aplicacion de votos para perritos o gatitos:
```console
$ git clone https://github.com/docker/example-voting-app
$ cd example-voting/
```
2. Revisar el archivo docker-compose.yml, luego los Dockerfile de cada carpeta.
Levantar project:

```console
$ docker-compose up -d
```
3. Abrir en browser (reemplazar localhost por su IP pública del VPS nodo1):
- Para votar: http://51.68.161.171:5000/  
- Para ver resultados: http://51.68.161.171:5001/  

# Nota: 
- Mala practica tener Makefiles : make clean, make generate-dirs, make permissions, make scripts.
- Buena practica Build docker image (imagen con todo lo necesario) <- Tasks (scripts ejem: db limpia, entrypoint) 
Separar el Runtime (todas las veces ejecutar un app o alguna tasks).
