package upfdata

import (
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/muthuramanecs03g/nfcli/lib"
)

var UpfDataSuggestion = []prompt.Suggest{
	{Text: "connect", Description: "Connect to the UPF data plane thrift server"},
	{Text: "stats", Description: "Get the statistics report"},
	{Text: "clear", Description: "Clear the statistics report of ports"},
	{Text: "close", Description: "Close the currently connected UPF data plane"},
	{Text: "status", Description: "Status the currently connected UPF data plane"},
	{Text: "list", Description: "List of connected UPF data plane"},
	{Text: "uplink", Description: "UPF data plane uplink rules"},
	{Text: "downlink", Description: "UPF data plane uplink rules"},
	{Text: "log", Description: "Log the statistics report"},
	{Text: "pcap", Description: "Start/Stop packet capture on port"},
	{Text: "help", Description: "List of UPF data plane commands"},
	{Text: "exit", Description: "Exit the UPF data plane"},
}

var optionHelp = []prompt.Suggest{
	{Text: "help", Description: "List of UPF data plane commands"},
}

var globalOptions = []prompt.Suggest{
	// {Text: "--namespace", Description: "temporarily set the namespace for a request"},
}

var connectOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--ipv4", Description: "UPF data plane Thrift server Ipv4 address"},
	prompt.Suggest{Text: "-i", Description: "UPF data plane Thrift server Ipv4 address"},
	prompt.Suggest{Text: "--port", Description: "UPF data plane Thrfit server listening port"},
	prompt.Suggest{Text: "-p", Description: "UPF data plane Thrfit server listening port"},
}

var statsOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--n3", Description: "UPF data plane N3 port number"},
	prompt.Suggest{Text: "--n6", Description: "UPF data plane N6 port number"},
	prompt.Suggest{Text: "--n9", Description: "UPF data plane N9 port number"},
	prompt.Suggest{Text: "--times", Description: "Number of times query the stats per sec(default 1)"},
	prompt.Suggest{Text: "-t", Description: "Number of times query the stats per sec(default 1)"},
}

var clearOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--n3", Description: "UPF data plane N3 port number"},
	prompt.Suggest{Text: "--n6", Description: "UPF data plane N6 port number"},
	prompt.Suggest{Text: "--n9", Description: "UPF data plane N9 port number"},
}

var closeOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--id", Description: "Identifier of UPF data plane to be close"},
	prompt.Suggest{Text: "-i", Description: "Identifier of UPF data plane to be close"},
	prompt.Suggest{Text: "--all", Description: "Close all connected UPF data plane"},
	prompt.Suggest{Text: "-a", Description: "Close all connected UPF data plane"},
}

var statusOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--id", Description: "Identfier of UPF data plane Thrfit server"},
	prompt.Suggest{Text: "-i", Description: "Identfier of UPF data plane Thrfit server"},
	prompt.Suggest{Text: "--start", Description: "Starting index of UPD data plane Thrfit server"},
	prompt.Suggest{Text: "-s", Description: "Starting index of UPD data plane Thrfit server"},
	prompt.Suggest{Text: "--count", Description: "Number of UPF data plane Thrfit server status to be display"},
	prompt.Suggest{Text: "-c", Description: "Number of UPF data plane Thrfit server status to be display"},
	prompt.Suggest{Text: "--all", Description: "List of connected UPF data plane Thrift sever"},
	prompt.Suggest{Text: "-a", Description: "List of connected UPF data plane Thrift sever"},
}

var listOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--count", Description: "Number of UPF data plane Thrfit server to be display"},
	prompt.Suggest{Text: "-c", Description: "Number of UPF data plane Thrfit server to be display"},
}

var logOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--write", Description: "Write port statistics report on file"},
	prompt.Suggest{Text: "-w", Description: "Write port statistics report on file"},
}

var uplinkFlowOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--id", Description: "Identfier of UPF data plane Thrfit server"},
	prompt.Suggest{Text: "--all", Description: "List all uplink flow rules"},
	prompt.Suggest{Text: "--add", Description: "Add a new link uplink fow rules"},
	prompt.Suggest{Text: "--del", Description: "Delete a new link uplink fow rules"},
	prompt.Suggest{Text: "--session", Description: "Number of Tunnel session established"},
	prompt.Suggest{Text: "--input", Description: "Read rule information from text file"},
}

var downlinkFlowOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--id", Description: "Identfier of UPF data plane Thrfit server"},
	prompt.Suggest{Text: "--all", Description: "List all uplink flow rules"},
	prompt.Suggest{Text: "--add", Description: "Add a new link uplink fow rules"},
	prompt.Suggest{Text: "--del", Description: "Delete a new link uplink fow rules"},
	prompt.Suggest{Text: "--session", Description: "Number of Tunnel session established"},
	prompt.Suggest{Text: "--input", Description: "Read rule information from text file"},
}

var pcapOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--port", Description: "DPDK port number of n3/n6/n9"},
	prompt.Suggest{Text: "--start", Description: "Start port capture"},
	prompt.Suggest{Text: "--stop", Description: "Stop port capture"},
}

func excludeOptions(args []string) ([]string, bool) {
	l := len(args)
	if l == 0 {
		return nil, false
	}
	cmd := args[0]
	filtered := make([]string, 0, l)

	var skipNextArg bool
	for i := 0; i < len(args); i++ {
		if skipNextArg {
			skipNextArg = false
			continue
		}

		if cmd == "logs" && args[i] == "-f" {
			continue
		}

		for _, s := range []string{
			"-s", "--server",
			"--user",
			"-o", "--output",
		} {
			if strings.HasPrefix(args[i], s) {
				if strings.Contains(args[i], "=") {
					// we can specify option value like '-o=json'
					skipNextArg = false
				} else {
					skipNextArg = true
				}
				continue
			}
		}
		if strings.HasPrefix(args[i], "-") {
			continue
		}

		filtered = append(filtered, args[i])
	}
	return filtered, skipNextArg
}

func optionCompleter(args []string, long bool) []prompt.Suggest {
	l := len(args)
	if l <= 1 {
		if long {
			return prompt.FilterHasPrefix(optionHelp, "--", false)
		}
		return optionHelp
	}

	var suggests []prompt.Suggest
	commandArgs, _ := excludeOptions(args)
	switch commandArgs[0] {
	case "connect":
		suggests = connectOptions
	case "stats":
		suggests = statsOptions
	case "clear":
		suggests = clearOptions
	case "close":
		suggests = clearOptions
	case "status":
		suggests = statusOptions
	case "list":
		suggests = listOptions
	case "uplink":
		suggests = uplinkFlowOptions
	case "downlink":
		suggests = downlinkFlowOptions
	case "pcap":
		suggests = pcapOptions
	case "log":
		suggests = logOptions
	default:
		suggests = optionHelp
	}

	suggests = append(suggests, globalOptions...)
	if long {
		return prompt.FilterContains(
			prompt.FilterHasPrefix(suggests, "--", false),
			strings.TrimLeft(args[l-1], "--"),
			true,
		)
	}
	return prompt.FilterContains(suggests, strings.TrimLeft(args[l-1], "-"), true)
}

func getPreviousOption(d prompt.Document) (cmd, option string, found bool) {
	args := strings.Split(d.TextBeforeCursor(), " ")
	l := len(args)
	if l >= 2 {
		option = args[l-2]
	}
	if strings.HasPrefix(option, "-") {
		return args[0], option, true
	}
	return "", "", false
}

func completeOptionArguments(d prompt.Document) ([]prompt.Suggest, bool) {
	cmd, _, found := getPreviousOption(d)
	if !found {
		return []prompt.Suggest{}, false
	}

	// commands
	switch cmd {
	case "help":
		return prompt.FilterHasPrefix(
			UpfDataSuggestion,
			d.GetWordBeforeCursor(),
			true,
		), true
	}
	return []prompt.Suggest{}, false
}

func argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(UpfDataSuggestion, args[0], true)
	}

	first := args[0]
	switch first {
	case "connect":
		// second := args[1]
		// fmt.Println("Second: ", second)
	case "stats":
		// second := args[1]
		// fmt.Println("Second: ", second)
	case "clear":
		// second := args[1]
		// fmt.Println("Second: ", second)
	case "close":
		// second := args[1]
		// fmt.Println("Second: ", second)
	case "help":
		return UpfDataSuggestion
	default:
		return []prompt.Suggest{}
	}
	return []prompt.Suggest{}
}

func CompleterData(in prompt.Document, promptConfig *lib.Prompt) []prompt.Suggest {
	if in.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(in.TextBeforeCursor(), " ")
	w := in.GetWordBeforeCursor()

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(w, "-") {
		return optionCompleter(args, strings.HasPrefix(w, "--"))
	}

	// Return suggestions for option
	if suggests, found := completeOptionArguments(in); found {
		return suggests
	}

	commandArgs, skipNext := excludeOptions(args)
	if skipNext {
		// when type 'get pod -o ', we don't want to complete pods. we want to type 'json' or other.
		// So we need to skip argumentCompleter.
		return []prompt.Suggest{}
	}

	return argumentsCompleter(commandArgs)
}
