package users

import (
	"database/sql"
)

func getUserById(db *sql.DB, id int) (*User, error) {
	user := new(User)
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func updateUser(db *sql.DB, payload *UpdateUserPayload, id int) (*User, error) {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := db.Exec(query, payload.Name, payload.Email, id)
	if err != nil {
		return nil, err
	}

	user := new(User)
	query = "SELECT * from users WHERE id = ?"
	if err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return user, nil
}

// func scanUserIntoRow(rows *sql.Rows, user *User) (*User, error) {
// 	err := rows.Scan(
// 		&user.ID,
// 		&user.Name,
// 		&user.Email,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }
