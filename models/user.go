package models

import (
	"database/sql"
	"encoding/json"
)

// agemt -> controller -> process data & save to db -> show user when requested the data using graph n shit

// User request -> controolers -> database/models

// POST /api/users -> Create()

type User struct {
	ID                    int    `db:"id" json:"id"`
	FirstName             string `db:"first_name,NOT NULL"` // _name
	LastName              string `db:"last_name,NOT NULL"`  // _name
	Email                 string `db:"email"`
	RegistrationTimestamp string `db:"registration_timestamp"`
	Verified              bool   `db:"verified"` // todo email verification, todo 2fa
	LastLoginTimestamp    string `db:"last_login_timestamp"`
	PasswordHash          string `db:"password_hash"` // todo sha512 hash?
	RoleID                int    `db:"role"`          // global role (level 10), must auto create entry in role table
	json.RawMessage
}

func (u *User) Exists(db *sql.DB) {

}

func (u *User) Create(db *sql.DB) error {
	// TODO apparently queries can also return other things at the same time...?
	//.In("INSERT INTO Users(FirstName, Organization, AccountOwner) VALUES ($1, $2, $3);", u.FirstName)
	stmt, err := db.Prepare("INSERT INTO Users(id, first_name, last_name, email, registration_timestamp, verified, ) VALUES( ?, ?, ?, ? )")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt) // Prepared statements take up server resources and should be closed after use.
	return nil
}

// when added to an organization, an entry is added to map their role, to the agent organization.

type UserOrganizations struct {
	ID                  int `db:"id"`
	AccountID           int `db:"account_id"`
	RoleID              int `db:"role_id"`
	AgentOrganizationID int `db:"org_id"`
}

// roles are not organization based, either, read-only, editor, or full permission, system level overrides

// 1 = read-only, 2=editor, 3=run tests, 10=

type UserRole struct {
	ID    int
	Name  string
	Level int
}

func createUser(database sql.DB, firstName string, lastName string, email string, password string) {
	//todo check if user with email exists, validate email and other fields
}
