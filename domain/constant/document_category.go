package constant

const (
	DocumentCategoryPickupLetter             = "pickup_letter"
	DocumentCategoryShippingInstruction      = "shipping_instruction"
	DocumentCategoryChronologyLetter         = "chronology_letter"
	DocumentCategoryReceiptOfGoods           = "receipt_of_goods"
	DocumentCategoryShippingInstructionTemas = "shipping_instruction_temas"
	DocumentCategoryPackingList              = "packing_list"
	DocumentCategoryBast                     = "bast"

	DocumentCategoryPickupLetterValue             = 1
	DocumentCategoryShippingInstructionValue      = 2
	DocumentCategoryChronologyLetterValue         = 3
	DocumentCategoryReceiptOfGoodsValue           = 4
	DocumentCategoryShippingInstructionTemasValue = 5
	DocumentCategoryPackingListValue              = 6
	DocumentCategoryBastValue                     = 7

	DocumentCategoryPickupLetterLabel             = "Surat Pengambilan Barang"
	DocumentCategoryShippingInstructionLabel      = "Shipping Instruction"
	DocumentCategoryChronologyLetterLabel         = "Surat Kronologi"
	DocumentCategoryReceiptOfGoodsLabel           = "Tanda Terima Barang"
	DocumentCategoryShippingInstructionTemasLabel = "Shipping Instruction TEMAS"
	DocumentCategoryPackingListLabel              = "Packing List"
	DocumentCategoryBastLabel                     = "Berita Penyerahan Barang"
)

var DocumentCategoryName = map[int64]string{
	1: DocumentCategoryPickupLetter,
	2: DocumentCategoryShippingInstruction,
	3: DocumentCategoryChronologyLetter,
	4: DocumentCategoryReceiptOfGoods,
	5: DocumentCategoryShippingInstructionTemas,
	6: DocumentCategoryPackingList,
	7: DocumentCategoryBast,
}

var DocumentCategoryValue = map[string]int64{
	DocumentCategoryPickupLetter:             1,
	DocumentCategoryShippingInstruction:      2,
	DocumentCategoryChronologyLetter:         3,
	DocumentCategoryReceiptOfGoods:           4,
	DocumentCategoryShippingInstructionTemas: 5,
	DocumentCategoryPackingList:              6,
	DocumentCategoryBast:                     7,
}

var DocumentCategoryLabel = map[string]string{
	DocumentCategoryPickupLetter:             DocumentCategoryPickupLetterLabel,
	DocumentCategoryShippingInstruction:      DocumentCategoryShippingInstructionLabel,
	DocumentCategoryChronologyLetter:         DocumentCategoryChronologyLetterLabel,
	DocumentCategoryReceiptOfGoods:           DocumentCategoryReceiptOfGoodsLabel,
	DocumentCategoryShippingInstructionTemas: DocumentCategoryShippingInstructionTemasLabel,
	DocumentCategoryPackingList:              DocumentCategoryPackingListLabel,
	DocumentCategoryBast:                     DocumentCategoryBastLabel,
}
