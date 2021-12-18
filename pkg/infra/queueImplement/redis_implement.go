/**
 * @project Pheidippides
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/18 11:58 12月
 **/
package queueImplement



import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// RedisDriver redis 缓存驱动
type RedisDriver struct {
	client *redis.Client
}

// NewRedis 实例化 redis 客户端
func NewRedis(
	password,
	host string,
	port,
	db,
	poolSize,
	idleConns int,
	readTimeout,
	writeTimeout time.Duration,
) *RedisDriver {

	rd := &RedisDriver{}
	addr := fmt.Sprintf("%s:%d", host, port)
	rd.client = redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         addr,
		Password:     password,
		DB:           db,
		PoolSize:     poolSize,
		MinIdleConns: idleConns,
		ReadTimeout: readTimeout,
		WriteTimeout: writeTimeout,
	})

	_, err := rd.client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return rd
}

// Send
func (rd *RedisDriver)Send(topic string, contents ...string) error {
	return nil
}

// Get
func (rd *RedisDriver)Get(topic string) (string, error) {
	return "", nil
}
