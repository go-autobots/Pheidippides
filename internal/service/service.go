/**
 * @project Pheidippides
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/18 15:53 12月
 **/
package service

import (
	"github.com/google/wire"
	v1 "pheidippides/api/pheidiqueue/v1"
	"pheidippides/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPheidiQueueService)


type PheidiQueueService struct {
	v1.UnimplementedPheidiQueueServer
	useCase  *biz.PheidiUseCase
}

func NewPheidiQueueService() *PheidiQueueService {
	return &PheidiQueueService{}
}

