package users

import (
	"database/sql"
)

func addUser(db *sql.DB, payload CreateUserPayload) error {
	_, err := db.Exec("INSERT INTO users(name, email) VALUES (?, ?)", payload.Name, payload.Email)
	if err != nil {
		return err
	}

	return nil
}

func getUserById(db *sql.DB, id int) (*User, error) {
	user := new(User)
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func getAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * from users LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
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
