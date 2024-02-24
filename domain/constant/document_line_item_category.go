package constant

const (
	DocumentLineItemCategoryMain      = "main"
	DocumentLineItemCategorySecondary = "secondary"

	DocumentLineItemCategoryMainValue      = 1
	DocumentLineItemCategorySecondaryValue = 2
)

var DocumentLineItemCategoryName = map[int64]string{
	1: DocumentLineItemCategoryMain,
	2: DocumentLineItemCategorySecondary,
}

var DocumentLineItemCategoryValue = map[string]int64{
	DocumentLineItemCategoryMain:      1,
	DocumentLineItemCategorySecondary: 2,
}
