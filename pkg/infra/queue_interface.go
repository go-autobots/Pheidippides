/**
 * @project Pheidippides
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/18 11:57 12月
 **/
package infra

import (
	"Pheidippides/pkg/infra/queueImplement"
	"time"

	"Pheidippides/internal/conf"
)

type QueueType int

// QueueType 中间件 类型
const (
	// Redis 使用 redis
	Redis QueueType = iota + 1
)

// Queue 缓存通用接口
type Queue interface {
	Send(topic string, contents ...string) error
	Get(topic string) (string, error)
}

// NewQueue 返回实现 Cache 接口的对象
func NewQueue(
	queueType QueueType,
	dataConfig *conf.Redis,
) Queue {
	switch queueType {
	case Redis:
		if dataConfig.WriteTimeout > 1 || dataConfig.ReadTimeout > 1 {
			panic("redis timeout too lang")
		}
		return queueImplement.NewRedis(
			dataConfig.Password,
			dataConfig.Host,
			int(dataConfig.Port),
			int(dataConfig.Db),
			int(dataConfig.PoolSize),
			int(dataConfig.IdleConns),
			// 变为秒
			time.Duration(dataConfig.ReadTimeout * 10 * 1000000000),
			time.Duration(dataConfig.WriteTimeout * 10 * 1000000000),
		)
	default:
		panic("暂不支持的队列类型")
	}
}
