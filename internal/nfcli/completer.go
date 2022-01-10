package nfcli

import (
	"github.com/c-bata/go-prompt"
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
