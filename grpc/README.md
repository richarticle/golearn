# gRPC
This is an example of using gRPC with protocol buffer v3.

**Installation**
```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
wget https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip
unzip protoc-3.0.0-linux-x86_64.zip
cp bin/protoc /go/bin/
go get -u google.golang.org/grpc
```

**Compile protocol.pb.go**
```
protoc --go_out=. *.proto
```

**Working Directory**
```
grpc
├── protocol
│   ├── protocol.pb.go
│   └── protocol.proto
├── client
│   ├── client
│   └── main.go
└── server
    ├── main.go
    └── server
```

**Run**

Server

```
./server
2016/08/03 07:29:59 Numbers: [1 3 8 3]
2016/08/03 07:29:59 Numbers: [8 1 4 2]
2016/08/03 07:30:04 Numbers: [5 1 5 1]
2016/08/03 07:30:04 Numbers: [5 4 9 4]
2016/08/03 07:30:09 Numbers: [3 0 0 2]
2016/08/03 07:30:09 Numbers: [5 8 7 1]
2016/08/03 07:30:14 Numbers: [9 3 7 1]
2016/08/03 07:30:14 Numbers: [9 8 2 1]
2016/08/03 07:30:19 Numbers: [7 2 6 2]
2016/08/03 07:30:19 Numbers: [6 0 9 5]
```

Client

```
./client
2016/08/03 07:29:59 Try to compute sum of [8 1 4 2]
2016/08/03 07:29:59 Try to compute sum of [1 3 8 3]
2016/08/03 07:30:00 Sum of [8 1 4 2] is 15
2016/08/03 07:30:00 Sum of [1 3 8 3] is 15
2016/08/03 07:30:04 Try to compute sum of [5 1 5 1]
2016/08/03 07:30:04 Try to compute sum of [5 4 9 4]
2016/08/03 07:30:05 Sum of [5 1 5 1] is 12
2016/08/03 07:30:05 Sum of [5 4 9 4] is 22
2016/08/03 07:30:09 Try to compute sum of [5 8 7 1]
2016/08/03 07:30:09 Try to compute sum of [3 0 0 2]
2016/08/03 07:30:10 Sum of [5 8 7 1] is 21
2016/08/03 07:30:10 Sum of [3 0 0 2] is 5
2016/08/03 07:30:14 Try to compute sum of [9 8 2 1]
2016/08/03 07:30:14 Try to compute sum of [9 3 7 1]
2016/08/03 07:30:15 Sum of [9 8 2 1] is 20
2016/08/03 07:30:15 Sum of [9 3 7 1] is 20
2016/08/03 07:30:19 Try to compute sum of [6 0 9 5]
2016/08/03 07:30:19 Try to compute sum of [7 2 6 2]
2016/08/03 07:30:20 Sum of [6 0 9 5] is 20
2016/08/03 07:30:20 Sum of [7 2 6 2] is 17
```
