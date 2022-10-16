package repository_test

import (
	"final-project/models"
	psqlrepo "final-project/repository"
	"final-project/repository/helper"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/assert"
)

func Test_GetUser(t *testing.T) {
	s := helper.Setup()
	defer s.Conn.Close()

	repo := psqlrepo.UserRepo{Conn: s.Db}

	t.Run("success", func(t *testing.T) {
		query := `SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`

		dataUser := models.User{
			Username: "deni",
			Email:    "deni@gmail.com",
			Password: "123456",
			Age:      25,
		}
		id := uint(1)
		rows := sqlmock.
			NewRows([]string{"id", "username", "password", "age"}).
			AddRow(id, dataUser.Username, dataUser.Password, dataUser.Age)

		s.Mock.
			ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(dataUser.Email).
			WillReturnRows(rows)

		err := repo.GetUserByEmail(&dataUser)

		assert.Equal(t, nil, s.Mock.ExpectationsWereMet())
		assert.Equal(t, nil, err)
		assert.Nil(t, deep.Equal(dataUser.ID, id))
	})

	t.Run("error connection", func(t *testing.T) {
		s.Conn.Close()
		user := models.User{Email: "deni@gmail.com"}
		err := repo.GetUserByEmail(&user)

		assert.Error(t, err)
	})
}
