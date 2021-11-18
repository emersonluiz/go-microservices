package repository

import (
	"context"
	"database/sql"

	"github.com/emersonluiz/go-user/logger"
	"github.com/emersonluiz/go-user/models"
	"github.com/emersonluiz/go-user/service"
)

func CreateUser(db *sql.DB, user *models.User) (*models.User, error) {
	var id int

	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	pass := service.Encoder(user.Password)

	error := db.QueryRow("INSERT INTO users (name, email, password) VALUES (@name, @email, @password); SELECT SCOPE_IDENTITY();", sql.Named("name", user.Name), sql.Named("email", user.Email), sql.Named("password", pass)).Scan(&id)
	if error != nil {
		logger.SetLog(error.Error())
		return nil, error
	}
	user.Id = id
	user.Password = ""
	logger.SetLog("User create with success")
	return user, nil
}

func FindAllUser(db *sql.DB) ([]models.User, error) {
	var id int
	var name, email string

	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		logger.SetLog(err.Error())
		return nil, err
	}

	rows, err := db.Query(`SELECT id, name, email FROM users;`)
	if err != nil {
		logger.SetLog(err.Error())
		return nil, err
	}
	defer rows.Close()
	defer db.Close()

	var users []models.User
	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			logger.SetLog(err.Error())
			return nil, err
		}

		users = append(users, models.User{Id: id, Name: name, Email: email})
	}
	if err := rows.Err(); err != nil {
		logger.SetLog(err.Error())
		return nil, err
	}
	return users, nil
}

func DeleteUser(db *sql.DB, id int) error {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		logger.SetLog(err.Error())
		return err
	}

	defer db.Close()

	res, err := db.Exec("DELETE FROM users WHERE id = @id", sql.Named("id", id))

	if err != nil {
		logger.SetLog(err.Error())
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		logger.SetLog(err.Error())
		return err
	}

	if rowsAffected > 0 {
		return nil
	}

	return nil
}

func FindOneUser(db *sql.DB, id int) (models.User, error) {
	var user models.User

	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		logger.SetLog(err.Error())
		return user, err
	}

	defer db.Close()

	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = @id", sql.Named("id", id))
	err = row.Scan(&user.Id, &user.Name, &user.Email)

	switch err {
	case sql.ErrNoRows:
		logger.SetLog("No rows were returned!")
		return user, err
	default:
		return user, err
	}
}

func UpdateUser(db *sql.DB, id int, user *models.User) error {

	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		logger.SetLog(err.Error())
		return err
	}

	defer db.Close()

	pass := service.Encoder(user.Password)

	res, err := db.Exec("UPDATE users SET name = @name, email = @email, password = @password);", sql.Named("name", user.Name), sql.Named("email", user.Email), sql.Named("password", pass))

	if err != nil {
		logger.SetLog(err.Error())
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		logger.SetLog(err.Error())
		return err
	}

	if rowsAffected > 0 {
		return nil
	}

	return nil
}
