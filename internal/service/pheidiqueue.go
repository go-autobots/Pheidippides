package service

import (
	"context"

	v1 "Pheidippides/api/pheidiqueue/v1"
)

func (s *PheidiQueueService) SendTo(ctx context.Context, req *v1.Pheidippides) (*v1.SendResponse, error) {
	// todo Pheidippides 的 contents 存入 topic 里
	return &v1.SendResponse{}, nil
}
func (s *PheidiQueueService) GetFrom(ctx context.Context, req *v1.Pheidippides) (*v1.PhiediResponse, error) {
	// todo valid 不需要传 contents 从 topic 拿出一个 content
	return &v1.PhiediResponse{}, nil
}
