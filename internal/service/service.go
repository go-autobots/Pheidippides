/**
 * @project Pheidippides
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/18 15:53 12月
 **/
package service

import (
	v1 "Pheidippides/api/pheidiqueue/v1"
	"Pheidippides/pkg/infra"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRedisType, infra.NewQueue, NewPheidiQueueService)

func NewRedisType() infra.QueueType {
	return infra.Redis
}

type PheidiQueueService struct {
	v1.UnimplementedPheidiQueueServer
	q infra.Queue
}

func NewPheidiQueueService(queue infra.Queue) *PheidiQueueService {
	return &PheidiQueueService{q: queue}
}
