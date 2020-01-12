package main

import (
	"context"
	"example"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/apache/thrift/lib/go/thrift"
)

const (
	HOST = "127.0.0.1"
	PORT = "8080"
)

func main() {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	transport, _ := transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := example.NewFormatDataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST+":"+PORT)
	}
	defer transport.Close()
	ctx, _ := context.WithCancel(context.Background())

	data := example.Data{Text: os.Args[1]}
	d, err := client.ArtificialIdiot(ctx, &data)
	fmt.Println(d.Text)
}
