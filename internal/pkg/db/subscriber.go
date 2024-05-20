package db

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"

	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/models"
	"github.com/gofiber/fiber/v2/log"
	"github.com/lib/pq"
)

func CreateSubscriber(email string) error {
	deleteKeyNum, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		log.Warn("error: ", err)
		return err
	}
	deleteKey := deleteKeyNum.Text(62)
	db := GetDB()
	_, err = db.Exec("INSERT INTO newsletter_subs (email, delete_key) VALUES ($1, $2)", email, deleteKey)
	if err == nil {
		return nil
	}
	log.Warn("error: ", err)
	if err, ok := err.(*pq.Error); ok && err.Code != UNIQUE_VIOLATION {
		return err
	} else if ok && err.Column == "delete_key" {
		return CreateSubscriber(email)
	} else if ok && err.Column == "email" {
		return fmt.Errorf("subscriber with email %s already exists", email)
	}
	return err
}

func DeleteSubscriberByEmail(email, deleteKey string) error {
	db := GetDB()

	var dbDeleteKey string
	err := db.QueryRow("SELECT delete_key FROM newsletter_subs WHERE email = $1", email).Scan(&dbDeleteKey)
	if err != nil {
		log.Warn("error: ", err)
		return err
	}
	log.Info("dbDeleteKey: ", dbDeleteKey, "deleteKey: ", deleteKey)
	if dbDeleteKey != deleteKey {
		return fmt.Errorf("invalid delete key")
	}

	result, err := db.Exec("DELETE FROM newsletter_subs WHERE email = $1", email)
	if rows, err := result.RowsAffected(); err != nil {
		log.Warn("error: ", err)
	} else if rows == 0 {
		return fmt.Errorf("no subscriber with email %s", email)
	}
	return err
}

func DeleteSubscriberByID(id int, deleteKey string) error {
	db := GetDB()

	var dbDeleteKey string
	err := db.QueryRow("SELECT delete_key FROM newsletter_subs WHERE id = $1", id).Scan(&dbDeleteKey)
	if err != nil {
		log.Warn("error: ", err)
		return err
	}
	log.Info("dbDeleteKey: ", dbDeleteKey, "deleteKey: ", deleteKey)
	if dbDeleteKey != deleteKey {
		return fmt.Errorf("invalid delete key")
	}

	result, err := db.Exec("DELETE FROM newsletter_subs WHERE id = $1", id)
	if rows, err := result.RowsAffected(); err != nil {
		log.Warn("error: ", err)
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
		log.Warn("error: ", err)
		return subs, fmt.Errorf("no subscribers")
	}

	defer rows.Close()

	for rows.Next() {
		var sub models.Subscriber
		err := rows.Scan(&sub.Email)
		if err != nil {
			log.Warn("error: ", err)
			return subs, fmt.Errorf("no subscribers")
		}
		subs = append(subs, sub)
	}

	return subs, nil
}

func GetSubscriberByid(id int, subscriber *models.Subscriber) error {
	db := GetDB()
	err := db.QueryRow("SELECT email FROM newsletter_subs WHERE id = $1", id).Scan(&subscriber.Email)
	if err != nil {
		log.Warn("error: ", err)
		return err
	}

	return nil
}

func GetSubscriberByEmail(email string, subscriber *models.Subscriber) error {
	db := GetDB()
	err := db.QueryRow("SELECT email FROM newsletter_subs WHERE email = $1", email).Scan(&subscriber.Email)
	if err != nil {
		log.Warn("error: ", err)
		return err
	}

	return nil
}
