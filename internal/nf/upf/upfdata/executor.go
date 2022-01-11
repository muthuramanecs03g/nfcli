package upfdata

import (
	"context"
	"fmt"
	"strings"

	"github.com/muthuramanecs03g/nfcli/lib"
)

var ctx = context.Background()

// Connect to the Thrift Server
// connect --ipv4 10.20.0.10 --port 9090
func upfDataConnect(in string) {
	cmd := strings.Split(strings.TrimSpace(in), " ")
	l := len(cmd)

	if l < 5 {
		return
	}

	fmt.Println("upfDataConnect: ")
	for idx, str := range cmd {
		fmt.Printf("Idx: %d Str: %s\n", idx, str)
	}

	connectClient(cmd[2], cmd[4])
}

// stats --n3 1 --n4 2 --n6 3
func upfDataStats(in string) {
	cmd := strings.Split(strings.TrimSpace(in), " ")
	l := len(cmd)

	if l < 3 {
		return
	}

	if upfDataClient.upfDataConn == nil {
		fmt.Println("upfDataStats: No connection exists, try connect again")
		return
	}

	// fmt.Println("upfDataStats: Given ")
	// for idx, str := range cmd {
	// 	fmt.Printf("Idx: %d Str: %s\n", idx, str)
	// }

	fmt.Println("upfDataStats: Splits port ")
	split := strings.Split(cmd[2], ",")
	for idx, str := range split {
		fmt.Printf("Idx: %d Str: %s\n", idx, str)
	}

	port, err := lib.StringToInt32(cmd[2])
	if err != nil {
		fmt.Println("upfDataStats: Invalid port: ", cmd[2])
		return
	}

	rsp, err := upfDataClient.upfDataConn.GetStats(ctx, port)
	if err != nil {
		fmt.Printf("upfDataStats: Thrift err: %v", err)
	}

	fmt.Printf("P: %d Rx: %d Tx: %d\n", port, rsp.GetStats().GetRxPktCount(),
		rsp.GetStats().TxPktCount)
}

// clear --n3 1 --n6 2 --n9 3
func upfDataClear(in string) {
	cmd := strings.Split(strings.TrimSpace(in), " ")
	l := len(cmd)

	if l < 3 {
		return
	}

	if upfDataClient.upfDataConn == nil {
		fmt.Println("upfDataStats: No connection exists, try connect again")
		return
	}

	// fmt.Println("upfDataClear: ")
	// for idx, str := range cmd {
	// 	fmt.Printf("Idx: %d Str: %s\n", idx, str)
	// }

	port, err := lib.StringToInt32(cmd[2])
	if err != nil {
		fmt.Println("upfDataStats: Invalid port: ", cmd[2])
		return
	}

	err = upfDataClient.upfDataConn.ClearStats(ctx, port)
	if err != nil {
		fmt.Printf("upfDataStats: Thrift err: %v", err)
	}
}

// close
func upfDataClose(in string) {
	if upfDataClient.upfDataConn == nil {
		fmt.Println("upfDataStats: No connection exists")
		return
	}
	closeClient()
}

func ExecutorData(in string, promptConfig *lib.Prompt) {
	if strings.HasPrefix(in, "exit") {
		Exit()
		return
	}
	if strings.HasPrefix(in, "connect") {
		upfDataConnect(in)
		return
	}

	if strings.HasPrefix(in, "stats") {
		upfDataStats(in)
		return
	}

	if strings.HasPrefix(in, "clear") {
		upfDataClear(in)
		return
	}

	if strings.HasPrefix(in, "close") {
		upfDataClose(in)
		return
	}

	fmt.Println("ExcutorData: Unhandled Given: ", in)
}
