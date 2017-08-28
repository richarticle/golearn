# gobindata
This is an example showing how to use go-bindata package (github.com/elazarl/go-bindata-assetfsi).

## Build
Install go-bindata executable file.
```
go get github.com/jteeuwen/go-bindata/...
```

Generate go-bindata Golang files automatically.
```
go generate
```

## Test
Start web server.
```
./gobindata
```

Browse the following links.
* http://127.0.0.1:8080/normal/index.html
* http://127.0.0.1:8080/bind/index.html
* http://127.0.0.1:8080/bind/deep/index.html
* http://127.0.0.1:8080/template



