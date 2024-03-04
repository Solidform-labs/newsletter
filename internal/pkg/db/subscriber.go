package db

import "fmt"

func CreateSubscriber(email string) error {
	db := GetDB()
	fmt.Println("db", db)
	_, err := db.Exec("INSERT INTO newsletter_subs (email) VALUES ($1)", email)
	fmt.Println("error", err)
	return err
}
