package main

import (
	"context"
)

type (
	// Service interface for define app service
	Service interface {
		AddLink(ctx context.Context, link string) error
		UpdateLink(ctx context.Context, linkID string, link string, oldLinkID string) error
		GetLinks(ctx context.Context) ([]Link, error)
		GetLinkByID(ctx context.Context, linkID string) (*Link, error)
		//GetLinkByUserID(ctx context.Context, userID string) ([]Link, error)
	}

	// Link models
	Link struct {
		ID   string `json:"id"`
		Link string `json:"link"`
	}
)
