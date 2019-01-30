# Guptime

> Tiny dockerized API to check the availability of a list of URLs

## Getting started

```bash
$ go mod download
$ URLS=www.google.fr:https,nicolas.ritouet.com:https go run main.go
```

## Build

### Using go

```bash
$ go build main.go
```

### Using docker

```bash
docker build . -t guptime
docker run -e URLS=www.google.fr:https,nicolas.ritouet.com:https guptime
```

## Deployment

This docker image is an easy way to make sure your ecs/eks/k8s/swarm cluster has access to the outside world and to given resources.

```bash
docker build . -t nicolasritouet/guptime
docker push nicolasritouet/guptime
```

Once the image has been deployed on docker hub, you can use it wherever you need:

`docker run -e URLS=www.google.fr:https,nicolas.ritouet.com:https nicolasritouet/guptime`