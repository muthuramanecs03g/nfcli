package upfdata

import (
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/muthuramanecs03g/nfcli/gen-go/Upf"
)

func GetClient(addr string) *Upf.UpfServiceClient {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
	}

	// protocol
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	// no buffered
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("error running client:", err)
	}

	if err := transport.Open(); err != nil {
		fmt.Println("error running client:", err)
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := Upf.NewUpfServiceClient(thrift.NewTStandardClient(iprot, oprot))
	return client
}
