package dbsql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PgSQLConfig defines config for postgres driver
type PgSQLConfig struct {
	// SSLMode mode of ssl that using to connect to Database
	//
	// Optional. Default is true
	SSLMode bool

	// TimeZone that want to use as a default timezone in Database
	//
	// Optional. Default is "Asia/Jakarta"
	TimeZone string

	cfg Config
}

// Driver name identity
const Postgres Dialect = "postgres"

// PgOptions create new options of postgres driver
func (d Dialect) PgOptions(c PgSQLConfig) *PgSQLConfig {
	return &PgSQLConfig{
		SSLMode:  c.SSLMode,
		TimeZone: c.TimeZone,
	}
}

// Return postgres dialector
func (c *PgSQLConfig) dialector() gorm.Dialector {
	return postgres.Open(resolveDSN(c))
}

// Return postgres dsn
func resolveDSN(c *PgSQLConfig) string {
	// Variable
	var (
		sslMode  string
		timezone = "Asia/Jakarta"
	)

	// Determine option
	if c.SSLMode {
		sslMode = "enable"
	} else {
		sslMode = "disable"
	}
	if len(c.TimeZone) > 0 && c.TimeZone != " " {
		timezone = c.TimeZone
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.cfg.Host,
		c.cfg.User,
		c.cfg.Password,
		c.cfg.DBName,
		c.cfg.Port,
		sslMode,
		timezone,
	)
}
