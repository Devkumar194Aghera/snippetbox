package mysql

import (
	"database/sql"
	"strings"

	"goModule/snippetbox/pkg/models"

	"github.com/go-sql-driver/mysql" // New import
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name string, email string, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES(?, ?, ?, UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "users.users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}

	}

	return err

}

func (m *UserModel) Get(id int) (*models.User, error) {

	return nil, nil

}

func (m *UserModel) Authneticate(email string, password string) (int, error) {

	var id int
	var hashedPassword []byte
	row := m.DB.QueryRow("select id,hashed_password FROM users WHERE email = ?", email)
	err := row.Scan(&id, &hashedPassword)

	if err == sql.ErrNoRows {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	return id, nil

}