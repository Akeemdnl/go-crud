package users

import (
	"database/sql"
	"fmt"
)

func getUserBy(db *sql.DB, id int, column string) (*User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE %s = ?", column)
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	user := new(User)
	for rows.Next() {
		user, err = scanUserIntoRow(rows, user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func scanUserIntoRow(rows *sql.Rows, user *User) (*User, error) {
	err := rows.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
