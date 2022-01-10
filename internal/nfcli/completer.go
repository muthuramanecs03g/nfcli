package nfcli

import (
	"github.com/c-bata/go-prompt"
	"github.com/muthuramanecs03g/nfcli/internal/nf/upf"
	"github.com/muthuramanecs03g/nfcli/lib"
)

const (
	NF_MAIN_STR = "nf"
)

const (
	NF_MAIN = 0xff
)

var MainSuggestion = []prompt.Suggest{
	{Text: "amf", Description: "Connect with AMF"},
	{Text: "smf", Description: "Connect with SMF"},
	{Text: "upf", Description: "Connect with UPF"},
	{Text: "exit", Description: "Exit from NFCli"},
}

// Completer is responsible for the autocompletion of the CLI
func Completer(in prompt.Document) []prompt.Suggest {
	if PromptConfig.IsNf {
		if PromptConfig.Nf == lib.NF_UPF {
			return upf.CompleterUPF(in, PromptConfig)
		}
	}

	w := in.TextBeforeCursor()
	return prompt.FilterHasPrefix(*PromptConfig.Suggestion, w, true)
}
