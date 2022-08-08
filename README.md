# Core Protocol Buffer

Protocol Buffer (protobuf) para firmadores.

## Instrucciones (Go)

Este proceso indica como generar los archivos respectivos en Go:

1. Instalar [protobuf](https://github.com/protocolbuffers/protobuf).

2. Instalar los plugins:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

3. Definir la variable de ambiente donde se instalaron los plugins, ejemplo:

```
export PATH=$PATH:~/go/bin
```

4. Generar los archivos:

```
protoc --go_out=. --go_opt=paths=source_relative \
 --go-grpc_out=. --go-grpc_opt=paths=source_relative \
 docencia.proto
```

4. Generar los archivos JS:

```
protoc-gen-grpc --js_out=import_style=commonjs,binary:./js --grpc_out=./js --proto_path . docencia.proto
```

## Referencias

- [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
- [How to install Protobuf on Ubuntu](https://linuxhint.com/install-protobuf-ubuntu/)

