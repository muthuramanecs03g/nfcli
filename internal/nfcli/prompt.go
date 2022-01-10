package nfcli

import (
	"github.com/muthuramanecs03g/nfcli/lib"
)

var PromptConfig *lib.Prompt

func InitializePrompt() {
	PromptConfig = &lib.Prompt{
		IsEnable:   false,
		IsNf:       false,
		Prefix:     "nfcli# ",
		Suggestion: &MainSuggestion,
		Title:      "nfcli - A Thrift based CLI tool for free5gc NFs",
		Nf:         lib.NF_END,
		SubNf:      lib.NF_INVALID,
	}
}
