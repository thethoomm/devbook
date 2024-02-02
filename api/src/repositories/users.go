package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func (u Users) Create(user models.User) (uint64, error) {

	stmt, err := u.db.Prepare("INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	response, err := stmt.Exec(user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := response.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (u Users) Find(nameOrUsername string) (*[]models.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername)

	rows, err := u.db.Query(
		"SELECT id, name, username, email, createdAt FROM users WHERE name LIKE ? or username LIKE ?",
		nameOrUsername, nameOrUsername,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}

func (u Users) FindById(id uint64) (*models.User, error) {
	rows, err := u.db.Query("SELECT id, name, username, email, createdAt FROM users WHERE id = ?", id)
	if err != nil {
		return &models.User{}, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return &models.User{}, err
		}
	}

	return &user, nil
}

func (u Users) FindByEmail(email string) (*models.User, error) {
	rows, err := u.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return &models.User{}, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Password); err != nil {
			return &models.User{}, err
		}
	}

	return &user, nil
}

func (u Users) Update(id uint64, user models.User) error {
	stmt, err := u.db.Prepare("UPDATE users SET name = ?, username = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Username, user.Email, id)
	if err != nil {
		return err
	}

	return nil
}

func (u Users) Delete(id uint64) error {
	stmt, err := u.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (u Users) Follow(userID, followerID uint64) error {
	stmt, err := u.db.Prepare("INSERT IGNORE INTO followers (userID, followerID) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (u Users) Unfollow(userID, followerID uint64) error {
	stmt, err := u.db.Prepare("DELETE FROM followers WHERE userID = ? AND followerID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (u Users) FindFollowers(userID uint64) (*[]models.User, error) {
	rows, err := u.db.Query("SELECT u.id, u.name, u.username, u.email, u.createdAt FROM users u INNER JOIN followers f ON u.id = f.followerID WHERE f.userID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		followers = append(followers, user)
	}

	return &followers, nil
}

func (u Users) FindFollowing(userID uint64) (*[]models.User, error) {
	rows, err := u.db.Query("SELECT u.id, u.name, u.username, u.email, u.createdAt FROM users u INNER JOIN followers f ON u.id = f.userID WHERE f.followerID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, nil
}

func (u Users) FindPassword(userID uint64) (string, error) {
	rows, err := u.db.Query("SELECT password FROM users WHERE id = ?", userID)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err := rows.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (u Users) UpdatePassword(userID uint64, passwordHashed string) error {
	stmt, err := u.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(passwordHashed, userID); err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db: db}
}
