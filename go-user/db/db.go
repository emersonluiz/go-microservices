package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/emersonluiz/go-user/logger"
)

var db *sql.DB

var server = os.Getenv("DATABASE_HOST")
var port = os.Getenv("DATABASE_PORT")
var user = os.Getenv("DATABASE_USER")
var password = os.Getenv("DATABASE_PASS")
var database = os.Getenv("DATABASE_NAME")

func Connect() *sql.DB {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)

	var err error

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		logger.SetLog(("Error creating connection pool: " + err.Error()))
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		logger.SetLog(err.Error())
	}
	logger.SetLog("DB Connected!\n")
	return db
}
