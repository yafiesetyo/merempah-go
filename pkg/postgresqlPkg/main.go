package postgresqlpkg

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Connection ...
type Connection struct {
	Host                    string
	DbName                  string
	User                    string
	Password                string
	Port                    int
	Location                *time.Location
	SslMode                 string
	SslCert                 string
	SslKey                  string
	SslRootCert             string
	DBMaxConnection         int
	DBMAxIdleConnection     int
	DBMaxLifeTimeConnection int
}

// Connect ...
func (c Connection) Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=%s&TimeZone=UTC", c.User, c.Password, c.Host, c.DbName, c.SslMode,
	)

	if c.SslMode == "require" {
		connStr = fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=UTC&sslcert=%s&sslkey=%s&sslrootcert=%s",
			c.User, c.Password, c.Host, c.Port, c.DbName, c.SslMode, c.SslCert, c.SslKey, c.SslRootCert,
		)
	}
	db, err := sql.Open("postgres", connStr)
	db.SetMaxOpenConns(c.DBMaxConnection)
	db.SetMaxIdleConns(c.DBMAxIdleConnection)
	db.SetConnMaxLifetime(time.Duration(c.DBMaxLifeTimeConnection) * time.Second)

	return db, err
}
