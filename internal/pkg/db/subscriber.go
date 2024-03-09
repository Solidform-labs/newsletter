package db

import (
	"fmt"
	"log"

	"github.com/lib/pq"
)

func CreateSubscriber(email string) error {
	db := GetDB()
	_, err := db.Exec("INSERT INTO newsletter_subs (email) VALUES ($1)", email)
	if err != nil {
		log.Println("error", err)
		if err, ok := err.(*pq.Error); ok && err.Code == UNIQUE_VIOLATION {
			return fmt.Errorf("subscriber with email %s already exists", email)
		}
	}
	return err
}

func DeleteSubscriberByEmail(email string) error {
	db := GetDB()
	result, err := db.Exec("DELETE FROM newsletter_subs WHERE email = $1", email)
	if rows, err := result.RowsAffected(); err != nil {
		log.Println("error", err)
	} else if rows == 0 {
		return fmt.Errorf("no subscriber with email %s", email)
	}
	return err
}

func DeleteSubscriberByID(id int) error {
	db := GetDB()
	result, err := db.Exec("DELETE FROM newsletter_subs WHERE id = $1", id)
	if rows, err := result.RowsAffected(); err != nil {
		log.Println("error", err)
	} else if rows == 0 {
		return fmt.Errorf("no subscriber with id %d", id)
	}
	return err
}
