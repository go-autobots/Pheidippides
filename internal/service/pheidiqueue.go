package service

import (
	"context"

	v1 "Pheidippides/api/pheidiqueue/v1"
)

func (s *PheidiQueueService) SendTo(ctx context.Context, req *v1.Pheidippides) (*v1.SendResponse, error) {
	return &v1.SendResponse{}, nil
}
func (s *PheidiQueueService) GetFrom(ctx context.Context, req *v1.Pheidippides) (*v1.PhiediResponse, error) {
	return &v1.PhiediResponse{}, nil
}
