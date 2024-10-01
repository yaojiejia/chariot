package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PSQL struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	db       *sql.DB
}

func NewPSQL(host, port, username, password, database string) *PSQL {
	return &PSQL{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
	}
}

func (p *PSQL) Connect() (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.Username, p.Password, p.Database)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	p.db = db
	return db, nil

}

func (p *PSQL) Read() string {
	db := p.db
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return err.Error()
	}
	defer rows.Close()

}
