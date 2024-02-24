package constant

const (
	InvoiceStatusPublished      = "published"
	InvoiceStatusCanceled       = "canceled"
	InvoiceStatusPublishedValue = 1
	InvoiceStatusCanceledValue  = 2
)

var InvoiceStatusName = map[int64]string{
	1: InvoiceStatusPublished,
	2: InvoiceStatusCanceled,
}

var InvoiceStatusValue = map[string]int64{
	InvoiceStatusPublished: 1,
	InvoiceStatusCanceled:  2,
}
