package upf

import (
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/muthuramanecs03g/nfcli/internal/nf/upf/upfcontrol"
	"github.com/muthuramanecs03g/nfcli/internal/nf/upf/upfdata"
)

var UpfSuggestion = []prompt.Suggest{
	{Text: "control", Description: "UPF control plane"},
	{Text: "data", Description: "UPF data plane"},
	{Text: "exit", Description: "Exit the UPF module"},
}

func CompleterUPF(in prompt.Document) []prompt.Suggest {
	a := in.TextBeforeCursor()
	var split = strings.Split(a, " ")
	w := in.GetWordBeforeCursor()
	if len(split) > 1 {
		var v = split[0]
		if v == "control" {
			return upfcontrol.CompleterControl(in)
		}
		if v == "data" {
			return upfdata.CompleterData(in)
		}
		return prompt.FilterHasPrefix(UpfSuggestion, v, true)
	}
	return prompt.FilterHasPrefix(UpfSuggestion, w, true)
}
