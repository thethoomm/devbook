package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (u *User) Prepare(stage string) error {
	if err := u.validate(stage); err != nil {
		return err
	}

	if err := u.format(stage); err != nil {
		return err
	}

	return nil
}

func (u *User) validate(stage string) error {
	if u.Name == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if u.Username == "" {
		return errors.New("o apelido é obrigatório e não pode estar em branco")
	}

	if u.Email == "" {
		return errors.New("o email é obrigatório e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("o email é inválido")
	}

	if stage == "signup" && u.Password == "" {
		return errors.New("o password é obrigatório e não pode estar em branco")
	}

	return nil
}

func (u *User) format(stage string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)

	if stage == "signup" {
		hashedPassword, err := security.Hash(u.Password)
		if err != nil {
			return err
		}

		u.Password = string(hashedPassword)
	}

	return nil
}
