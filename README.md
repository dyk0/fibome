# Welcome to Fibome(Fibø-meh)
Give me Fibonocci! An API to satisfy that Fibonocci need.

## Description
This package returns Fibonocci numbers in sequence of size N

## Build
Go 1.10 has been tested with Fibome
```
$ mkdir -p $GOPATH/src/github.com/dyk0
$ git clone https://github.com/dyk0/fibome.git $GOPATH/src/github.com/dyk0/fibome
$ cd $GOPATH/src/github.com/dyk0/fibome
$ make
```

Run and Test locally
```
$ make run
$ curl http://localhost:8000/5
[0 1 1 2 3]
```

Docker build is supported
```
$ make docker-build
```

## Run
Docker image is available on Dockerhub
```
$ docker pull dyk0/fibome
$ docker run -p 8000:8000 dyk0/fibome
```
