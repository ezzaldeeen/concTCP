# concTCP

This repository contains an implementation of Concurrent TCP. 
It enables concurrent handling of multiple TCP connections.


## Usage

```bash
go run server.go
```

```bash
# you can run multiple clients
# in different sessions 
go run client.go
```

```bash
# client:
Hello Server # msg
2023/06/29 02:22:52 127.0.0.1:50418 -> 127.0.0.1:8000
2023/06/29 02:22:52 number of written bytes: 11
2023/06/29 02:22:52 recevied: Server Received
...
# terminate the session
```

```bash
# server:
2023/06/29 02:22:52 received from - 127.0.0.1:50418: Hello Server
2023/06/29 02:23:14 connection closed by:  127.0.0.1:50418 
```

