# RPC
This is an example of using RPC.

**Working Directory**
```
rpc
├── protocol
│   └── protocol.go
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
2016/08/03 15:12:09 Numbers: [0 3 0 4]
2016/08/03 15:12:09 Numbers: [7 2 3 3]
2016/08/03 15:12:14 Numbers: [9 5 2 0]
2016/08/03 15:12:14 Numbers: [8 6 2 8]
2016/08/03 15:12:19 Numbers: [4 9 6 0]
2016/08/03 15:12:19 Numbers: [8 1 7 0]
2016/08/03 15:12:24 Numbers: [2 2 8 7]
2016/08/03 15:12:24 Numbers: [3 5 5 1]
2016/08/03 15:12:29 Numbers: [5 3 7 8]
2016/08/03 15:12:29 Numbers: [1 0 8 2]
```

Client

```
./client
2016/08/03 15:12:09 Try to compute sum of [7 2 3 3]
2016/08/03 15:12:09 Try to compute sum of [0 3 0 4]
2016/08/03 15:12:10 Sum of [7 2 3 3] is 15
2016/08/03 15:12:10 Sum of [0 3 0 4] is 7
2016/08/03 15:12:14 Try to compute sum of [8 6 2 8]
2016/08/03 15:12:14 Try to compute sum of [9 5 2 0]
2016/08/03 15:12:15 Sum of [8 6 2 8] is 24
2016/08/03 15:12:15 Sum of [9 5 2 0] is 16
2016/08/03 15:12:19 Try to compute sum of [8 1 7 0]
2016/08/03 15:12:19 Try to compute sum of [4 9 6 0]
2016/08/03 15:12:20 Sum of [4 9 6 0] is 19
2016/08/03 15:12:20 Sum of [8 1 7 0] is 16
2016/08/03 15:12:24 Try to compute sum of [3 5 5 1]
2016/08/03 15:12:24 Try to compute sum of [2 2 8 7]
2016/08/03 15:12:25 Sum of [3 5 5 1] is 14
2016/08/03 15:12:25 Sum of [2 2 8 7] is 19
2016/08/03 15:12:29 Try to compute sum of [1 0 8 2]
2016/08/03 15:12:29 Try to compute sum of [5 3 7 8]
2016/08/03 15:12:30 Sum of [1 0 8 2] is 11
2016/08/03 15:12:30 Sum of [5 3 7 8] is 23
```
