package nfcli

import (
	"fmt"
	"nfcli/internal/nf"
	"nfcli/internal/nf/upf"
	"os"
	"strings"
)

func Executor(in string) {
	if PromptConfig.Module == nf.NF_UPF {
		upf.Executor(in)
	}

	if strings.HasPrefix(in, "upf") {
		PromptConfig.Suggestion = &upf.UpfSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "upf# "
		PromptConfig.IsModule = true
		PromptConfig.Module = nf.NF_UPF
		upf.Initialize()
		return
	}

	if in == "exit" {
		if PromptConfig.IsModule {
			PromptConfig.Suggestion = &MainSuggestion
			PromptConfig.IsEnable = true
			PromptConfig.Prefix = "nfcli> "
			PromptConfig.IsModule = false
			PromptConfig.Module = NF_MAIN
			return
		}
		fmt.Println("Bye Bye !")
		os.Exit(0)
	}
}
