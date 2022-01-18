package upfdata

import (
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/muthuramanecs03g/nfcli/gen-go/Upf"
)

func connectClient(ipv4, port string) {
	var transport thrift.TTransport
	var err error
	addr := ipv4 + ":" + port
	upfDataClient.transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
	}

	// protocol
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	// no buffered
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	transport, err = transportFactory.GetTransport(upfDataClient.transport)
	if err != nil {
		fmt.Println("error running client:", err)
	}

	if err := transport.Open(); err != nil {
		fmt.Println("error running client:", err)
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	upfDataClient.upfDataConn = Upf.NewUpfServiceClient(thrift.NewTStandardClient(iprot, oprot))
	fmt.Println("Connected")
}

func closeClient() {
	if upfDataClient != nil && upfDataClient.upfDataConn != nil {
		if upfDataClient.transport.IsOpen() == true {
			upfDataClient.transport.Close()
		}
		upfDataClient.upfDataConn = nil
	}
	fmt.Println("Dis-Connected")
}
