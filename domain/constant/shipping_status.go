package constant

const (
	ShippingStatusInProgress      = "in_progress"
	ShippingStatusDelivered       = "delivered"
	ShippingStatusInProgressValue = 1
	ShippingStatusDeliveredValue  = 2

	ShippingStatusInProgressLabel = "In Progress"
	ShippingStatusDeliveredLabel  = "Delivered"
)

var ShippingStatusName = map[int64]string{
	1: ShippingStatusInProgress,
	2: ShippingStatusDelivered,
}

var ShippingStatusValue = map[string]int64{
	ShippingStatusInProgress: 1,
	ShippingStatusDelivered:  2,
}

var ShippingStatusLabel = map[string]string{
	ShippingStatusInProgress: ShippingStatusInProgressLabel,
	ShippingStatusDelivered:  ShippingStatusDeliveredLabel,
}
