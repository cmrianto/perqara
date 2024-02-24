package constant

const (
	ShippingCostTypeDraft          = "draft"
	ShippingCostTypeSubmitted      = "submitted"
	ShippingCostTypeDraftValue     = 1
	ShippingCostTypeSubmittedValue = 2
)

var ShippingCostTypeName = map[int64]string{
	1: ShippingCostTypeDraft,
	2: ShippingCostTypeSubmitted,
}

var ShippingCostTypeValue = map[string]int64{
	ShippingCostTypeDraft:     1,
	ShippingCostTypeSubmitted: 2,
}
