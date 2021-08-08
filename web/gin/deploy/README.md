
## grpc
```sh
protoc -I ${protoDir}/ ${protoDir}/*proto --go_out=plugins=grpc:${outDir}
```
example:

Golang

```sh
protoc -I rpc/grpc/protos/movie/ rpc/grpc/protos/movie/*proto --go_out=plugins=grpc:rpc/grpc/protos/movie
```

Java

```sh
protoc --plugin=protoc-gen-grpc-java --grpc-java_out="$OUTPUT_FILE" --proto_path="$DIR_OF_PROTO_FILE" "$PROTO_FILE"

mvn protobuf:compile-custom

mvn spring-boot:run

```


# build
```sh
_deploy/build.sh
```

# run
```sh
docker-compose -f projects/docker/docker-compose.yml up
```

# build init data job
```sh
docker build -t roandocker/initdata-job:1.0.0 -f projects/docker/initdata-job/Dockerfile  projects/docker/initdata-job
```

helm install apps
```sh
helm dep update k8s/helm/apps/ticket-app/
helm upgrade --install ticket-app k8s/helm/apps/ticket-app/
```

istio injecton
```sh
kubectl get ns  -L istio-injection
kubectl label ns default istio-injection=disabled
```