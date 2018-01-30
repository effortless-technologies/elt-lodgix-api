# Effortless Lodgix Middleware API

## Dependencies

[Go 1.9.1](https://golang.org/doc/install)
[Dep](https://github.com/golang/dep#dep)
[Docker](https://docs.docker.com/engine/installation/)

### API SPEC

[API Spec](./docs/API_SPEC.md)

## Development

`cd` into <gopath>/src/github/effortless-technologies clone repo:

```
$ cd $GOPATH/src/github/effortless-technologies
$ git clone git@github.com:effortless-technologies/el-connect-api.git
```

`cd` into project root and get dependencies:
```
$ cd el-connect-api
$ dep ensure
```

## Build

#### Simple local build

`cd` into project and run:

```
$ cd <root/of/project>
$ go build -tags dev
$ ./elt-delivery
```

Now a server is runnning on localhost:1323:

```
$ curl http://localhost:1323
```

#### Containerized build via Docker

`cd` into project and run:

```
$ cd <root/of/project>
$ go run main.go
```

Now build the Docker image:

```
docker build -t <registry/cluster/image> .
```

Now run the Docker image:

```
docker run -i -t -p 1323:1323 <registry/cluster/image>
```

Now a server is runnning on localhost:1323:

```
$ curl http://localhost:1323
```

## Testing

To run tests locally, use the following command

```
go test ./... -tags dev -v
```