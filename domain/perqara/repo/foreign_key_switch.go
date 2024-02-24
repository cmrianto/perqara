package repo

func (r *perqaraRepo) ForeignKeySwitch(in int) {
	stringQuery := ""
	switch in {
	case 1:
		stringQuery = "SET FOREIGN_KEY_CHECKS = 1"
	case 0:
		stringQuery = "SET FOREIGN_KEY_CHECKS = 0"
	default:
		return
	}

	r.db.Exec(stringQuery)
}
