package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/packr/v2"
	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
)

func DatabaseConnection() *sql.DB {

	db, err := sql.Open("mysql", "user:password@tcp(localhost:3399)/database")
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}

	fmt.Println("database connected!")

	return db
}

func DatabaseMigration(db *sql.DB) {

	//migration := &migrate.PackrMigrationSource{
	//	Box: packr.New("migrations", "./db/migration"),
	//}
	//
	//n, err := migrate.Exec(db, "mysql", migration, migrate.Up)
	//if err != nil {
	//	// Handle errors!
	//	log.Error().Msg(err.Error())
	//	return
	//}
	//log.Info().Msg("migrate success" + strconv.Itoa(n))

	// OR: Use migrations from a packr box
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./db/migration"),
	}
	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		// Handle errors!
		log.Error().Msg("Monyeeeeetttttt... " + err.Error())
		return
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
