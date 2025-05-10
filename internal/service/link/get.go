package link

import (
	"context"

	"github.com/llamaunicorn/grpc-basics-03/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.Link, error) {
	link, err := s.linkRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return link, nil
}