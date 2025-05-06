package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"imobiliaria/internal/models"
	repositories "imobiliaria/internal/repositories/database"

	// Importing MySQL driver for its side effects
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type MySQL struct {
	Database *sql.DB
	// Cache    *redis.Redis ??
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewRepository(config *Config) (*MySQL, error) {
	logrus.Trace("Creating MySQL repository")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Database))
	if err != nil {
		logrus.WithError(err).Error("Error opening database")

		return nil, err
	}

	if err := db.Ping(); err != nil {
		logrus.WithError(err).Error("Error pinging database")

		return nil, err
	}

	logrus.Infoln("Connected to MySQL")

	MySQL := &MySQL{
		Database: db,
	}

	if err = MySQL.Migrate(); err != nil {
		logrus.WithError(err).Error("Error migrating database")

		return nil, err
	}

	return MySQL, nil
}

func (m *MySQL) Migrate() error {
	/// TODO: Better migration system
	ctx := context.TODO()

	_, err := m.Database.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), email VARCHAR(255) UNIQUE, password VARCHAR(255), age INT)")
	if err != nil {
		return err
	}

	_, err = m.Database.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS houses (id INT AUTO_INCREMENT PRIMARY KEY, street VARCHAR(255), number VARCHAR(255), city VARCHAR(255), state VARCHAR(255), zip_code VARCHAR(255), price FLOAT, owner_id INT)")
	if err != nil {
		return err
	}

	return nil
}

// CreateUser implements repositories.Repositories.
func (m *MySQL) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	logrus.Trace("Executing CreateUser", user)
	logrus.Trace(user.PasswordHash)

	password, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithError(err).Trace("Error hashing password")

		return nil, err
	}

	// TODO: check if email already exists

	result, err := m.Database.ExecContext(ctx, "INSERT INTO users (name, email, age, password) VALUES (?, ?, ?, ?)", user.Name, user.Email, user.Age, string(password))
	logrus.Trace("User created?", result)

	if err != nil {
		logrus.WithError(err).Trace("Error exec creating user")

		return nil, err
	}

	rows, err := m.Database.QueryContext(ctx, "SELECT * FROM users WHERE email = ?", user.Email)
	if err != nil {
		logrus.WithError(err).Trace("Error creating user")

		return nil, err
	}

	defer rows.Close()

	var u models.User
	logrus.Trace("User created with success? Maybe")
	for rows.Next() {
		logrus.Trace("User on next")

		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.Age)
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
	rows, err := m.Database.QueryContext(ctx, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		logrus.WithError(err).Trace("Error getting user")

		return nil, err
	}

	defer rows.Close()

	var u models.User
	logrus.Trace("Getting user enter row")
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.Age)
		logrus.Trace("Getting user on next")
		if err != nil {
			logrus.WithError(err).Trace("Error on next")

			return nil, err
		}
	}

	u.PasswordHash = ""

	return &u, nil
}

func (m *MySQL) Login(ctx context.Context, email, password string) (*models.User, error) {
	logrus.Trace("Executing Login")

	if email == "" || password == "" {
		logrus.Trace("email or password is empty at login")

		return nil, fmt.Errorf("email or password is empty")
	}

	row := m.Database.QueryRowContext(ctx, "SELECT * FROM users WHERE email = ?", email)

	if row.Err() != nil {
		logrus.WithError(row.Err()).Trace("Error getting user")

		return nil, row.Err()
	}

	var u models.User
	logrus.Trace("Getting user enter row")

	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.Age); err != nil {
		logrus.WithError(err).Trace("Error on next")

		return nil, err
	}

	logrus.Info([]byte(u.PasswordHash), []byte(password))
	logrus.Info(u.PasswordHash, password)

	// $2a$10$94tS/w/5SIFShxTa5B4crO/CI.6ueQpSlIOvGmDgYwFXjuNMLAZJy

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		logrus.WithError(err).Trace("Error comparing password")

		return nil, err
	}

	return &u, nil
}

// GetHouse Gets a house by ID
//
// id must be greater than 0
func (m *MySQL) GetHouse(ctx context.Context, id uint) (*models.House, error) {
	logrus.Trace("Getting House")
	rows, err := m.Database.QueryContext(ctx, "SELECT * FROM houses WHERE id = ?", id)
	if err != nil {
		logrus.WithError(err).Trace("Error getting house")

		return nil, err
	}

	defer rows.Close()

	var house models.House
	logrus.Trace("Getting houses enter row")
	for rows.Next() {
		// validar os campos aqui também
		err := rows.Scan(&house.ID, &house.Street, &house.Number, &house.City, &house.State, &house.ZipCode, &house.Price, &house.OwnerID)
		if err != nil {
			logrus.WithError(err).Trace("Error on next")

			return nil, err
		}
	}

	return &house, nil
}

