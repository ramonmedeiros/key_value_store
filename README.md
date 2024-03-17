# Key-value store

## Before run

Expectations:

1. Have golang installed
2. Have makefile installed
3. Have internet to download modules

## How to run

```
$ GIN_MODE=release go run cmd/main.go
```
OR

```
make run-prod
```

You will see the rest frawework log

## How to interact

`GET /:key`

```
$ curl http://localhost:8080/key
```

`POST /:key` value

```
$ curl http://localhost:8080/key -d 'value'
```

## Tests

```
make tests
```

## Design considerations

The proposed solution is for a standalone in in-memory cache. When scaling, only vertical scaling is supported. For a horizontal scaling, the project will need a massive change to support partitions and to communicate between nodes

### Original request
```
Backend Work Test

Create a simple in-memory cache with an HTTP interface

Interface
* HTTP `POST /<key>` with the value as UTF-8 body
* HTTP `GET /<key>` replies with the value as body or 404 if no such key exists

Keep in mind
- Use idiomatic Go
- Stored key expires after 30 minutes
- The server is expected to handle a large load with an 80/20 ratio between reads and
writes
- Write it as if you would be expected to continue to support it in production for the
foreseeable future
- Add a README file with instructions on how to build, use and any additional information you might have

Remember, we want to see how you would build an in-memory cache with the given specifications. With this said; you're free to use external dependencies but for the core application (i.e. the cache) we want your implementation 

Important! Please don't publish the test assignment or your result on the internet
```