package lib

const (
	NF_AMF_STR = "amf"
	NF_SMF_STR = "smf"
	NF_UPF_STR = "upf"
)

const (
	NF_AMF int = iota
	NF_SMF
	NF_UPF

	/* End-of NFs */
	NF_END
)

const (
	NF_UPF_CONTROL int = iota
	NF_UPF_DATA
	NF_INVALID
)
