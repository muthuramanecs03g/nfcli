package upf

import (
	"strings"

	"github.com/muthuramanecs03g/nfcli/internal/nf"
	"github.com/muthuramanecs03g/nfcli/internal/nf/upf/upfcontrol"
	"github.com/muthuramanecs03g/nfcli/internal/nf/upf/upfdata"
)

func ExecutorUpf(in string) {
	if strings.HasPrefix(in, "control") {
		PromptConfig.Suggestion = &upfcontrol.UpfControlSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "upf>control# "
		return
	} else if strings.HasPrefix(in, "data") {
		PromptConfig.Suggestion = &upfdata.UpfDataSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "upf>data# "
		return
	}

	if strings.HasPrefix(in, "exit") {

	}

	if PromptConfig.SubModule == nf.NF_UPF_CONTROL {
		upfcontrol.ExecutorControl(in)
	} else if PromptConfig.SubModule == nf.NF_UPF_DATA {
		upfdata.ExecutorData(in)
	}
}
