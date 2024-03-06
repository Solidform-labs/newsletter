package db

import (
	"fmt"
	"log"
)

func CreateSubscriber(email string) error {
	db := GetDB()
	_, err := db.Exec("INSERT INTO newsletter_subs (email) VALUES ($1)", email)
	if err != nil {
		log.Println("error", err)
	}
	return err
}

func DeleteSubscriber(email string) error {
	db := GetDB()
	result, err := db.Exec("DELETE FROM newsletter_subs WHERE email = $1", email)
	if rows, err := result.RowsAffected(); err != nil {
		log.Println("error", err)
	} else if rows == 0 {
		return fmt.Errorf("no subscriber with email %s", email)
	}
	return err
}
