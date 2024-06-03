package sqlstorage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cheef/hw-final-project/internal/config"
	"github.com/cheef/hw-final-project/internal/domain/models"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"time"
)

type Storage struct {
	db  *sql.DB
	log *slog.Logger
}

func NewStorage(_ context.Context, cfg config.Storage, log *slog.Logger) (*Storage, error) {
	const op = "storage.Connect"

	storage := Storage{log: log}
	db, err := storage.Connect(cfg)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	storage.db = db

	return &storage, nil
}

func (s *Storage) Connect(cfg config.Storage) (*sql.DB, error) {
	const op = "sql.Open"

	dsn := GetStorageDSN(cfg)
	s.log.Debug("Connecting to storage", slog.String("DSN", dsn))
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionMaxLifetime) * time.Second)

	return db, nil
}

func (s *Storage) Stop(_ context.Context) error {
	return s.db.Close()
}

func (s *Storage) CreateExceptionList(_ context.Context, listType, CIDR string) (int64, error) {
	const op = "storage.CreateExceptionList"

	s.log.Debug(op, "creating record", slog.String("listType", listType), slog.String("cidr", CIDR))
	stmt, err := s.db.Prepare("INSERT INTO exception_lists (type, cidr) VALUES ($1, $2) RETURNING id")

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var id int64
	rows := stmt.QueryRow(listType, CIDR)

	err = rows.Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	s.log.Debug(op, "created", slog.Int64("ID", id))
	defer stmt.Close()

	return id, nil
}

func (s *Storage) DeleteExceptionList(_ context.Context, listType, CIDR string) error {
	const op = "storage.DeleteExceptionList"

	s.log.Debug(op, "deleting record", slog.String("listType", listType), slog.String("cidr", CIDR))
	stmt, err := s.db.Prepare("DELETE FROM exception_lists WHERE type = $1 AND cidr = $2")

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(listType, CIDR)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) ShowExceptionLists(_ context.Context) ([]models.ExceptionList, error) {
	const op = "storage.ShowExceptionLists"

	stmt, err := s.db.Prepare("SELECT id, type, cidr FROM exception_lists")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var exceptions []models.ExceptionList

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		exception := new(models.ExceptionList)

		err := rows.Scan(
			&exception.ID,
			&exception.Type,
			&exception.CIDR,
		)

		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		exceptions = append(exceptions, *exception)
	}

	return exceptions, nil
}

func GetStorageDSN(cfg config.Storage) string {
	res := os.Getenv("STORAGE_DSN")

	if res != "" {
		return res
	}

	return cfg.DSN
}
