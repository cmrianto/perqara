package constant

const (
	ShippingTrackingStatusInProgress      = "in_progress"
	ShippingTrackingStatusDelivered       = "delivered"
	ShippingTrackingStatusInProgressValue = 1
	ShippingTrackingStatusDeliveredValue  = 2

	ShippingTrackingStatusInProgressLabel = "In Progress"
	ShippingTrackingStatusDeliveredLabel  = "Delivered"
)

var ShippingTrackingStatusName = map[int64]string{
	1: ShippingTrackingStatusInProgress,
	2: ShippingTrackingStatusDelivered,
}

var ShippingTrackingStatusValue = map[string]int64{
	ShippingTrackingStatusInProgress: 1,
	ShippingTrackingStatusDelivered:  2,
}

var ShippingTrackingStatusLabel = map[string]string{
	ShippingTrackingStatusInProgress: ShippingTrackingStatusInProgressLabel,
	ShippingTrackingStatusDelivered:  ShippingTrackingStatusDeliveredLabel,
}
