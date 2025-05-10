package converter

import (
	"github.com/llamaunicorn/grpc-basics-03/internal/model"
	modelRepo "github.com/llamaunicorn/grpc-basics-03/internal/repository/link/model"
)

func ToLinkFromRepo(link *modelRepo.Link) *model.Link {
	return &model.Link{
		ID:        link.ID,
		Info:      ToLinkInfoFromRepo(link.Info),
		CreatedAt: link.CreatedAt,
		UpdatedAt: link.UpdatedAt,
	}
}

func ToLinkInfoFromRepo(info modelRepo.LinkInfo) model.LinkInfo {
	return model.LinkInfo{
		URL:         info.URL,
		Title:       info.Title,
		Description: info.Description,
	}
}