package epg

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"reflect"
	"testing"
)

func SetupTestDB() (*sqlx.DB, *migrate.Migrate) {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=test password=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	//if err != nil {
	//	log.Fatal(err)
	//}
	return db, m
}

func CleanUpTestDB(db *sqlx.DB, m *migrate.Migrate) {
	err := m.Down()
	//if err != nil {
	//	log.Fatal(err)
	//}
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateAccount(t *testing.T) {
	tests := []struct {
		name    string
		args    *Account
		want    *Account
		wantErr bool
	}{
		{
			name: "add new account",
			args: &Account{
				Name:       "woop",
				Email:      "woop@woop.de",
				ImapServer: "imap.outlook.com",
				ImapPort:   993,
				UseSSL:     true,
				Password:   "qqqq",
			},
			want: &Account{
				ID:         1,
				Name:       "woop",
				Email:      "woop@woop.de",
				ImapServer: "imap.outlook.com",
				ImapPort:   993,
				UseSSL:     true,
				Password:   "qqqq",
			},
			wantErr: false,
		},
		{
			name: "add new with no data account",
			args: &Account{},
			want: &Account{
				ID:         1,
				Name:       "",
				Email:      "",
				ImapServer: "",
				ImapPort:   0,
				UseSSL:     false,
				Password:   "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, m := SetupTestDB()
			defer CleanUpTestDB(db, m)
			got, err := CreateAccount(db, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryAccounts(t *testing.T) {
	var empty []Account

	type args struct {
		accounts []Account
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "query with no data",
			args: args{
				accounts: empty,
			},
			wantErr: false,
		},
		{
			name: "query with entries",
			args: args{
				accounts: []Account{
					{
						ID:         1,
						Name:       "woop",
						Email:      "woop@woop.de",
						ImapServer: "imap.outlook.com",
						ImapPort:   993,
						UseSSL:     true,
						Password:   "qqqq",
					},
					{
						ID:         2,
						Name:       "test",
						Email:      "test@test.de",
						ImapServer: "test.test.test",
						ImapPort:   1337,
						UseSSL:     false,
						Password:   "test",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, m := SetupTestDB()
			defer CleanUpTestDB(db, m)

			for _, account := range tt.args.accounts {
				_, err := CreateAccount(db, &account)
				if err != nil {
					log.Fatal("fail to create account" + err.Error())
				}
			}

			got, err := QueryAccounts(db)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.args.accounts) {
				t.Errorf("QueryAccounts() got = %v, want %v", got, tt.args.accounts)
			}
		})
	}
}
