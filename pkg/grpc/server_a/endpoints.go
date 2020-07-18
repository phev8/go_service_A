package server_a

import (
	"context"

	"github.com/phev8/go_commons/pkg/api/shared"
	"github.com/phev8/go_service_A/pkg/api"
)

func (s *serviceAServer) GetDataFromB(ctx context.Context, req *shared.RequestObject) (*api.DataObjectA, error) {
	resp, err := s.clients.ServiceB.GetData(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.DataObjectA{
		ExtraInfo:    "hello from A",
		CommonObject: resp.CommonObject,
	}, nil
}
