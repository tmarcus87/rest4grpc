Rest4gRpc
=========

Rest4gRPC is a simple REST proxy for gRPC.

## Usage

```bash
tty1 $ r4g -target grpc://127.0.0.1:5000

tty2 $ curl -X GET http://localhost:8888/hoge.fuga.TestService/TestMethod?hoge=fuga

or 

tty2 $ curl -X POST -d '{"hoge":"fuga"}' http://localhost:8888/hoge.fuga.TestService/TestMethod
```

For more information `$ r4g -help`

## Development

**Compile pb file**

```bash
$ cd path/to/checkout
$ docker run --rm \
  -v `pwd`/pb:/go/src/proto \
  -w /go/src/proto \
  namely/protoc-all \
  -f test.proto \
  -l go \
  -o /go/src/proto
```
