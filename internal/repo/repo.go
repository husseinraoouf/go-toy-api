package repo

import (
	"fmt"

	"scenario/internal/setting"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	tables    []interface{}
	seedFuncs []func() error
)

// GetDatabase returns a database connection.
func GetDatabase() *gorm.DB {
	return db
}

// Tables returns tables registered by repo models.
func Tables() []interface{} {
	return tables
}

// SeedFuncs returns seed functions registered by repo models.
func SeedFuncs() []func() error {
	return seedFuncs
}

// RegisterModel registers model, if initfunc provided, it will be invoked after data model sync.
func RegisterModel(bean interface{}, seedFunc ...func() error) {
	tables = append(tables, bean)

	if len(seedFuncs) > 0 && seedFunc[0] != nil {
		seedFuncs = append(seedFuncs, seedFunc[0])
	}
}

// InitDatabase initializes the database and store it in global variable `db`.
func InitDatabase() error {
	gormDB, err := newDatabaseConnection()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	setDefaultDatabaseConnection(gormDB)

	return nil
}

// newDatabaseConnection returns a new gorm DB from the configuration.
func newDatabaseConnection() (*gorm.DB, error) {
	connStr, err := DBConnStr()
	if err != nil {
		return nil, err
	}

	var conn *gorm.DB

	if setting.Config.Database.Type == "postgres" {
		conn, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	} else {
		conn, err = gorm.Open(sqlite.Open(connStr), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	return conn, nil
}

// DBConnStr returns database connection string.
func DBConnStr() (string, error) {
	connStr := ""

	switch setting.Config.Database.Type {
	case "postgres":
		connStr = setting.Config.Database.ConnectionString
	case "sqlite":
		connStr = fmt.Sprintf("file:%s", setting.Config.Database.Path)
	case "memory":
		connStr = "file::memory:"
	default:
		return "", fmt.Errorf("unknown database type: %s", setting.Config.Database.Type)
	}

	return connStr, nil
}

// Set gorm DB globally.
func setDefaultDatabaseConnection(conn *gorm.DB) {
	db = conn
}
