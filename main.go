package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/bruhjeshhh/gatorCLI/internal/config"
	"github.com/bruhjeshhh/gatorCLI/internal/database"
	_ "github.com/lib/pq"
	// "github.com/bruhjeshhh/gatorCLI/sql"
)

func main() {

	resp, _ := config.Read()
	dbURL := resp.DbURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("cant connect")
	}
	dbQueries := database.New(db)
	s := &state{
		db:  dbQueries,
		cfg: &resp,
	}
	appState := commands{
		commands: make(map[string]func(*state, command) error),
	}

	appState.register("login", handlerLogin)
	appState.register("register", handlerRegister)
	appState.register("reset", resetDb)
	appState.register("users", getUsers)
	params := os.Args
	// fmt.Println(params[0])
	// fmt.Println(params[1])
	// fmt.Println(params[2])
	// fmt.Println(len(params))
	// if len(params) <  {
	// 	log.Fatal("not enough args")
	// }

	cmd := command{
		name: params[1],
		omfo: params[1:],
	}

	ersr := appState.run(s, cmd)
	if ersr != nil {
		log.Fatal(ersr)
	}

}
