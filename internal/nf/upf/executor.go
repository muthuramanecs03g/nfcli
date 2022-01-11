package upf

import (
	"strings"

	"github.com/muthuramanecs03g/nfcli/internal/nf/upf/upfcontrol"
	"github.com/muthuramanecs03g/nfcli/internal/nf/upf/upfdata"
	"github.com/muthuramanecs03g/nfcli/lib"
)

func ExecutorUpf(in string, promptConfig *lib.Prompt) {
	if strings.HasPrefix(in, "control") {
		promptConfig.Suggestion = &upfcontrol.UpfControlSuggestion
		promptConfig.Prefix = "upf>control# "
		promptConfig.SubNf = lib.NF_UPF_CONTROL
		return
	} else if strings.HasPrefix(in, "data") {
		promptConfig.Suggestion = &upfdata.UpfDataSuggestion
		promptConfig.Prefix = "upf>data# "
		promptConfig.SubNf = lib.NF_UPF_DATA
		return
	}

	switch promptConfig.SubNf {
	case lib.NF_UPF_CONTROL:
		upfcontrol.ExecutorControl(in, promptConfig)
	case lib.NF_UPF_DATA:
		upfdata.ExecutorData(in, promptConfig)
	}

	if in == "exit" {
		// fmt.Println("ExecutorUPF: ", in, "SubNf: ", promptConfig.SubNf)
		if promptConfig.SubNf != lib.NF_INVALID {
			// fmt.Println("ExecutorUPF: UPF")
			promptConfig.Suggestion = &UpfSuggestion
			promptConfig.Prefix = "upf# "
			promptConfig.SubNf = lib.NF_INVALID
			return
		}
		promptConfig.Nf = lib.NF_END
		Exit()
	}
}
