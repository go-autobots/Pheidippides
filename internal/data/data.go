/**
 * @project Pheidippides
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/18 15:22 12月
 **/
package data

import (
	"Pheidippides/internal/biz"
	"Pheidippides/pkg/infra"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewRedisType, infra.NewQueue, NewPeidiQueue)

func NewPeidiQueue(queue infra.Queue) biz.PheidiQueue {
	return &pheidiQueue{q: queue}
}

func NewRedisType() infra.QueueType {
	return infra.Redis
}
