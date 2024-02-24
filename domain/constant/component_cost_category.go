package constant

const (
	ComponentCostCategoryCostOfOrigin              = "cost_of_origin"
	ComponentCostCategoryFreightOfCost             = "freight_of_cost"
	ComponentCostCategoryCostFinalDestination      = "cost_final_destination"
	ComponentCostCategoryCostOfOriginValue         = 1
	ComponentCostCategoryFreightOfCostValue        = 2
	ComponentCostCategoryCostFinalDestinationValue = 3
)

var ComponentCostCategoryName = map[int64]string{
	1: ComponentCostCategoryCostOfOrigin,
	2: ComponentCostCategoryFreightOfCost,
	3: ComponentCostCategoryCostFinalDestination,
}

var ComponentCostCategoryValue = map[string]int64{
	ComponentCostCategoryCostOfOrigin:         1,
	ComponentCostCategoryFreightOfCost:        2,
	ComponentCostCategoryCostFinalDestination: 3,
}
