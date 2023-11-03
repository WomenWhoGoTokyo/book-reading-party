package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var dsn string
	flag.StringVar(&dsn, "d", "database.sqlite", "database")
	flag.Usage = func() {
		usage := "このコマンドラインツールは◯◯です。使い方〜"
		fmt.Print(usage)
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch flag.Arg(0) {
	case "authors":
		// err = showAuthors(db)
	case "titles":
		if flag.NArg() != 2 {
			flag.Usage()
			os.Exit(2)
		}
		// err = showTitles(db, flag.Arg(1))
	case "content":
		if flag.NArg() != 3 {
			flag.Usage()
			os.Exit(2)
		}
		// err = showContent(db, flag.Arg(1), flag.Arg(2))
	case "query":
		if flag.NArg() != 2 {
			flag.Usage()
			os.Exit(2)
		}
		// err = queryCongent(db, flag.Arg(1))
	}

	if err != nil {
		log.Fatal(err)
	}
}
