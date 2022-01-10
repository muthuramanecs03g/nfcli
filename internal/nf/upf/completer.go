package upf

import "github.com/c-bata/go-prompt"

var UpfSuggestion = []prompt.Suggest{
	{Text: "control", Description: "UPF control plane"},
	{Text: "data", Description: "UPF data plane"},
	{Text: "exit", Description: "Exit the UPF module"},
}
