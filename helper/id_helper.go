package helper

func EmptyId(id int64) *int64 {
	if id == 0 {
		return nil
	}
	return &id
}
