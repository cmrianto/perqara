package constant

const (
	MenuPermissionCanView   = "can view"
	MenuPermissionCanCreate = "can create"
	MenuPermissionCanEdit   = "can edit"
	MenuPermissionCanDelete = "can delete"
)

var MenuPermissionList = []string{
	MenuPermissionCanView,
}

var MenuPermissionDetail = []string{
	MenuPermissionCanView,
	MenuPermissionCanEdit,
}

var MenuPermissionCreate = []string{
	MenuPermissionCanView,
	MenuPermissionCanCreate,
}

var MenuPermissionDelete = []string{
	MenuPermissionCanView,
	MenuPermissionCanEdit,
	MenuPermissionCanDelete,
}

var MenuPermissionUpdate = []string{
	MenuPermissionCanView,
	MenuPermissionCanEdit,
	MenuPermissionCanCreate,
}
