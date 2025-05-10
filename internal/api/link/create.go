package link

import (
	"context"
	"log"

	"github.com/llamaunicorn/grpc-basics-03/internal/converter"
	desc "github.com/llamaunicorn/grpc-basics-03/pkg/link_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.linkService.Create(ctx, converter.ToLinkInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted link with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
