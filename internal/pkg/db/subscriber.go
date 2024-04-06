package db

import (
	"fmt"
	"log"

	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/models"
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

func GetSubscribers() ([]models.Subscriber, error) {
	db := GetDB()

	var subs []models.Subscriber
	rows, err := db.Query("SELECT email FROM newsletter_subs")

	if err != nil {
		log.Println("error", err)
		return subs, fmt.Errorf("no subscribers")
	}

	defer rows.Close()

	for rows.Next() {
		var sub models.Subscriber
		err := rows.Scan(&sub.Email)
		if err != nil {
			log.Println("error", err)
			return subs, fmt.Errorf("no subscribers")
		}
		subs = append(subs, sub)
	}

	return subs, nil
}

func GetSubscriberByid(id int, subscriber *models.Subscriber) error {
	db := GetDB()
	err := db.QueryRow("SELECT * newsletter_subs WHERE id = $1", id).Scan(&subscriber)
	if err != nil {
		log.Println("error", err)
		return err
	}

	return nil
}

func GetSubscriberByEmail(email string, subscriber *models.Subscriber) error {
	db := GetDB()
	err := db.QueryRow("SELECT * newsletter_subs WHERE email = $1", email).Scan(&subscriber)
	if err != nil {
		log.Println("error", err)
		return err
	}

	return nil
}
