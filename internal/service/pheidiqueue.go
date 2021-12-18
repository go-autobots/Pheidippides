package service

import (
	"Pheidippides/internal/service/validator"
	"context"

	v1 "Pheidippides/api/pheidiqueue/v1"
)

const (
	topicPrefix = "Pheidippides_"
)

func pheidippidesKey(originKey string) string {
	return topicPrefix + originKey
}

func (s *PheidiQueueService) SendTo(ctx context.Context, req *v1.Pheidippides) (*v1.SendResponse, error) {
	// Pheidippides 的 contents 存入 topic 里
	valideErr := validator.Send(req)
	if valideErr != nil {
		return nil, valideErr
	}
	sendErr := s.q.ReceiveFrom(req.GetContent(), pheidippidesKey(req.GetTopic()))
	if sendErr != nil {
		return nil, validator.QueueError(sendErr.Error())
	}
	return &v1.SendResponse{IsSuccess: true}, nil
}

func (s *PheidiQueueService) GetFrom(ctx context.Context, req *v1.Pheidippides) (*v1.PhiediResponse, error) {
	// valid 不需要传 contents 从 topic 拿出一个 content
	valideErr := validator.Get(req)
	if valideErr != nil {
		return nil, valideErr
	}
	res, getErr := s.q.GetFrom(pheidippidesKey(req.GetTopic()))
	if getErr != nil {
		return nil, validator.QueueError(getErr.Error())
	}
	return &v1.PhiediResponse{Content: res}, nil
}
