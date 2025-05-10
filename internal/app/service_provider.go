package app

import (
	"context"
	linkRepository "github.com/llamaunicorn/grpc-basics-03/internal/repository/link"
	linkService "github.com/llamaunicorn/grpc-basics-03/internal/service/link"
	"log"

	"github.com/llamaunicorn/grpc-basics-03/internal/api/link"
	"github.com/llamaunicorn/grpc-basics-03/internal/api/note"
	"github.com/llamaunicorn/grpc-basics-03/internal/client/db"
	"github.com/llamaunicorn/grpc-basics-03/internal/client/db/pg"
	"github.com/llamaunicorn/grpc-basics-03/internal/client/db/transaction"
	"github.com/llamaunicorn/grpc-basics-03/internal/closer"
	"github.com/llamaunicorn/grpc-basics-03/internal/config"
	"github.com/llamaunicorn/grpc-basics-03/internal/repository"
	noteRepository "github.com/llamaunicorn/grpc-basics-03/internal/repository/note"
	"github.com/llamaunicorn/grpc-basics-03/internal/service"
	noteService "github.com/llamaunicorn/grpc-basics-03/internal/service/note"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager      db.TxManager
	noteRepository repository.NoteRepository
	linkRepository repository.LinkRepository

	noteService service.NoteService
	linkService service.LinkService

	noteImpl *note.Implementation
	linkImpl *link.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) NoteRepository(ctx context.Context) repository.NoteRepository {
	if s.noteRepository == nil {
		s.noteRepository = noteRepository.NewRepository(s.DBClient(ctx))
	}

	return s.noteRepository
}

func (s *serviceProvider) NoteService(ctx context.Context) service.NoteService {
	if s.noteService == nil {
		s.noteService = noteService.NewService(
			s.NoteRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.noteService
}

func (s *serviceProvider) NoteImpl(ctx context.Context) *note.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = note.NewImplementation(s.NoteService(ctx))
	}

	return s.noteImpl
}

func (s *serviceProvider) LinkRepository(ctx context.Context) repository.LinkRepository {
	if s.linkRepository == nil {
		s.linkRepository = linkRepository.NewRepository(s.DBClient(ctx))
	}

	return s.linkRepository
}

func (s *serviceProvider) LinkService(ctx context.Context) service.LinkService {
	if s.linkService == nil {
		s.linkService = linkService.NewService(
			s.LinkRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.linkService
}

func (s *serviceProvider) LinkImpl(ctx context.Context) *link.Implementation {
	if s.linkImpl == nil {
		s.linkImpl = link.NewImplementation(s.LinkService(ctx))
	}

	return s.linkImpl
}
