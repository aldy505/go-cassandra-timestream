package main

import (
	"context"
	"fmt"
	"time"
)

func (d *Deps) InsertEvent(ctx context.Context, event Event) error {
	if event.CreatedAt.IsZero() {
		event.CreatedAt = time.Now()
	}

	query := d.DB.Query(
		`INSERT INTO 
			event_timestream
			(actor, events)
		VALUES
			(
				?, 
				{ 
					{
						id: `+event.ID.String()+`, 
						type: `+Quote(event.Type)+`, 
						message: `+Quote(event.Message)+`, 
						created_at: `+Quote(event.CreatedAt.Format(time.RFC3339))+`
					}
				}
			)
		IF NOT EXISTS`,
		event.Actor,
	)

	err := query.WithContext(ctx).Exec()
	if err != nil {
		return fmt.Errorf("failed to insert event: %v", err)
	}

	return nil
}

func (d *Deps) UpdateEvent(ctx context.Context, event Event) error {
	if event.CreatedAt.IsZero() {
		event.CreatedAt = time.Now()
	}

	query := d.DB.Query(
		`UPDATE
			event_timestream
		SET
			events = events + {
				{
					id: `+event.ID.String()+`, 
					type: `+Quote(event.Type)+`, 
					message: `+Quote(event.Message)+`, 
					created_at: `+Quote(event.CreatedAt.Format(time.RFC3339))+`
				}
			}
		WHERE
			actor = ?`,
		event.Actor,
	)

	err := query.WithContext(ctx).Exec()
	if err != nil {
		return fmt.Errorf("failed to update event: %v", err)
	}

	return nil
}
