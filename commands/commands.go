package commands

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/urfave/cli"
	tinysearch "github.com/wasuken/go-tinysearch"
)

var engine *tinysearch.Engine

func Main() {
	app := cli.NewApp()
	app.Name = "tinysearch"
	app.Usage = `simple and small search engine for learning`
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		createIndexCommand,
		searchCommand,
	}
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/tinysearch")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	engine = tinysearch.NewSearchEngine(db)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

const (
	exactArgs = iota
	minArgs
	maxArgs
)

func checkArgs(context *cli.Context, expected, checkType int) error {
	var err error
	cmdName := context.Command.Name
	switch checkType {
	case exactArgs:
		if context.NArg() != expected {
			err = fmt.Errorf(
				"%s: %q requres exactly %d argument(s)",
				os.Args[0], cmdName, expected,
			)
		}
	case minArgs:
		if context.NArg() < expected {
			err = fmt.Errorf(
				"%s: %q requres minimum %d argument(s)",
				os.Args[0], cmdName, expected,
			)
		}
	case maxArgs:
		if context.NArg() < expected {
			err = fmt.Errorf(
				"%s: %q requres maximum %d argument(s)",
				os.Args[0], cmdName, expected,
			)
		}
	}
	if err != nil {
		fmt.Printf("Incorrect Usage.\n\n")
		cli.ShowCommandHelp(context, cmdName)
		return err
	}
	return nil

}
