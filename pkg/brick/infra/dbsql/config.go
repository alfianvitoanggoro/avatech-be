package dbsql

import (
	"encoding/json"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

// Config defines the config for App Database
type Config struct {
	// Debugging mode
	Debug bool `json:"debug" yml:"debug" toml:"debug" mapstructure:"debug"`

	// Host name where the Database is hosted
	//
	// Required. Default is "127.0.0.1"
	Host string `json:"host" yml:"host" toml:"host" mapstructure:"host"`

	// Port number of Database connection
	//
	// Required. Default 5432
	Port int `json:"port" yml:"port" toml:"port" mapstructure:"port"`

	// User name of Database connection
	//
	// Required. Default is "test"
	User string `json:"user" yml:"user" toml:"user" mapstructure:"user"`

	// Password of Database conenction
	//
	// Required. Default is "test"
	Password string `json:"password" yml:"password" toml:"password" mapstructure:"password"`

	// DBName name of Database that want to connect
	//
	// Required. Default is "test"
	DBName string `json:"db" yml:"db" toml:"db" mapstructure:"db"`

	// Dialect is varian or type of database query language
	//
	// Required. Default is "postgres"
	Dialect Dialect `json:"dialect" yml:"dialect" toml:"dialect" mapstructure:"dialect"`

	// Options is config of gorm
	//
	// Optional. Default is nil
	Options gorm.Option `json:"options" yml:"options" toml:"options" mapstructure:"options"`

	TimeZone string `json:"time_zone" yml:"time_zone" toml:"time_zone" mapstructure:"time_zone"`

	OrderOpen      int  `json:"order_open" yml:"order_open" toml:"order_open" mapstructure:"order_open"`
	OrderClose     int  `json:"order_close" yml:"order_close" toml:"order_close" mapstructure:"order_close"`
	MustCloseFirst bool `json:"must_close_first" yml:"must_close_first" toml:"must_close_first" mapstructure:"must_close_first"`
}

type DialectOptions func(d Dialect, c interface{}) interface{}

type Dialect string

type connectionConfig struct {
	dialector gorm.Dialector
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Host:     "127.0.0.1",
	Port:     5432,
	User:     "testing",
	Password: "testing123",
	DBName:   "testing",
	Dialect:  "postgres",
	Options:  &gorm.Config{},
}

// GenConfig generate dbsql configs.
func GenConfig(cfgDialect string) (ret string, err error) {
	type tmp struct {
		DBSQL Config `json:"db_sql" yml:"db_sql" toml:"db_sql" mapstructure:"db_sql"`
	}

	var b []byte
	switch cfgDialect {
	case "json":
		b, err = json.Marshal(tmp{})

	case "yaml", "yml":
		b, err = yaml.Marshal(tmp{})

	default:
		b, err = toml.Marshal(tmp{})
	}

	if b != nil {
		ret = string(b)
	}
	return
}

// Helper function to set default config
func configDefault(config Config, driverOptions ...interface{}) *connectionConfig {
	// Overide default config
	cfg := config

	// Validating
	if cfg.Host == "" {
		cfg.Host = ConfigDefault.Host
	}
	if cfg.Port == 0 {
		cfg.Port = ConfigDefault.Port
	}
	if cfg.User == "" {
		cfg.User = ConfigDefault.User
	}
	if cfg.Password == "" {
		cfg.Password = ConfigDefault.Password
	}
	if cfg.DBName == "" {
		cfg.DBName = ConfigDefault.DBName
	}
	if cfg.Dialect == "" {
		cfg.Dialect = ConfigDefault.Dialect
	}
	if cfg.Options == nil {
		cfg.Options = config.Options
	}

	return finalConfigBasedOnDriver(cfg, driverOptions...)
}

// Helper function to set final "connection-use" config based on choosen driver
func finalConfigBasedOnDriver(config Config, driverOptions ...interface{}) *connectionConfig {
	switch config.Dialect {
	// PostgreSQL Driver
	case Postgres:
		pgc := &PgSQLConfig{cfg: config}
		if len(driverOptions) > 0 {
			dop, ok := driverOptions[0].(*PgSQLConfig)
			if !ok {
				panic("Given option is not belongs to postgres")
			}
			pgc.SSLMode = dop.SSLMode
			pgc.TimeZone = dop.TimeZone
		}
		return &connectionConfig{dialector: pgc.dialector()}

	// SQLite Driver
	case SQLite:
		sqc := &SQLiteConfig{cfg: config}
		if len(driverOptions) > 0 {
			dop, ok := driverOptions[0].(*SQLiteConfig)
			if !ok {
				panic("Given option is not belongs to sqlite")
			}
			sqc.DBPath = dop.DBPath
			sqc.TimeZone = dop.TimeZone
		} else {
			sqc.DBPath = config.DBName
		}
		return &connectionConfig{dialector: sqc.dialector()}

	// Unknown Driver
	default:
		panic("Driver name is not available. Current supported driver: postgres, sqlite")
	}
}
