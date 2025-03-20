package dbsql

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// SQLiteConfig defines config for sqlite driver
type SQLiteConfig struct {
	// DBPath database location path.
	//
	// Optional. Default root directory of runner.
	DBPath string

	// TimeZone that want to use as a default timezone in Database
	//
	// Optional. Default is "Asia/Jakarta"
	TimeZone string

	cfg Config
}

// Driver name identity
const SQLite Dialect = "sqlite"

// PgOptions create new options of postgres driver
func (d Dialect) SqliteOptions(c SQLiteConfig) *SQLiteConfig {
	return &SQLiteConfig{
		DBPath:   c.DBPath,
		TimeZone: c.TimeZone,
	}
}

// Return postgres dialector
func (c *SQLiteConfig) dialector() gorm.Dialector {
	return sqlite.Open(c.DBPath)
}
