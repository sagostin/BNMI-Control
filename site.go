package main

import (
	"database/sql"
	"log"
)

type Site struct {
	ID           int    `db:"name"`
	Name         string `db:"name"`
	Organization string `db:"name"`
	AccountID    int    `db:"name"`
}

func createTable(database sql.DB) {

}

func createSite(database sql.DB, name string, organization string, accountOwner User) {
	// TODO apparently queries can also return other things at the same time...?
	_, err := database.Query("INSERT INTO Sites(Name, Organization, AccountOwner) VALUES ($1, $2, $3);", organization)

	if err != nil {
		log.Println("Error creating site.")
	}
}
