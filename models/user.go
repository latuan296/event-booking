package models

import (
	"errors"

	"exmpale.com/event-booking/db"
	"exmpale.com/event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	newUser, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := newUser.LastInsertId()

	u.ID = userId

	return err
}

func (u *User) ValidateCredential() error {

	query := "SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credential is not valid")
	}

	passwordIsValid := utils.CheckPassword(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credential is not valid")
	}

	return nil

}
