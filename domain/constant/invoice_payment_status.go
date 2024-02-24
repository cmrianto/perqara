package constant

const (
	InvoicePaymentStatusPublished      = "published"
	InvoicePaymentStatusCanceled       = "canceled"
	InvoicePaymentStatusPublishedValue = 1
	InvoicePaymentStatusCanceledValue  = 2
)

var InvoicePaymentStatusName = map[int64]string{
	1: InvoicePaymentStatusPublished,
	2: InvoicePaymentStatusCanceled,
}

var InvoicePaymentStatusValue = map[string]int64{
	InvoicePaymentStatusPublished: 1,
	InvoicePaymentStatusCanceled:  2,
}
