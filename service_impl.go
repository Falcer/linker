package linker

import(
	"context"
)

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddLink(ctx context.Context, link string) error {
	
}

func (s *service) UpdateLink(ctx context.Context, linkID string, link string) error {
	
}

func (s *service) GetLinks(ctx context.Context) ([]Link, error) {
	
}

func (s *service) GetLinkByID(ctx context.Context) (*Link, error) {
	
}
