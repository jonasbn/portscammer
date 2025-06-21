# portscammer

## Introduction

I did a basic Go course entitled ["Go in 3 Weeks"](https://learning.oreilly.com/live-events/go-in-3-weekswith-interactivity/0636920060986/) with [Johnny Boursiquot](https://github.com/jboursiquot) via [O'Reilly Safari](https://www.oreilly.com/publisher/safari-books-online/) back in 2022 and one of the exercises was on network and demonstrated [a port scanner](https://github.com/jboursiquot/portscan) implementation to get to learn how to work with the network.

Based on my newly acquired knowledge I decided to write a tool that could let me monitor for port scarns. It came out of the idea of how do you test a port scanner.

## Usage

Currently it only listens on port `8080` and responds to any TCP connection with a message that includes the port number.

```bash
go run main.go
```

Or if you have the binary built, you can run it like this:

```bash
./portscammer
```

## Build

To build the binary, you can use the following command:

```bash
go build -o portscammer main.go
```
