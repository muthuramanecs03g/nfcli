package upfdata

import (
	"github.com/c-bata/go-prompt"
)

var UpfDataSuggestion = []prompt.Suggest{
	{Text: "connect", Description: "Connect to the UPF data plane thrift server"},
	{Text: "stats", Description: "Get the statistics report"},
	{Text: "clear", Description: "Clear the statistics report of ports"},
	{Text: "exit", Description: "Exit the UPF data plane"},
}

func CompleterData(in prompt.Document) []prompt.Suggest {
	return UpfDataSuggestion
}
