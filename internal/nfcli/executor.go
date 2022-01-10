package nfcli

import (
	"fmt"
	"os"
	"strings"

	"github.com/muthuramanecs03g/nfcli/internal/nf"
	"github.com/muthuramanecs03g/nfcli/internal/nf/upf"
)

func Executor(in string) {
	if PromptConfig.Module == nf.NF_UPF {

		upf.ExecutorUpf(in)
	}

	if strings.HasPrefix(in, "upf") {
		PromptConfig.Suggestion = &upf.UpfSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "upf# "
		PromptConfig.IsModule = true
		PromptConfig.Module = nf.NF_UPF
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
		upf.Exit()
		fmt.Println("Bye Bye !")
		os.Exit(0)
	}
}
