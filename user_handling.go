package main

import (
	"context"
	"database/sql"
	_ "database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bruhjeshhh/gatorCLI/internal/database"
	"github.com/google/uuid"
)

func (c *commands) run(s *state, cmd command) error {
	callback, ok := c.commands[cmd.name]
	if !ok {
		log.Fatal("map lookup failed")
	}
	err := callback(s, cmd)
	if err != nil {
		return fmt.Errorf("the function did not execute %w", err)
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.omfo) < 2 {
		return errors.New("lavde")
	}
	_, erro := s.db.GetUser(context.Background(), cmd.omfo[1])
	if erro != nil {
		return fmt.Errorf("couldn't find user: %w", erro)
	}

	err := s.cfg.SetUser(cmd.omfo[1])
	if err != nil {
		return err
	}
	fmt.Println("The user has been set", cmd.omfo[1])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	// fmt.Println("About to run command...")
	if len(cmd.omfo) == 1 {
		return errors.New("not enough args")
	}
	_, err := s.db.GetUser(context.Background(), cmd.omfo[1])
	if err == nil {
		log.Fatal("user already exists")
	}
	// fmt.Println("About to craete...")
	user, errbr := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: timeConverter(time.Now().UTC()),
		UpdatedAt: timeConverter(time.Now().UTC()),
		Name:      cmd.omfo[1],
	})
	if errbr != nil {
		return fmt.Errorf("couldn't create user: %w", errbr)
	}
	// fmt.Println("created, now setting")
	s.cfg.SetUser(cmd.omfo[1])

	fmt.Println("user was created:")
	fmt.Print(user)
	return nil
}

func timeConverter(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func resetDb(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("couldnt reset")
	}
	fmt.Print("reset successful")
	return nil
}

func getUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("query failed for sm reason uhuh")
	}
	currentuser := s.cfg.CurrentUserName
	for _, user := range users {
		if user == currentuser {
			fmt.Println(user, "(current)")
		} else {
			fmt.Println(user)
		}
	}
	return nil
}
