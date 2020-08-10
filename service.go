package linker

import (
	"context"
)

type(
	Service interface {
		AddLink(ctx context.Context, link string) error
		UpdateLink(ctx context.Context, linkID string, link string) error
		GetLinks(ctx context.Context) ([]Link, error)
		GetLinkByID(ctx context.Context) (*Link, error)
		//GetLinkByUserID(ctx context.Context, userID string) ([]Link, error)
	}

	Link struct {
		ID string `json:"id"`
		Link string `json:"link"`
	}
)