package main

import (
	"fmt"

	"github.com/muthuramanecs03g/nfcli/internal/nfcli"
)

func main() {
	fmt.Println("free5GC Thrift NF CLI version: 0.1.0")
	nfcli.Initialize()
	nfcli.Run()
}
