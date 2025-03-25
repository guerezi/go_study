package mysql

import (
	"context"
	"database/sql"

	"imobiliaria/internal/models"
	"imobiliaria/internal/repositories"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type MySQL struct {
	Database *sql.DB
}

func NewRepository() *MySQL {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/database")

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	logrus.Infoln("Connected to MySQL")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), email VARCHAR(255), age INT)")

	if err != nil {
		panic(err)
	}

	return &MySQL{
		Database: db,
	}
}

// CreateUser implements repositories.Repositories.
func (m *MySQL) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	logrus.Trace("Executing CreateUser", user)
	result, err := m.Database.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)", user.Name, user.Email, user.Age)

	logrus.Trace("User created?", result)

	if err != nil {
		logrus.WithError(err).Trace("Error exec creating user")

		return nil, err
	}

	logrus.Trace("User created?")
	rows, err := m.Database.Query("SELECT * FROM users WHERE email = ?", user.Email)

	if err != nil {
		logrus.WithError(err).Trace("Error creating user")

		return nil, err
	}

	defer rows.Close()

	var u models.User
	logrus.Trace("User created with success? Maybe")
	for rows.Next() {
		logrus.Trace("User on next")

		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		if err != nil {
			logrus.WithError(err).Trace("Error on next")

			return nil, err
		}
	}

	logrus.Trace("User created with success", u)

	return &u, nil
}

// GetUser implements repositories.Repositories.
func (m *MySQL) GetUser(ctx context.Context, id int) (*models.User, error) {
	logrus.Trace("Getting user")
	rows, err := m.Database.Query("SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		logrus.WithError(err).Trace("Error getting user")

		return nil, err
	}

	defer rows.Close()

	var u models.User
	logrus.Trace("Getting user enter row")
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		logrus.Trace("Getting user on next")
		if err != nil {
			logrus.WithError(err).Trace("Error on next")

			return nil, err
		}
	}

	return &u, nil
}

var _ repositories.Repositories = &MySQL{}
