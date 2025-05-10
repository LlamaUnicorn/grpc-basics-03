package link

import (
	"github.com/llamaunicorn/grpc-basics-03/internal/service"
	desc "github.com/llamaunicorn/grpc-basics-03/pkg/link_v1"
)

type Implementation struct {
	desc.UnimplementedLinkV1Server
	linkService service.LinkService
}

func NewImplementation(linkService service.LinkService) *Implementation {
	return &Implementation{
		linkService: linkService,
	}
}
