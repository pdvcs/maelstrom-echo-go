# A Simple Echo Server in Go

This is a Go version of the Echo Server
[described here](https://github.com/jepsen-io/maelstrom/blob/main/doc/02-echo/index.md).

Requires Go version 1.18+ and Linux (to run maelstrom).
1.18 is required because one of the utility functions uses generics.

## Build and Deploy

```shell
go build -o go-echo
cp go-echo ~/repos/maelstrom-bin/
```

## Run

```shell
./maelstrom test -w echo --bin ./go-echo --time-limit 5
```
