package epg

import "github.com/jmoiron/sqlx"

func QueryAccounts(db *sqlx.DB) ([]Account, error) {
	var accounts []Account

	err := db.Select(&accounts, "SELECT id, name, email, imap_server, imap_port, use_ssl, password FROM public.accounts;")
	if err != nil {
		return nil, err
	}
	return accounts, err
}

func CreateAccount(db *sqlx.DB, newAccount *Account) (*Account, error) {
	result, err := db.NamedQuery(
		`
			INSERT INTO 
    			public.accounts(name, email, imap_server, imap_port, use_ssl, password) 
    		VALUES (:name, :email, :imap_server, :imap_port, :use_ssl, :password) 
			RETURNING id, name, email, imap_server, imap_port, use_ssl, password
			`,
		newAccount,
	)
	if err != nil {
		return nil, err
	}
	for result.Next() {
		err = result.StructScan(newAccount)
		if err != nil {
			return nil, err
		}
	}
	return newAccount, nil
}
