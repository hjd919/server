package pagination

func GetOffsetByPage(page int64, limit int64) (offset int64) {
	return (page - 1) * limit
}
