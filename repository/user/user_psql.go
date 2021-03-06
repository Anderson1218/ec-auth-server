package userRepository

import (
	"database/sql"
	"github.com/Anderson1218/ec-auth-server/models"
	"log"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (u UserRepository) Signup(db *sql.DB, user models.User) (models.User, error) {
	stmt := "insert into users (email, password, name) values($1, $2, $3) RETURNING id;"
	err := db.QueryRow(stmt, user.Email, user.Password, user.Name).Scan(&user.ID)

	if err != nil {
		return user, err
	}

	user.Password = ""
	return user, nil
}

func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("select * from users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserRepository) GetUserProfile(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("select id, name from users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Name)

	if err != nil {
		return user, err
	}

	return user, nil
}
