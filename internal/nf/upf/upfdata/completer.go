package upfdata

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/muthuramanecs03g/nfcli/lib"
)

var UpfDataSuggestion = []prompt.Suggest{
	{Text: "connect", Description: "Connect to the UPF data plane thrift server"},
	{Text: "stats", Description: "Get the statistics report"},
	{Text: "clear", Description: "Clear the statistics report of ports"},
	{Text: "exit", Description: "Exit the UPF data plane"},
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
func CompleterData(in prompt.Document, promptConfig *lib.Prompt) []prompt.Suggest {
	a := in.TextBeforeCursor()
	var split = strings.Split(a, " ")
	// w := in.GetWordBeforeCursor()
	if len(split) > 1 {
		var v = split[0]
		if v == "connect" {
			fmt.Println("CompleterData: connect")
			return completerConnect(in)
		}
		return prompt.FilterHasPrefix(*promptConfig.Suggestion, v, true)
	}
	return prompt.FilterHasPrefix(*promptConfig.Suggestion, a, true)
}
