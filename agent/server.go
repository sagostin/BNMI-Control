package agent

import "BNMI-Control/db"

func Start() {
	db.Connect()
}
