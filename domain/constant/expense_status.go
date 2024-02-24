package constant

const (
	ExpenseStatusPublished      = "published"
	ExpenseStatusCanceled       = "canceled"
	ExpenseStatusPublishedValue = 1
	ExpenseStatusCanceledValue  = 2
)

var ExpenseStatusName = map[int64]string{
	1: ExpenseStatusPublished,
	2: ExpenseStatusCanceled,
}

var ExpenseStatusValue = map[string]int64{
	ExpenseStatusPublished: 1,
	ExpenseStatusCanceled:  2,
}
