package constant

const (
	ProofOfCutPphPending        = "pending"
	ProofOfCutPphSubmitted      = "submitted"
	ProofOfCutPphPendingValue   = 1
	ProofOfCutPphSubmittedValue = 2
)

var ProofOfCutPphName = map[int64]string{
	1: ProofOfCutPphPending,
	2: ProofOfCutPphSubmitted,
}

var ProofOfCutPphValue = map[string]int64{
	ProofOfCutPphPending:   1,
	ProofOfCutPphSubmitted: 2,
}
