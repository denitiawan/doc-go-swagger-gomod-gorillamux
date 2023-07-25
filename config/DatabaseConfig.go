package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
)

func DatabaseConnection(appConfig AppConfig) *sql.DB {

	db, err := sql.Open(appConfig.DBDriverName, appConfig.DBUsername+":"+appConfig.DBPassword+"@tcp("+appConfig.DBHost+":"+appConfig.DBPort+")/"+appConfig.DBName+"?parseTime=true")
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}

	fmt.Println("database connected!")

	return db
}

func DatabaseMigration(appConfig AppConfig, db *sql.DB) {

	// Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migration",
	}

	// Execution all migration files
	n, err := migrate.Exec(db, appConfig.DBDriverName, migrations, migrate.Up)
	if err != nil {
		log.Error().Msg("Errror Migration... " + err.Error())
		return
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
