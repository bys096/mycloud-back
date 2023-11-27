package repository

import (
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"mycloud/errors"
	"time"
)

type User struct {
	Id        uint64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserModel() *User {
	return &User{}
}

func (u *User) Save(user *User) (*User, error) {
	stmt, err := Db.Prepare("INSERT INTO users(name, email, password, created_at, updated_at) VALUES($1, $2, $3, $4, $5)")
	checkError(err)
	defer stmt.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	_, err = stmt.Exec(user.Name, user.Email, hashedPassword, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return user, errors.Errorf(errors.UserExist, "user is already exist: %s", err)

	}

	return user, nil
}

func (u *User) Login(user *User) {

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
