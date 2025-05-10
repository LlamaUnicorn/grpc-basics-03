package link

import (
	"context"

	"github.com/llamaunicorn/grpc-basics-03/internal/model"
)

func (s *serv) Create(ctx context.Context, info *model.LinkInfo) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.linkRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.linkRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}