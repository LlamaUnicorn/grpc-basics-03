package link

import (
	"context"
	"log"

	"github.com/llamaunicorn/grpc-basics-03/internal/converter"
	desc "github.com/llamaunicorn/grpc-basics-03/pkg/link_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	linkObj, err := i.linkService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, url: %s, title: %s, description: %s, created_at: %v, updated_at: %v\n",
		linkObj.ID, linkObj.Info.URL, linkObj.Info.Title, linkObj.Info.Description, linkObj.CreatedAt, linkObj.UpdatedAt)

	return &desc.GetResponse{
		Link: converter.ToLinkFromService(linkObj),
	}, nil
}
