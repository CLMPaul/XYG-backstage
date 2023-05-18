package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	dialector "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMySQL() *gorm.DB {
	var databaseName string
	if database := os.Getenv("MYSQL_DATABASE"); database != "" {
		databaseName = database
	} else {
		databaseName = defaultDatabaseName
	}

	cfg := mysql.NewConfig()
	cfg.Addr = os.Getenv("MYSQL_ADDR")
	cfg.User = os.Getenv("MYSQL_USER")
	cfg.Passwd = os.Getenv("MYSQL_PASSWORD")
	cfg.Params = map[string]string{
		"charset": "utf8mb4,utf8",
	}
	cfg.ParseTime = true
	cfg.Loc = time.Local

	connector, err := mysql.NewConnector(cfg)
	if err != nil {
		logrus.Fatalf("Cannot connect to mysql: %v", err)
	}
	pool := sql.OpenDB(connector)
	_, err = pool.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET 'utf8mb4'", databaseName))
	if err != nil {
		logrus.Fatalf("Failed creating mysql database: %v", err)
	}

	cfg.DBName = databaseName
	connector, err = mysql.NewConnector(cfg)
	if err != nil {
		logrus.Fatalf("Cannot connect to mysql: %v", err)
	}
	pool = sql.OpenDB(connector)
	logrus.Infof("connecting to mysql: %s", cfg.Addr)

	db, err := gorm.Open(
		&dialector.Dialector{Config: &dialector.Config{Conn: pool}},
	)
	if err != nil {
		logrus.Fatalln(err)
	}

	return db
}
