package linker

import(
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/segmentio/ksuid"
	"time"
)

type postgresRepository {
	db *pg.Conn
}

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
	
}

func (r *postgresRepository) UpdateLink(ctx context.Context, linkID string, link string) error {
	
}

func (r *postgresRepository) GetLinks(ctx context.Context) ([]Link, error) {
	
}

func (r *postgresRepository) GetLinkByID(ctx context.Context) (*Link, error) {
	
}
