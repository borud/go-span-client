# go-span-client

This repository contains a simple DTLS client library for communicating with Span as an IoT device. If you want to access the Span API for managing Span resources, the library you want is the [Go Span API Client](https://github.com/lab5e/go-spanapi).

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/borud/go-span-client)

## Minimal example

```go
package main

import (
    "log"
    "time"
    client "github.com/borud/go-span-client"
)

func main() {
    client, err := client.Connect(client.NewDefaultConfig())
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
   
    n, err := client.Write([]byte("this is a test"), time.Second)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("wrote %d bytes", n)
   
    buffer := make([]byte, 1024)
   
    n, err = client.Read(buffer, time.Second*5)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("read %d bytes: [%s]", n, string(buffer))
}
```
