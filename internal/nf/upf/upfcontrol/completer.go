package upfcontrol

import (
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/muthuramanecs03g/nfcli/lib"
)

var UpfControlSuggestion = []prompt.Suggest{
	{Text: "connect", Description: "Connect to the UPF control plane thrift server"},
	{Text: "exit", Description: "Exit the UPF module"},
}

func CompleterControl(in prompt.Document, promptConfig *lib.Prompt) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	d := in.TextBeforeCursor()
	if len(strings.Split(d, " ")) > 2 {
		return []prompt.Suggest{}
	}
	promptConfig.Suggestion = &UpfControlSuggestion
	return prompt.FilterHasPrefix(*promptConfig.Suggestion, a, true)
}
