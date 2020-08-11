package main

import (
	"context"
)

// Repository for connecting app and database
type Repository interface {
	Close()

	AddLink(ctx context.Context, link string) error
	UpdateLink(ctx context.Context, linkID string, link string, oldLinkID string) error
	GetLinks(ctx context.Context) ([]Link, error)
	GetLinkByID(ctx context.Context, linkID string) (*Link, error)
	//GetLinkByUserID(ctx context.Context, userID string) ([]Link, error)
}
