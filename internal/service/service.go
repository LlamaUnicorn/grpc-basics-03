package service

import (
	"context"

	"github.com/llamaunicorn/grpc-basics-03/internal/model"
)

type NoteService interface {
	Create(ctx context.Context, info *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
}

type LinkService interface {
	Create(ctx context.Context, info *model.LinkInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Link, error)
}
