package nfcli

import (
	"github.com/muthuramanecs03g/lib"
)

var PromptConfig *lib.Prompt

func InitializePrompt() {
	PromptConfig = &lib.Prompt{
		IsEnable:   false,
		IsModule:   false,
		Prefix:     "nfcli# ",
		Suggestion: &MainSuggestion,
		Title:      "nfcli - A Thrift based CLI tool for free5gc NFs",
	}
}
