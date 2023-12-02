package store

func CountActiveRequests() int {
	query := `
        SELECT COUNT(*) FROM Request WHERE status = 'active'
    `

	var count int
	Db_Conn.Get(&count, query)

	return count
}
