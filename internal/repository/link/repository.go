package link

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/llamaunicorn/grpc-basics-03/internal/client/db"
	"github.com/llamaunicorn/grpc-basics-03/internal/model"
	"github.com/llamaunicorn/grpc-basics-03/internal/repository"
	"github.com/llamaunicorn/grpc-basics-03/internal/repository/link/converter"
	modelRepo "github.com/llamaunicorn/grpc-basics-03/internal/repository/link/model"
)

const (
	tableName = "link"

	idColumn          = "id"
	urlColumn         = "url"
	titleColumn       = "title"
	descriptionColumn = "description"
	createdAtColumn   = "created_at"
	updatedAtColumn   = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.LinkRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.LinkInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(urlColumn, titleColumn, descriptionColumn).
		Values(info.URL, info.Title, info.Description).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "link_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Link, error) {
	builder := sq.Select(idColumn, urlColumn, titleColumn, descriptionColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "link_repository.Get",
		QueryRaw: query,
	}

	var link modelRepo.Link
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&link.ID, &link.Info.URL, &link.Info.Title, &link.Info.Description, &link.CreatedAt, &link.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToLinkFromRepo(&link), nil
}