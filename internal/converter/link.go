package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/llamaunicorn/grpc-basics-03/internal/model"
	desc "github.com/llamaunicorn/grpc-basics-03/pkg/link_v1"
)

func ToLinkFromService(link *model.Link) *desc.Link {
	var updatedAt *timestamppb.Timestamp
	if link.UpdatedAt.Valid {
		updatedAt = timestamppb.New(link.UpdatedAt.Time)
	}

	return &desc.Link{
		Id:        link.ID,
		Info:      ToLinkInfoFromService(link.Info),
		CreatedAt: timestamppb.New(link.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToLinkInfoFromService(info model.LinkInfo) *desc.LinkInfo {
	return &desc.LinkInfo{
		Url:         info.URL,
		Title:       info.Title,
		Description: info.Description,
	}
}

func ToLinkInfoFromDesc(info *desc.LinkInfo) *model.LinkInfo {
	return &model.LinkInfo{
		URL:         info.Url,
		Title:       info.Title,
		Description: info.Description,
	}
}