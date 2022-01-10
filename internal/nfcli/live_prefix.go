package nfcli

func ChangeLivePrefix() (string, bool) {
	return PromptConfig.Prefix, PromptConfig.IsEnable
}
