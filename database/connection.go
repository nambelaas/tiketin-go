package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
)

var (
	DBConn *sql.DB
	err    error
)

func Init() {
	dbData := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("Database.Host"),
		viper.GetInt("Database.Port"),
		viper.GetString("Database.User"),
		viper.GetString("Database.Pass"),
		viper.GetString("Database.DbName"),
	)

	DBConn, err = sql.Open("postgres", dbData)

	err = DBConn.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success connect database")

	RunMigration()
}

func RunMigration() {
	migrations := &migrate.FileMigrationSource{
		Dir: "database/migrations",
	}

	n, err := migrate.Exec(DBConn, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(fmt.Sprintf("Migration failed: %v", err))
	}

	fmt.Printf("Applied %d migrations!\n", n)
}