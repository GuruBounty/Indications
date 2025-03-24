package repository

import (
	"context"
	"indication/internal/domain"

	"github.com/jmoiron/sqlx"
	//"golang.org/x/crypto/bcrypt"
)

type Users struct {
	dbx *sqlx.DB
}

func NewUsers(db *sqlx.DB) *Users {
	return &Users{dbx: db}
}

// func (u *Users) Create(ctx context.Context, user domain.User) error {
// 	//_, err := u.db.Exec("INSERT INTO users(name, password, registerAt) values ($1, $2, $3)", user.Name, user.Password, user.RegisterAt)
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	query := "INSERT INTO users (username, password, registered_at ) VALUES ($1, $2, $3)"
// 	_, err = u.dbx.ExecContext(ctx, query, user.Name, string(hashedPassword), user.RegisterAt)

// 	// stm, err := u.dbx.Prepare("INSERT INTO users(username, password, registered_at) VALUES ($1, $2, $3)")
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// defer stm.Close()
// 	// _, err = stm.Exec(user.Email, user.Password, user.RegisterAt)
// 	return err
// }

func (u *Users) GetDB() *sqlx.DB {
	return u.dbx
}

func (u *Users) GetByCredentials(ctx context.Context, email string, password string) (domain.User, error) {
	var user domain.User
	query := "SELECT id, email FROM users WHERE email= $1 AND password= $2"
	//err := u.dbx.QueryRow("SELECT id, email FROM users WHERE email= $1 AND password= $2", email, password).Scan(&user.UserID, &user.Email)
	err := u.dbx.QueryRowxContext(ctx, query, email, password).Scan(&user.UserID, &user.Email)
	if err != nil {
		return user, err
	}
	return user, err
}
