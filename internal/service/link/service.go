package link

import (
	"github.com/llamaunicorn/grpc-basics-03/internal/client/db"
	"github.com/llamaunicorn/grpc-basics-03/internal/repository"
	"github.com/llamaunicorn/grpc-basics-03/internal/service"
)

type serv struct {
	linkRepository repository.LinkRepository
	txManager      db.TxManager
}

func NewService(
	linkRepository repository.LinkRepository,
	txManager db.TxManager,
) service.LinkService {
	return &serv{
		linkRepository: linkRepository,
		txManager:      txManager,
	}
}