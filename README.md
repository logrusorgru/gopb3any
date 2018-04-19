gopb3any
========

Example how to use protocol buffers v3.0.0-alpha-2 Any type in golang

### Project tree

```
# --- our types to store
/cat
/user
# --- lifo in-memory storage
/lis
# --- required for lifo storage, example struct
/msg
# --- example program
main.go
```

### Prepare

- Learn about Protocol Buffers [here](https://developers.google.com/protocol-buffers/).
- Read about proto3 and get release v3.0.0-alpha-2+ [here](https://github.com/google/protobuf/releases)
- Install it
- Get golang protocol buffers
  `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`

### How to use

Get
```bash
go get -d guthub.com/logrusorgru/gopb3any
```

See sources

Generate golang messages
```
cd $GOPATH/src/guthub.com/logrusorgru/gopb3any
protoc --go_out=. msg/*.proto
protoc --go_out=. user/*.proto
protoc --go_out=. cat/*.proto
```

Run
```
go run main.go
```

### Licensing

Copyright &copy; 2015 Konstantin Ivanov <kostyarin.ivanov@gmail.com>  
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file for more details.

