package main

import (
	"context"
	"fmt"
)

func (d *Deps) Migrate(ctx context.Context) error {
	err := d.DB.Query(
		`CREATE TYPE IF NOT EXISTS event(
			id uuid,
			type varchar,
			message varchar,
			created_at timestamp
		)`).WithContext(ctx).Exec()
	if err != nil {
		return fmt.Errorf("failed to create event table: %v", err)
	}
	err = d.DB.Query(
		`CREATE TABLE IF NOT EXISTS event_timestream(
			actor varchar,
			events set<frozen<event>>,
			PRIMARY KEY (actor)
		)`).WithContext(ctx).Exec()
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	return nil
}
