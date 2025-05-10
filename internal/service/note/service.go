package note

import (
	"github.com/llamaunicorn/grpc-basics-03/internal/client/db"
	"github.com/llamaunicorn/grpc-basics-03/internal/repository"
	"github.com/llamaunicorn/grpc-basics-03/internal/service"
)

type serv struct {
	noteRepository repository.NoteRepository
	txManager      db.TxManager
}

func NewService(
	noteRepository repository.NoteRepository,
	txManager db.TxManager,
) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
		txManager:      txManager,
	}
}
