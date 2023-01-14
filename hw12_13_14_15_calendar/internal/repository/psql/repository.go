package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	_ "github.com/jackc/pgx/v4/stdlib" //nolint:blank-imports
)

var ErrNotFound = errors.New("event not found")

type Repo struct {
	db *sql.DB
}

func NewRepo() repository.DatabaseRepo {
	return new(Repo)
}

func (r *Repo) Connect(ctx context.Context) (err error) {
	r.db, err = sql.Open("pgx", getDSN())
	if err != nil {
		return err // failed to load driver
	}
	return r.db.PingContext(ctx)
}

func (r *Repo) Close() error {
	return r.db.Close()
}

func getDSN() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.Configuration.DB.Host,
		config.Configuration.DB.Port,
		config.Configuration.DB.User,
		config.Configuration.DB.Password,
		config.Configuration.DB.Database)
}

func (r *Repo) Create(ctx context.Context, event repository.Event) (repository.EventID, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return event.ID, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	query := `INSERT INTO events (title, description, date_start, 
                    date_end, owner_id, notify_in) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = tx.QueryRowContext(
		ctx,
		query,
		event.Title, event.Description, event.DateStart, event.DateEnd, event.OwnerID, event.NotifyIn).Scan(&event.ID)
	if err != nil {
		return event.ID, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return event.ID, err
	}

	return event.ID, nil
}

func (r *Repo) Update(ctx context.Context, id repository.EventID, event repository.Event) (repository.EventID, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return id, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	query := `UPDATE events SET title = $1, description = $2, date_start = $3, 
                  date_end = $4, owner_id = $5, notify_in = $6 WHERE id = $7`

	result, err := tx.ExecContext(
		ctx,
		query,
		event.Title, event.Description, event.DateStart, event.DateEnd, event.OwnerID, event.NotifyIn, id)
	if err != nil {
		return id, err
	}
	notFound, err := result.RowsAffected()
	if err != nil {
		return id, err
	}
	if notFound == 0 {
		return id, ErrNotFound
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return event.ID, err
	}

	return id, nil
}

func (r *Repo) Delete(ctx context.Context, id repository.EventID) (repository.EventID, error) {
	query := `DELETE FROM events WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return id, err
	}
	notFound, err := result.RowsAffected()
	if err != nil {
		return id, err
	}
	if notFound == 0 {
		return id, ErrNotFound
	}

	return id, nil
}

func (r *Repo) PeriodList(
	ctx context.Context,
	startPeriod time.Time,
	period repository.Period,
) ([]repository.Event, error) {
	var endPeriod time.Time
	switch period {
	case "Day":
		endPeriod = startPeriod.AddDate(0, 0, 1)
	case "Week":
		endPeriod = startPeriod.AddDate(0, 0, 7)
	case "Month":
		endPeriod = startPeriod.AddDate(0, 1, 0)
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	query := "SELECT * FROM events e WHERE e.date_end > $1 AND $2 > e.date_start"

	rows, err := tx.QueryContext(
		ctx,
		query,
		startPeriod.Format("2006-01-02 15:04"),
		endPeriod.Format("2006-01-02 15:04"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []repository.Event
	for rows.Next() {
		var event repository.Event
		if err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.DateStart,
			&event.DateEnd,
			&event.OwnerID,
			&event.NotifyIn); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *Repo) DayList(ctx context.Context, startDate time.Time) ([]repository.Event, error) {
	period := repository.Period("Day")
	return r.PeriodList(ctx, startDate, period)
}

func (r *Repo) WeekList(ctx context.Context, startDate time.Time) ([]repository.Event, error) {
	period := repository.Period("Week")
	return r.PeriodList(ctx, startDate, period)
}

func (r *Repo) MonthList(ctx context.Context, startDate time.Time) ([]repository.Event, error) {
	period := repository.Period("Month")
	return r.PeriodList(ctx, startDate, period)
}
