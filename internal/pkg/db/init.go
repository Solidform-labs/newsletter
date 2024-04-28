package db

import "github.com/gofiber/fiber/v2/log"

func Init() {
	db := GetDB()

	createTablesSQL := `
	CREATE TABLE IF NOT EXISTS newsletter_subs (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE
	);
	`

	if res, err := db.Exec(createTablesSQL); err != nil {
		log.Fatalf("Error creating tables: %s", err)
	} else {
		rows, _ := res.RowsAffected()
		if rows == 0 {
			log.Info("Tables already exist")
		} else {
			log.Info("Tables created")
		}
	}
}
