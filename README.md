Rest4gRpc
=========

Rest4gRPC is a simple REST proxy for gRPC.

## Usage

```bash
tty1 $ r4g -target grpc://127.0.0.1:5000
tty2 $ curl http://localhost:8888/hoge.fuga.TestService/TestMethod
```

For more information `$ r4g -help`
