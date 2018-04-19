package factory

import "database/sql"

//The first implementation.
type PostgreSQLDataStore struct {
	DSN string
	DB  sql.DB
}

func (pds *PostgreSQLDataStore) Name() string {
	return "PostgreSQLDataStore"
}

func (pds *PostgreSQLDataStore) FindUserNameById(id int64) (string, error) {
	var username string
	rows, _ := pds.DB.Query("SELECT username FROM users WHERE id=", id)
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			if err == sql.ErrNoRows {
				return "", UserNotFoundError
			}
			return "", err
		}
	}
	return username, nil
}
