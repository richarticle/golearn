# Protocol Buffer
This is an example of using protocol buffer v3.

**Installation**
```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
wget https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip
unzip protoc-3.0.0-linux-x86_64.zip
cp bin/protoc /go/bin/
```

**Compile protocol.pb.go**
```
protoc --go_out=. *.proto
```

**Run**
```
go build
./protobuf
protobuf len = 24
[10 5 72 101 108 108 111 16 17 26 3 1 2 3 34 8 10 6 79 108 105 118 101 114]

newTest: Hello 17 [1 2 3] Oliver

json len = 66
{"label":"Hello","type":17,"reps":[1,2,3],"man":{"name":"Oliver"}}

```
