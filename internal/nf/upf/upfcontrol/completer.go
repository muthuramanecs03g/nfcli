package upfcontrol

import (
	"github.com/c-bata/go-prompt"
)

var UpfControlSuggestion = []prompt.Suggest{
	{Text: "connect", Description: "Connect to the UPF control plane thrift server"},
	{Text: "exit", Description: "Exit the UPF module"},
}

func CompleterControl(in prompt.Document) []prompt.Suggest {
	return UpfControlSuggestion
}
