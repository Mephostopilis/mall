# edu

## 简介

## 安装

``` bash
$ go get -u \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
    google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    github.com/go-kratos/kratos/cmd/kratos/v2@latest \
    github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest \
    github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest \
    github.com/go-kratos/kratos/tool/protobuf/protoc-gen-bswagger@latest
```

## 启动

### fileserver

```bash
./fileserver
```

### admin

```bash
./admin -conf ../../configs/admin.yaml
```

### sso

```bash
./sso -conf ../../configs/conf.yaml
```

### sys

```bash
./sso -conf ../../configs/conf.yaml
```

### tiku

```bash
./tiku -conf ../../configs/conf.yaml
```

### cms

```bash
./cms -conf ../../configs/conf.yaml
```

### uuid

```bash
./uuid -conf ../../configs/conf.yaml
```

### gateway

```bash
./gateway -conf ../../configs/conf.yaml
```

### caddy

```bash
./caddy
```
