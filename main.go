package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bruhjeshhh/gatorCLI/internal/config"
	_ "github.com/lib/pq"
)

func main() {

	resp, _ := config.Read()
	s := state{
		cfg: &resp,
	}
	appState := commands{
		commands: make(map[string]func(*state, command) error),
	}

	appState.register("login", handlerLogin)
	params := os.Args
	// fmt.Println(params[0])
	fmt.Println(len(params))
	if len(params) < 3 {
		log.Fatal("not enough args")
	}
	cmd := command{
		name: params[1],
		omfo: params[1:],
	}
	// fmt.Print(cmd, s)
	appState.run(&s, cmd)

}
