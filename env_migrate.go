package env_migrate

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	migrate "github.com/rubenv/sql-migrate"
)

const _UP_MIGRATE = "up_migrate"
const _DOWN_MIGRATE = "down_migrate"
const _SKIP_MIGRATE = "skip_migrate"
const _DIR_MIGRATE = "dir_migrate"

var BaseRootDir = ""

func loadMigrationsDir() *migrate.FileMigrationSource {
	var dir = os.Getenv(_DIR_MIGRATE)
	if dir == "" {
		dir = "migrations"
	}

	migrations := &migrate.FileMigrationSource{
		Dir: path.Join(BaseRootDir, dir),
	}
	return migrations
}

func UpMigration(db *sql.DB, max int) {
	dir := loadMigrationsDir()
	n, err := migrate.ExecMax(db, "postgres", dir, migrate.Up, 9999)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("UpMigration Applied %d migrations!\n", n)
}

func DownMigration(db *sql.DB, max int) {
	dir := loadMigrationsDir()
	n, err := migrate.ExecMax(db, "postgres", dir, migrate.Down, 9999)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("DownMigration Applied %d migrations!\n", n)
}

func SkipMigration(db *sql.DB, max int) {
	dir := loadMigrationsDir()
	n, err := migrate.SkipMax(db, "postgres", dir, migrate.Up, 9999)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("SkipMigration Applied %d migrations!\n", n)
}

func getEnvKeyValue(keys ...string) (key string, value string) {
	for _, k := range keys {
		v := os.Getenv(k)
		if v != "" {
			key = k
			value = v
			break
		}
	}
	return
}

func Auto(db *sql.DB) {

	var max int
	var onlyMigrate = false

	key, value := getEnvKeyValue(_UP_MIGRATE, _DOWN_MIGRATE, _SKIP_MIGRATE)

	if value == "" {
		fmt.Println("No need migrate.")
		return
	}

	if value == "all" {
		value = "99999"
	}

	if os.Getenv("only_migrate") != "" {
		onlyMigrate = true
	}

	// var err error
	max, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalln(err)
	}

	if key == _UP_MIGRATE {
		UpMigration(db, max)
	} else if key == _DOWN_MIGRATE {
		DownMigration(db, max)
	} else if key == _SKIP_MIGRATE {
		SkipMigration(db, max)
	}
	// fmt.Printf("Run %v=%v, Done!\n", key, value)
	if onlyMigrate || key != _UP_MIGRATE {
		fmt.Println("Only run and exit.")
		os.Exit(0)
	}
}
