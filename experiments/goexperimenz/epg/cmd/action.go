package main

import (
	epg "epg/pkg"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db , err := sqlx.Connect("postgres", "user=postgres dbname=dev password=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	account, err := epg.CreateAccount(db, &epg.Account{
		Name:       "woop",
		Email:      "woop@woop.de",
		ImapServer: "imap.outlook.com",
		ImapPort:   993,
		UseSSL:     true,
		Password:   "qqqq",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(account)

	accounts, err := epg.QueryAccounts(db)
	log.Println(accounts)
}
