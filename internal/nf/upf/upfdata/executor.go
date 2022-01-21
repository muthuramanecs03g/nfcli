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
	if l < 3 { // Minimum command lengths
		fmt.Printf("upfDataConnect: %v\n", cmd)
		return
	}

	var ipv4 string = ""
	var port string = "9090"
	// fmt.Println("upfDataConnect: ")
	// for idx, str := range cmd {
	// 	fmt.Printf("Idx: %d Str: %s\n", idx, str)
	// }
	for idx, str := range cmd {
		switch str {
		case "--ipv4", "-i":
			ipv4 = cmd[idx+1]
		case "--port", "-p":
			port = cmd[idx+1]
		default:
			continue
		}
	}
	// fmt.Println("upfDataConnect: Ip: ", ipv4, " Port: ", port)
	if ipv4 != "" {
		connectClient(ipv4, port)
	}
}

func dumpPortStatistics(stats *Upf.PortStats) {
	/* Rx Handler */
	fmt.Printf("--------------Rx Packet --------------\n")
	fmt.Printf("Count: %v\n", stats.RxPktCnt)
	fmt.Printf("Drop: %v\n", stats.RxPktDropCnt)
	// TX Handler
	// Non-Qfi
	fmt.Printf("--------------Non-QFI Tx Packet --------------\n")
	fmt.Printf("Count: %d\n", stats.TxNQfiPktCnt)
	fmt.Printf("Drop: %d\n", stats.TxNQfiPktDropCnt)
	// Dropper WRED
	fmt.Printf("--------------Dropper WRED Packet --------------\n")
	fmt.Printf("Count: %v\n", stats.DropperPkts)
	fmt.Printf("Drop: %v\n", stats.DropperDrops)
	/* Scheduler */
	fmt.Printf("--------------Scheduler Queue Drop --------------\n")
	fmt.Printf("Drop: %v\n", stats.SchedulerQDropCnt)
	/* DPDK Tx Queue */
	fmt.Printf("--------------DPDK Tx Packet --------------\n")
	fmt.Printf("Count: %v\n", stats.TxPktCnt)
	fmt.Printf("Drop: %v\n", stats.TxPktDropCnt)
	fmt.Printf("==================================================\n")
}

func getPortStats(name, portStr, times string) {
	port, err := lib.StringToInt32(portStr)
	if err != nil {
		fmt.Println("upfDataStats: Invalid port: ", portStr, " Times: ", times)
		return
	}

	rsp, err := upfDataClient.upfDataConn.GetStats(ctx, port)
	if err != nil {
		fmt.Printf("upfDataStats: Thrift err: %v", err)
		return
	}
	if rsp.Stats != nil {
		fmt.Println(name + ": " + portStr)
		dumpPortStatistics(rsp.Stats)
	}
}

// stats --n3 1 --n6 2 --n9 3 --times 1
func upfDataStats(cmd []string) {
	l := len(cmd)
	if l < 3 { // Minimum command length
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
	var n3 string = ""
	var n6 string = ""
	var n9 string = ""
	var times string = "1"
	for idx, str := range cmd {
		next := idx + 1
		switch str {
		case "--n3":
			if l > next {
				n3 = cmd[next]
			}
		case "--n6":
			if l > next {
				n6 = cmd[next]
			}
		case "--n9":
			if l > next {
				n9 = cmd[next]
			}
		case "--times", "-t":
			if l > next {
				times = cmd[next]
			}
		default:
			continue
		}
	}

	if n3 != "" {
		getPortStats("N3", n3, times)
	}
	if n6 != "" {
		getPortStats("N6", n6, times)
	}
	if n9 != "" {
		getPortStats("N9", n9, times)
	}
}

func clearPortStats(portStr string) {
	port, err := lib.StringToInt32(portStr)
	if err != nil {
		fmt.Println("upfDataStats: Invalid port: ", portStr)
		return
	}

	err = upfDataClient.upfDataConn.ClearStats(ctx, port)
	if err != nil {
		fmt.Printf("upfDataStats: Thrift err: %v", err)
		return
	}
}

// clear --n3 1 --n6 2 --n9 3
func upfDataClear(cmd []string) {
	l := len(cmd)
	if l < 3 { // Minimum command length
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
	var n3 string = ""
	var n6 string = ""
	var n9 string = ""
	for idx, str := range cmd {
		next := idx + 1
		switch str {
		case "--n3":
			if l > next {
				n3 = cmd[next]
			}
		case "--n6":
			if l > next {
				n6 = cmd[next]
			}
		case "--n9":
			if l > next {
				n9 = cmd[next]
			}
		default:
			continue
		}
	}

	if n3 != "" {
		clearPortStats(n3)
	}
	if n6 != "" {
		clearPortStats(n6)
	}
	if n9 != "" {
		clearPortStats(n9)
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

func upfDataStatus(cmd []string) {
	if upfDataClient.upfDataConn == nil {
		fmt.Println("upfDataStats: No connection exists")
		return
	}
	if upfDataClient.getDataClientStatus() == true {
		fmt.Println("upfDataStatus: Active")
		return
	}
	fmt.Println("upfDataStatus: In-Active")
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
	case "status":
		upfDataStatus(args)
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
