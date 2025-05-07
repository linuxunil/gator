package main

import "context"

func handleReset(st *state, cmd command) error {
	err := st.db.Reset(context.Background())
	if err != nil {
		return err
	}
	return nil
}
