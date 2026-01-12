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

type state struct {
	db  *database.Queries
	cfg *config.Config
}

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
	// fmt.Print(rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml"))
	// log.Fatal()
	appState.register("login", handlerLogin)
	appState.register("register", handlerRegister)
	appState.register("reset", resetDb)
	appState.register("users", getUsers)
	appState.register("addfeed", addFeed)
	appState.register("feeds", fetchFeeds)
	appState.register("follow", addFollow)
	appState.register("following", getfollwedfeedsby_User)
	appState.register("unfollow", unfollow)

	params := os.Args

	cmd := command{
		name: params[1],
		omfo: params[1:],
	}

	ersr := appState.run(s, cmd)
	if ersr != nil {
		log.Fatal("error hui gya", ersr)
	}

}
