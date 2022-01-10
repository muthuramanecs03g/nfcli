package nfcli

import (
	"fmt"
	"os"
	"strings"

	"github.com/muthuramanecs03g/nfcli/internal/nf/upf"
	"github.com/muthuramanecs03g/nfcli/lib"
)

func Executor(in string) {
	if PromptConfig.Nf == lib.NF_UPF {
		fmt.Println("Main: UPF")
		upf.ExecutorUpf(in, PromptConfig)
	}

	if strings.HasPrefix(in, "upf") {
		fmt.Println("Main: UPF has prefix")
		PromptConfig.Suggestion = &upf.UpfSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "upf# "
		PromptConfig.IsNf = true
		PromptConfig.Nf = lib.NF_UPF
		return
	}

	if PromptConfig.IsNf && PromptConfig.Nf != lib.NF_END {
		return
	}

	if in == "exit" {
		fmt.Println("Main: exit")
		if PromptConfig.IsNf {
			fmt.Println("Main: Changed to main")
			PromptConfig.Suggestion = &MainSuggestion
			PromptConfig.IsEnable = true
			PromptConfig.Prefix = "nfcli# "
			PromptConfig.IsNf = false
			PromptConfig.Nf = NF_MAIN
			return
		}
		fmt.Println("Bye Bye !!!")
		os.Exit(0)
	}
}
