package upfdata

import (
	"fmt"
	"strings"

	"github.com/muthuramanecs03g/nfcli/lib"
)

// connect --ipv4 10.20.0.10 --port 9090
func upfDataConnect(in string) {
	cmd := strings.Split(strings.TrimSpace(in), " ")
	l := len(cmd)

	if l < 3 {
		return
	}

	fmt.Println("upfDataConnect: ")
	for idx, str := range cmd {
		fmt.Printf("Idx: %d Str: %s\n", idx, str)
	}
}

// stats --port 1
func upfDataStats(in string) {
	cmd := strings.Split(strings.TrimSpace(in), " ")
	l := len(cmd)

	if l < 3 {
		return
	}

	fmt.Println("upfDataStats: Given ")
	for idx, str := range cmd {
		fmt.Printf("Idx: %d Str: %s\n", idx, str)
	}

	fmt.Println("upfDataStats: Splits port ")
	split := strings.Split(cmd[2], ",")
	for idx, str := range split {
		fmt.Printf("Idx: %d Str: %s\n", idx, str)
	}
}

// clear --port 1
func upfDataClear(in string) {
	cmd := strings.Split(strings.TrimSpace(in), " ")
	l := len(cmd)

	if l < 3 {
		return
	}

	fmt.Println("upfDataClear: ")
	for idx, str := range cmd {
		fmt.Printf("Idx: %d Str: %s\n", idx, str)
	}
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

	fmt.Println("ExcutorData: Unhandled Given: ", in)
}
