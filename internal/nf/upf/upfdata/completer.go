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
	{Text: "exit", Description: "Exit the UPF data plane"},
}

var optionHelp = []prompt.Suggest{
	{Text: "-h"},
	{Text: "--help"},
}

var globalOptions = []prompt.Suggest{
	// {Text: "--namespace", Description: "temporarily set the namespace for a request"},
}

var connectOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--ipv4", Description: "UPF data plane Thrift server Ipv4 address"},
	prompt.Suggest{Text: "--port", Description: "UPF data plane Thrfit server listening port"},
}

var statsOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--n3", Description: "UPF data plane N3 port number"},
	prompt.Suggest{Text: "--n6", Description: "UPF data plane N6 port number"},
	prompt.Suggest{Text: "--n9", Description: "UPF data plane N9 port number"},
}

var clearOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--n3", Description: "UPF data plane N3 port number"},
	prompt.Suggest{Text: "--n6", Description: "UPF data plane N6 port number"},
	prompt.Suggest{Text: "--n9", Description: "UPF data plane N9 port number"},
}

func completerConnect(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	// d := strings.Split(in.TextBeforeCursor(), " ")

	// if d[1] ==  {
	// 	return prompt.FilterHasPrefix([]prompt.Suggest{
	// 		{Text: "--ipv4", Description: "Specify the Thrif IPv4 address of UPF"},
	// 	}, a, true)
	// }

	return prompt.FilterHasPrefix([]prompt.Suggest{
		{Text: "--ipv4", Description: "Specify the Thrif IPv4 address of UPF"},
	}, a, true)
}

// func CompleterData(in prompt.Document, promptConfig *lib.Prompt) []prompt.Suggest {
// 	a := in.TextBeforeCursor()
// 	var split = strings.Split(a, " ")
// 	// w := in.GetWordBeforeCursor()
// 	if len(split) > 1 {
// 		var v = split[0]
// 		if v == "connect" {
// 			fmt.Println("CompleterData: connect")
// 			return completerConnect(in)
// 		}
// 		return prompt.FilterHasPrefix(*promptConfig.Suggestion, v, true)
// 	}
// 	return prompt.FilterHasPrefix(*promptConfig.Suggestion, a, true)
// }

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
	//cmd, option, found := getPreviousOption(d)
	_, _, found := getPreviousOption(d)
	if !found {
		return []prompt.Suggest{}, false
	}

	// commands
	// switch cmd {
	// case "connect", "stats", "clear":
	// 	if option == "-c" || option == "--container" {
	// 		cmdArgs := getCommandArgs(d)
	// 		var suggestions []prompt.Suggest
	// 		if cmdArgs == nil || len(cmdArgs) < 2 {
	// 			suggestions = getContainerNamesFromCachedPods(c.client, c.namespace)
	// 		} else {
	// 			suggestions = getContainerName(c.client, c.namespace, cmdArgs[1])
	// 		}
	// 		return prompt.FilterHasPrefix(
	// 			suggestions,
	// 			d.GetWordBeforeCursor(),
	// 			true,
	// 		), true
	// 	}
	// }
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
