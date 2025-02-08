package user

import (
	"awesomeProject/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type IUserRepository interface {
	//WithTx(tx *sql.Tx) IUserRepository
	CreateUser(ctx context.Context, user *User) error
	GetUsers(ctx context.Context) ([]*User, error)
	DeleteUser(ctx context.Context, user *User) error
}

type Repository struct {
	db *sql.DB
	//tx *sql.Tx
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

//func (r *Repository) WithTx(tx *sql.Tx) IUserRepository {
//	return &Repository{db: r.db, tx: tx}
//}

func (r *Repository) CreateUser(ctx context.Context, user *User) error {
	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	query := `INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3)`
	result, err := r.db.ExecContext(ctx, query, user.Username, user.Password, user.CreatedAt)
	helper.PanicIfError(err)

	rowAffected, err := result.RowsAffected()
	helper.PanicIfError(err)

	if rowAffected == 0 {
		return errors.New("user already exists")
	}

	return nil
}

func (r *Repository) GetUsers(ctx context.Context) ([]*User, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, username, created_at FROM users`) // check err
	helper.PanicIfError(err)
	defer rows.Close()

	var users []*User

	//var users []user
	for rows.Next() {
		var u = &User{}
		err = rows.Scan(&u.ID, &u.Username, &u.CreatedAt)
		helper.PanicIfError(err)
		users = append(users, u)
	}
	err = rows.Err() // check err
	helper.PanicIfError(err)

	return users, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id) // check err
	helper.PanicIfError(err)

	rowsAffected, err := result.RowsAffected()
	helper.PanicIfError(err)

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *Repository) CreateTable() {
	query := `
	CREATE TABLE users (
	   id SERIAL PRIMARY KEY,
	   username TEXT NOT NULL,
	   password TEXT NOT NULL,
	   created_at TIMESTAMP
	);`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	_, err := r.db.Exec(query)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
}
