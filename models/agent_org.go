package models

import (
	"database/sql"
	"log"
)

type AgentOrg struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	Organization string `db:"organization"`
}

func createAgentOrgTable(database sql.DB) {

}

func createAgentOrg(database sql.DB, name string, organization string, accountOwner User) {
	// TODO apparently queries can also return other things at the same time...?
	_, err := database.Query("INSERT INTO Sites(Name, Organization, AccountOwner) VALUES ($1, $2, $3);", organization)

	if err != nil {
		log.Println("Error creating site.")
	}
}
