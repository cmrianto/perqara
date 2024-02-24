package constant

const (
	InvoicePaymentMethodTransfer      = "transfer"
	InvoicePaymentMethodCash          = "cash"
	InvoicePaymentMethodTransferValue = 1
	InvoicePaymentMethodCashValue     = 2
)

var InvoicePaymentMethodName = map[int64]string{
	1: InvoicePaymentMethodTransfer,
	2: InvoicePaymentMethodCash,
}

var InvoicePaymentMethodValue = map[string]int64{
	InvoicePaymentMethodTransfer: 1,
	InvoicePaymentMethodCash:     2,
}
