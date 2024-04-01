package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-pg/migrations/v8"
	"github.com/joho/godotenv"

	"github.com/ljlimjk10/users-ms/db"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.

Usage:
  go run *.go <command> [args]
`

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}
	flag.Usage = usage
	flag.Parse()

	dbConnector := db.InitDBConnection()

	oldVersion, newVersion, err := migrations.Run(dbConnector.DB, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
