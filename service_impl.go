package main

import (
	"context"
)

type service struct {
	repo Repository
}

// NewService created new service app
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddLink(ctx context.Context, link string) error {
	return s.repo.AddLink(ctx, link)
}

func (s *service) UpdateLink(ctx context.Context, linkID string, link string, oldLinkID string) error {
	return s.repo.UpdateLink(ctx, linkID, link, oldLinkID)
}

func (s *service) GetLinks(ctx context.Context) ([]Link, error) {
	return s.repo.GetLinks(ctx)
}

func (s *service) GetLinkByID(ctx context.Context, linkID string) (*Link, error) {
	return s.repo.GetLinkByID(ctx, linkID)
}
