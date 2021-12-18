/**
 * @project Pheidippides
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/18 15:22 12月
 **/
package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewPheidiUseCase)

func NewPheidiUseCase(queue PheidiQueue) *PheidiUseCase{
	return &PheidiUseCase{queue: queue}
}

