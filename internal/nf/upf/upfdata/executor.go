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
func upfDataConnect(cmd []string) {
	l := len(cmd)
	if l < 5 {
		return
	}

	// fmt.Println("upfDataConnect: ")
	// for idx, str := range cmd {
	// 	fmt.Printf("Idx: %d Str: %s\n", idx, str)
	// }

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

// stats --n3 1 --n6 2 --n9 3 --times 1
func upfDataStats(cmd []string) {
	l := len(cmd)
	if l < 3 {
		return
	}

	if upfDataClient.upfDataConn == nil {
		fmt.Println("upfDataStats: No connection exists, try connect command")
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
func upfDataClear(cmd []string) {
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
func upfDataClose(cmd []string) {
	if upfDataClient.upfDataConn == nil {
		fmt.Println("upfDataStats: No connection exists")
		return
	}
	closeClient()
}

func upfDataHelp(cmd []string) {
	l := len(cmd)
	if l < 2 {
		return
	}
	switch cmd[1] {
	case "connect":
		fmt.Println("Description: Connect to the UPF data plane thrift server")
		fmt.Println("Usage: connect --ipv4 10.10.0.10 --port 9090")
	case "stats":
		fmt.Println("Description: Get the statistics report of data plane ports")
		fmt.Println("Usage: stats --n3 1 --n6 2 --n9 3 --times 10")
	case "clear":
		fmt.Println("Description: Clear the statistics report of data plane ports")
		fmt.Println("Usage: clear --n3 1 --n6 2 --n9 3 ")
	case "close":
		fmt.Println("Description: Close the connected UPF data plane")
		fmt.Println("Usage: close --id 1 --all")
	case "status":
		fmt.Println("Description: Status the currently connected UPF data plane")
		fmt.Println("Usage: status --id 1 --all")
	case "list":
		fmt.Println("Description: List of connected UPF data plane")
		fmt.Println("Usage: list --id 0 --count 2")
	case "log":
		fmt.Println("Description: Log the port statistics report")
		fmt.Println("Usage: log --write filename.txt")
	case "exit":
		fmt.Println("Description: Exit from the UPF data plane")
		fmt.Println("Usage: exit")
	default:
		fmt.Println("Unrecognized the command: ", cmd[1])
	}
}

func ExecutorData(in string, promptConfig *lib.Prompt) {
	args := strings.Split(strings.TrimSpace(in), " ")
	switch args[0] {
	case "exit":
		Exit()
	case "connect":
		upfDataConnect(args)
	case "stats":
		upfDataStats(args)
	case "clear":
		upfDataClear(args)
	case "close":
		upfDataClose(args)
	case "help":
		upfDataHelp(args)
	default:
		fmt.Println("ExcutorData: Unhandled command: ", in)
	}
}
