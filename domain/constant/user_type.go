package constant

const (
	UserTypeInternal      = "internal"
	UserTypeExternal      = "external"
	UserTypeInternalValue = 1
	UserTypeExternalValue = 2
)

var UserTypeName = map[int64]string{
	1: UserTypeInternal,
	2: UserTypeExternal,
}

var UserTypeValue = map[string]int64{
	UserTypeInternal: 1,
	UserTypeExternal: 2,
}
