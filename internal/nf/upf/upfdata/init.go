package upfdata

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/muthuramanecs03g/nfcli/gen-go/Upf"
)

type upfData struct {
	transport   thrift.TTransport
	upfDataConn *Upf.UpfServiceClient
}

var upfDataClient *upfData

func init() {
	upfDataClient = &upfData{
		transport:   nil,
		upfDataConn: nil,
	}
}

func Exit() {
	if upfDataClient != nil && upfDataClient.upfDataConn != nil {
		if upfDataClient.transport.IsOpen() == true {
			upfDataClient.transport.Close()
		}
		upfDataClient.upfDataConn = nil
	}
}
