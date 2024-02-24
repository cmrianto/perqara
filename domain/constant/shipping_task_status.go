package constant

const (
	ShippingTaskStatusIncomplete      = "incomplete"
	ShippingTaskStatusCompleted       = "completed"
	ShippingTaskStatusIncompleteValue = 1
	ShippingTaskStatusCompletedValue  = 2

	ShippingTaskStatusIncompleteLabel = "Incomplete"
	ShippingTaskStatusCompletedLabel  = "Completed"
)

var ShippingTaskStatusName = map[int64]string{
	1: ShippingTaskStatusIncomplete,
	2: ShippingTaskStatusCompleted,
}

var ShippingTaskStatusValue = map[string]int64{
	ShippingTaskStatusIncomplete: 1,
	ShippingTaskStatusCompleted:  2,
}

var ShippingTaskStatusLabel = map[string]string{
	ShippingTaskStatusIncomplete: ShippingTaskStatusIncompleteLabel,
	ShippingTaskStatusCompleted:  ShippingTaskStatusCompletedLabel,
}