func (m *MySQL) CreateHouse(ctx context.Context, house *models.House) (*models.House, error) {
	logrus.Trace("Executing CreateHouse", house.Number)
	result, err := m.Database.ExecContext(ctx, "INSERT INTO houses (street, number, city, state, zip_code, price, owner_id) VALUES (?, ?, ?, ?, ?, ?, ?)", house.Street, house.Number, house.City, house.State, house.ZipCode, house.Price, *house.OwnerID)
	if err != nil {
		logrus.WithError(err).Trace("Error exec creating house")

		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logrus.WithError(err).Trace("Error getting last insert id")

		return nil, err
	}

	// Too much, mas não tenho unique aqui
	rows, err := m.Database.QueryContext(ctx, "SELECT * FROM houses WHERE id = ?", id)
	if err != nil {
		logrus.WithError(err).Trace("Error creating house")

		return nil, err
	}

	defer rows.Close()

	var returnHouse models.House
	for rows.Next() {
		logrus.Trace("House on next")

		err := rows.Scan(
			&returnHouse.ID,
			&returnHouse.Street,
			&returnHouse.Number,
			&returnHouse.City,
			&returnHouse.State,
			&returnHouse.ZipCode,
			&returnHouse.Price,
			&returnHouse.OwnerID,
		)
		if err != nil {
			logrus.WithError(err).Trace("Error on next")

			return nil, err
		}
	}

	return &returnHouse, nil
}

func (m *MySQL) GetHouses(ctx context.Context, limit uint, offset uint) ([]*models.House, error) {
	logrus.Trace("Getting Houses")
	/// TODO: sorted funciona?
	rows, err := m.Database.QueryContext(ctx, "SELECT * FROM houses SORTED BY `ID` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		logrus.WithError(err).Trace("Error getting houses")

		return nil, err
	}
	defer rows.Close()

	/// Parece que posso fazer isso de um jeito mais limpinho, criar objeto e dar append é meio feio, mas fácil
	var houses []*models.House
	for rows.Next() {
		var house models.House
		err := rows.Scan(&house.ID, &house.Street, &house.Number, &house.City, &house.State, &house.ZipCode, &house.Price, &house.OwnerID)
		if err != nil {
			logrus.WithError(err).Trace("Error on next")

			return nil, err
		}
		houses = append(houses, &house)
	}

	logrus.Trace("Getting houses with success")

	return houses, nil
}

func (m *MySQL) UpdateHouse(ctx context.Context, house *models.House) (*models.House, error) {
	logrus.Trace("Updating House")
	_, err := m.Database.ExecContext(ctx, "UPDATE houses SET street = ?, number = ?, city = ?, state = ?, zip_code = ?, price = ?, owner_id = ? WHERE id = ?", house.Street, house.Number, house.City, house.State, house.ZipCode, house.Price, house.OwnerID, house.ID)
	if err != nil {
		logrus.WithError(err).Trace("Error updating house")

		return nil, err
	}

	return house, nil
}

func (m *MySQL) DeleteHouse(ctx context.Context, id uint) error {
	logrus.Trace("Deleting House")
	_, err := m.Database.ExecContext(ctx, "DELETE FROM houses WHERE id = ?", id)
	if err != nil {
		logrus.WithError(err).Trace("Error deleting house")

		return err
	}

	return nil
}

func (m *MySQL) GetHousesByUserID(ctx context.Context, id uint, limit uint, offset uint) ([]*models.House, error) {
	logrus.Trace("Getting Houses by User ID")

	/// TODO: sorted funciona?
	rows, err := m.Database.QueryContext(ctx, "SELECT * FROM houses WHERE owner_id = ? SORTED BY `ID` LIMIT ? OFFSET ?", id, limit, offset)
	if err != nil {
		logrus.WithError(err).Trace("Error getting houses by user id")

		return nil, err
	}
	defer rows.Close()

	var houses []*models.House
	for rows.Next() {
		var house models.House
		err := rows.Scan(&house.ID, &house.Street, &house.Number, &house.City, &house.State, &house.ZipCode, &house.Price, &house.OwnerID)
		if err != nil {
			logrus.WithError(err).Trace("Error on next")

			return nil, err
		}
		houses = append(houses, &house)
	}

	logrus.Trace("Getting houses by user id with success")

	return houses, nil
}

var _ repositories.Repositories = &MySQL{}
