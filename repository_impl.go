package main

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/teris-io/shortid"
)

type postgresRepository struct {
	db *pgx.Conn
}

// NewPostgres created new postgres database connection
func NewPostgres(url string) (Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = r.db.Close(ctx)
}

func (r *postgresRepository) AddLink(ctx context.Context, link string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := r.db.Query(ctx, "INSERT INTO links (id, link) VALUES ($1, $2)", shortid.Generate, link)
	if res != nil {
		res.Close()
	}
	return err
}

func (r *postgresRepository) UpdateLink(ctx context.Context, linkID string, link string, oldLinkID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := r.db.Query(ctx, "UPDATE links SET id=$1, link=$2 WHERE=$3", linkID, link, oldLinkID)
	if res != nil {
		res.Close()
	}
	return err
}

func (r *postgresRepository) GetLinks(ctx context.Context) ([]Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := r.db.Query(ctx, "SELECT id, link FROM links")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var links []Link

	for rows.Next() {
		a := &Link{}
		if err = rows.Scan(&a.ID, &a.Link); err == nil {
			links = append(links, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return links, nil
}

func (r *postgresRepository) GetLinkByID(ctx context.Context, linkID string) (*Link, error) {
	a := &Link{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	row := r.db.QueryRow(ctx, "SELECT id, link FROM links WHERE id = $1", linkID)
	err := row.Scan(&a.ID, &a.Link)
	if err != nil {
		return nil, err
	}
	return a, nil
}
