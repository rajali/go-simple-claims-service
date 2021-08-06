# GO-SIMPLE-CLAIMS-SERVICE

## Description
- A simple rest api with two GET methods.

# Run locally
```shell
go run .
```

# Update Swagger spec
```shell
swagger generate spec -o ./swagger.json
```

# Run containerised
```shell
docker build -t simple-claims-service:latest .
docker run --rm --name claims-service -p 8080:8080 simple-claims-service:latest
```