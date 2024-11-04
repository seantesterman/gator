package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("cannot retrieve users: %w", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf(" * %v (current)\n", user.Name)
			continue
		}
		fmt.Printf(" * %v\n", user.Name)
	}
	return nil
}
