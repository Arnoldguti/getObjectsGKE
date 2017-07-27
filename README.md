# getObjectsGKE
API - Obtener objetos de Google Container Engine

API en Go que puede implementarse dentro del cluster de contenedores de Google Container Engine, para poder obtener información sobre servicios y pods (Dirección IP, Estado de las implementaciones, Número de contenedores por Pod y su información descriptiva)

Dockerfile para la creación de la imágen Docker:

Instrucciones:
1. Localizarse en el directorio de la aplicación
2. Empaquetar las dependencias (Godep)
3. Ejecutar para crear la imágen Docker: docker build -t appName .
4. docker run --publish 6060:8080 --appName test --rm appName
