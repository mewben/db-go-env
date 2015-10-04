package db

import (
	"os"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	DbHost  string
	DbPort  string
	DbName  string
	DbUser  string
	DbPass  string
	SslMode string
}

// Global DB
var Conn *sqlx.DB

func Setup(db Database) {
	conn_string := os.Getenv("DATABASE_URL")

	if conn_string == "" {
		// dev mode
		// manually setup connection string
		conn_string = `host=` + db.DbHost + ` ` +
			`port=` + db.DbPort + ` ` +
			`dbname=` + db.DbName + ` ` +
			`user=` + db.DbUser + ` ` +
			`password=` + db.DbPass + ` ` +
			`sslmode=` + db.SslMode
	}

	// connect to database
	registerDatabase(conn_string)
}

func registerDatabase(conn_string string) {
	Conn = sqlx.MustConnect("postgres", conn_string)
	Conn = Conn.Unsafe()
}
