package store

func DeleteStaff(username string) error {

    query := `
        DELETE FROM 
            User 
        WHERE 
            username = ?
    `

    _, err := Db_Conn.Exec(query, username)
    if err != nil {
        return err
    }

    return nil
}
