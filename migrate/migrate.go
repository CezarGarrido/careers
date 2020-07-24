package migrate

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/pressly/goose"
)

var flags = flag.NewFlagSet("migrate", flag.ExitOnError)

func New(dialect, directory string, db *sql.DB) error {

	dir := flags.String("dir", directory, "directory with migration files")
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return nil
	}

	command := args[0]

	switch command {
	case "create":
		if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
			return err
		}
		os.Exit(0)
	case "fix":
		if err := goose.Run("fix", nil, *dir); err != nil {
			return err
		}
		os.Exit(0)
	}
	defer db.Close()
	if err := goose.SetDialect(dialect); err != nil {
		return err
	}
	if err := goose.Run(command, db, *dir, args[1:]...); err != nil {
		return err
	}
	return nil
}
func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

var (
	usagePrefix = `Usage: migrate [OPTIONS] COMMAND
Examples:
    migrate status
Options:
`
	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
`
)
