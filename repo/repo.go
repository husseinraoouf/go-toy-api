package repo

import (
	"fmt"
	"scenario/setting"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	db     *gorm.DB
	tables []interface{}
)

// RegisterModel registers model, if initfunc provided, it will be invoked after data model sync
func RegisterModel(bean interface{}) {
	tables = append(tables, bean)
}

func InitDatabase() error {
	gormDB, err := newDatabaseConnection()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	setDefaultDatabaseConnection(gormDB)

	err = syncAllTables()
	if err != nil {
		return fmt.Errorf("failed to sync to database: %v", err)
	}

	err = seedDatabase()
	if err != nil {
		return fmt.Errorf("failed to seed the database: %v", err)
	}

	return nil
}

// newDatabaseConnection returns a new gorm DB from the configuration
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

// DBConnStr returns database connection string
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

// Set gorm DB globally
func setDefaultDatabaseConnection(conn *gorm.DB) {
	db = conn
}

// syncAllTables sync the schemas of all tables, is required by unit test code
func syncAllTables() error {
	err := db.AutoMigrate(tables...)
	if err != nil {
		return fmt.Errorf("models sync: %v", err)
	}

	return nil
}

func seedDatabase() error {

	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&allCards)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
