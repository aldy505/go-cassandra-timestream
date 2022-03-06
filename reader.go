package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gocql/gocql"
)

func (d *Deps) GetAllEvents(ctx context.Context) ([]EventTimestream, error) {
	rows := d.DB.Query(
		`SELECT 
			*
		FROM
			event_timestream`,
	).WithContext(ctx).Iter()

	records, err := rows.SliceMap()
	if err != nil {
		return []EventTimestream{}, fmt.Errorf("failed to iterate over rows: %v", err)
	}

	err = rows.Close()
	if err != nil {
		return []EventTimestream{}, fmt.Errorf("failed to close rows: %v", err)
	}

	data, err := json.Marshal(records)
	if err != nil {
		return []EventTimestream{}, fmt.Errorf("failed to marshal records: %v", err)
	}

	var events []EventTimestream
	if err := json.Unmarshal(data, &events); err != nil {
		return []EventTimestream{}, fmt.Errorf("failed to unmarshal records: %v", err)
	}

	return events, nil
}

func (d *Deps) GetEventsByActor(ctx context.Context, actor string) (EventTimestream, error) {
	rows := d.DB.Query(
		`SELECT 
			*
		FROM
			event_timestream
		WHERE
			actor = ?
		LIMIT 1`,
		actor,
	).WithContext(ctx).Iter()

	records, err := rows.SliceMap()
	if err != nil {
		return EventTimestream{}, fmt.Errorf("failed to iterate over rows: %v", err)
	}

	err = rows.Close()
	if err != nil {
		return EventTimestream{}, fmt.Errorf("failed to close rows: %v", err)
	}

	if len(records) < 1 {
		return EventTimestream{}, ErrNoRecordsFound
	}

	data, err := json.Marshal(records[0])
	if err != nil {
		return EventTimestream{}, fmt.Errorf("failed to marshal records: %v", err)
	}

	var events EventTimestream
	if err := json.Unmarshal(data, &events); err != nil {
		return EventTimestream{}, fmt.Errorf("failed to unmarshal records: %v", err)
	}

	return events, nil
}

func (d *Deps) ActorExists(ctx context.Context, actor string) (bool, error) {
	var exists bool
	err := d.DB.Query(
		`SELECT
			actor
		FROM
			event_timestream
		WHERE
			actor = ?
		LIMIT 1`,
		actor,
	).WithContext(ctx).Scan(&exists)
	if err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return false, nil
		}

		return false, fmt.Errorf("failed to check if actor exists: %v", err)
	}

	return exists, nil
}
