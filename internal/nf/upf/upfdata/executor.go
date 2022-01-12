package upfdata

import (
	"context"
	"fmt"
	"strings"

	"github.com/muthuramanecs03g/nfcli/gen-go/Upf"
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

func dumpPortStatistics(stats *Upf.PortStats) {
	/* Rx Handler */
	fmt.Println("Dump Rx Packet Count: ")
	fmt.Println(stats.RxPktCnt)
	fmt.Println("Dump Rx Packet Drop Count: ")
	fmt.Println(stats.RxPktDropCnt)
	/* Tx Handler*/
	/* Non-Qfi */
	fmt.Println("Tx NonQfi Packet Count: ", stats.TxNQfiPktCnt)
	fmt.Println("Tx NonQfi Packet Drop Count: ", stats.TxNQfiPktDropCnt)
	/* Dropper, WRED */
	fmt.Println("Dump Dropper Packets: ")
	fmt.Println(stats.DropperPkts)
	fmt.Println("Dump Dropper Packet Drop Count: ")
	fmt.Println(stats.DropperDrops)
	/* Scheduler */
	fmt.Println("Dump Scheduler Queue Drop Count: ")
	fmt.Println(stats.SchedulerQDropCnt)
	/* DPDK Tx Queue */
	fmt.Println("Dump Tx Packet Count: ")
	fmt.Println(stats.TxPktCnt)
	fmt.Println("Dump Tx Packet Drop Count: ")
	fmt.Println(stats.TxPktDropCnt)
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

	// fmt.Println("upfDataStats: Splits port ")
	// split := strings.Split(cmd[2], ",")
	// for idx, str := range split {
	// 	fmt.Printf("Idx: %d Str: %s\n", idx, str)
	// }

	port, err := lib.StringToInt32(cmd[2])
	if err != nil {
		fmt.Println("upfDataStats: Invalid port: ", cmd[2])
		return
	}

	rsp, err := upfDataClient.upfDataConn.GetStats(ctx, port)
	if err != nil {
		fmt.Printf("upfDataStats: Thrift err: %v", err)
		return
	}
	if rsp.Stats != nil {
		dumpPortStatistics(rsp.Stats)
	}
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
		return
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
