package main

import (
	"context"
	"example"
	"fmt"
	"log"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"
)

type FormatDataImpl struct{}

func (fdi *FormatDataImpl) ArtificialIdiot(ctx context.Context, data *example.Data) (r *example.Data, err error) {
	var rData example.Data
	rData.Text = strings.Trim(data.Text, "么吗？?")

	return &rData, nil
}

const (
	HOST = "127.0.0.1"
	PORT = "8080"
)

func main() {

	handler := &FormatDataImpl{}
	processor := example.NewFormatDataProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(HOST + ":" + PORT)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", HOST+":"+PORT)
	server.Serve()
}
