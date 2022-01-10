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
	{Text: "exit", Description: "Exit the UPF data plane"},
}

func CompleterData(in prompt.Document, promptConfig *lib.Prompt) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	d := in.TextBeforeCursor()
	if len(strings.Split(d, " ")) > 2 {
		return []prompt.Suggest{}
	}
	promptConfig.Suggestion = &UpfDataSuggestion
	return prompt.FilterHasPrefix(*promptConfig.Suggestion, a, true)
}
