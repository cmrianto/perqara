package constant

const (
	PaymentStatusPaid         = "paid"
	PaymentStatusPending      = "pending"
	PaymentStatusPaidValue    = 1
	PaymentStatusPendingValue = 2
)

var PaymentStatusName = map[int64]string{
	1: PaymentStatusPaid,
	2: PaymentStatusPending,
}

var PaymentStatusValue = map[string]int64{
	PaymentStatusPaid:    1,
	PaymentStatusPending: 2,
}
