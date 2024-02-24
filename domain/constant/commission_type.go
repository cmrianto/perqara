package constant

const (
	CommissionTypePercentage      = "percentage"
	CommissionTypeNominal         = "nominal"
	CommissionTypePercentageValue = 1
	CommissionTypeNominalValue    = 2
)

var CommissionTypeName = map[int64]string{
	1: CommissionTypePercentage,
	2: CommissionTypeNominal,
}

var CommissionTypeValue = map[string]int64{
	CommissionTypePercentage: 1,
	CommissionTypeNominal:    2,
}
