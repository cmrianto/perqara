package constant

const (
	MarginTypePercentage      = "percentage"
	MarginTypeNominal         = "nominal"
	MarginTypePercentageValue = 1
	MarginTypeNominalValue    = 2
)

var MarginTypeName = map[int64]string{
	1: MarginTypePercentage,
	2: MarginTypeNominal,
}

var MarginTypeValue = map[string]int64{
	MarginTypePercentage: 1,
	MarginTypeNominal:    2,
}
