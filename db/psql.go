// db/psql.go

package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// PSQL represents a PostgreSQL connection
type PSQL struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	db       *sql.DB
}

// Table represents a PostgreSQL table with its schema
type Table struct {
	Schema string
	Name   string
}

// NewPSQL creates a new PSQL instance with the provided connection parameters
func NewPSQL(host, port, username, password, database string) *PSQL {
	return &PSQL{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
	}
}

// Connect establishes a connection to the PostgreSQL database
func (p *PSQL) Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.Host, p.Port, p.Username, p.Password, p.Database,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	p.db = db
	return db, nil
}

// GetTables retrieves all user-defined tables along with their schemas
func (p *PSQL) GetTables() ([]Table, error) {
	if p.db == nil {
		return nil, fmt.Errorf("database connection is not established")
	}

	query := `
		SELECT schemaname, tablename
		FROM pg_catalog.pg_tables
		WHERE schemaname NOT IN ('pg_catalog', 'information_schema')
		ORDER BY schemaname, tablename
	`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GetTables query: %w", err)
	}
	defer rows.Close()

	var tables []Table
	for rows.Next() {
		var table Table
		if err := rows.Scan(&table.Schema, &table.Name); err != nil {
			return nil, fmt.Errorf("failed to scan table row: %w", err)
		}
		p.Grant(table)
		tables = append(tables, table)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over table rows: %w", err)
	}

	return tables, nil
}

// GetColumns retrieves all column names for a given table and schema
func (p *PSQL) GetColumns(tableName, schemaName string) ([]string, error) {
	if p.db == nil {
		return nil, fmt.Errorf("database connection is not established")
	}

	query := `
		SELECT column_name
		FROM information_schema.columns
		WHERE table_name = $1
		  AND table_schema = $2
		ORDER BY ordinal_position
	`

	rows, err := p.db.Query(query, tableName, schemaName)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GetColumns query: %w", err)
	}
	defer rows.Close()

	var columns []string
	for rows.Next() {
		var columnName string
		if err := rows.Scan(&columnName); err != nil {
			return nil, fmt.Errorf("failed to scan column name: %w", err)
		}
		columns = append(columns, columnName)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over column rows: %w", err)
	}

	return columns, nil
}

func (p *PSQL) Grant(table Table) error {
	_, err := p.db.Exec("GRANT ALL ON TABLE " + table.Schema + "." + table.Name + " TO " + p.Username)
	if err != nil {
		return fmt.Errorf("failed to grant privileges: %w", err)
	}
	return nil
}
