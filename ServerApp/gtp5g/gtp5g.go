package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/muthuramanecs03g/nfcli/gen-go/UpfGtp5g"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")

	flag.Parse()

	//protocol
	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		Usage()
		os.Exit(1)
	}

	// transport, no secure
	var err error
	var transport thrift.TServerTransport
	transport, err = thrift.NewTServerSocket(*addr)
	if err != nil {
		fmt.Printf("Run: Failed to create a thrift transport: %v", err)
	}

	// Buffered
	var transportFactory thrift.TTransportFactory
	if *buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	// framed
	if *framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	// processor
	handler := &Gtp5gThriftServer{}
	processor := UpfGtp5g.NewUpfGtp5gServiceProcessor(handler)

	fmt.Printf("Run: Starting the Thrift server... on %s\n", *addr)
	// start tcp server
	thriftServer := thrift.NewTSimpleServer4(processor,
		transport,
		transportFactory,
		protocolFactory)
	go thriftServer.Serve()

	if err != nil {
		fmt.Println("error running server:", err)
	}
}
