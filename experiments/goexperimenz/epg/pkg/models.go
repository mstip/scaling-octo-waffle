package epg

type Account struct {
	ID         int64  `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Email      string `db:"email" json:"email"`
	ImapServer string `db:"imap_server" json:"imap_server"`
	ImapPort   int    `db:"imap_port" json:"imap_port"`
	UseSSL     bool   `db:"use_ssl" json:"use_ssl"`
	Password   string `db:"password" json:"password"`
}

type Classification struct {
}
