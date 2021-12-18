/**
 * @project Pheidippides
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/18 15:53 12月
 **/
package service

import (
	v1 "Pheidippides/api/pheidiqueue/v1"
	"Pheidippides/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPheidiQueueService)

type PheidiQueueService struct {
	v1.UnimplementedPheidiQueueServer
	useCase *biz.PheidiUseCase
}

func NewPheidiQueueService(uc *biz.PheidiUseCase) *PheidiQueueService {
	return &PheidiQueueService{useCase: uc}
}
