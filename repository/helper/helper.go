package helper

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Suite is struct declarete variable use at setup
type Suite struct {
	Conn *sql.DB
	Db   *gorm.DB
	Mock sqlmock.Sqlmock
}

// Setup is helper for init psql connection interface in unit test
func Setup() *Suite {
	conn, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 conn,
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &Suite{
		Conn: conn,
		Db:   db,
		Mock: mock,
	}
}
