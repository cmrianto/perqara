package constant

const (
	TrackingProgressCategorySeaDelivery       = "sea_delivery"
	TrackingProgressCategorySeaTransit        = "sea_transit"
	TrackingProgressCategoryLandTransit       = "land_transit"
	TrackingProgressCategoryLandDelivery      = "land_delivery"
	TrackingProgressCategorySeaDeliveryValue  = 1
	TrackingProgressCategorySeaTransitValue   = 2
	TrackingProgressCategoryLandTransitValue  = 3
	TrackingProgressCategoryLandDeliveryValue = 4
)

var TrackingProgressCategoryeName = map[int64]string{
	1: TrackingProgressCategorySeaDelivery,
	2: TrackingProgressCategorySeaTransit,
	3: TrackingProgressCategoryLandTransit,
	4: TrackingProgressCategoryLandDelivery,
}

var TrackingProgressCategoryeValue = map[string]int64{
	TrackingProgressCategorySeaDelivery:  1,
	TrackingProgressCategorySeaTransit:   2,
	TrackingProgressCategoryLandTransit:  3,
	TrackingProgressCategoryLandDelivery: 4,
}
