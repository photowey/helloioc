package redis

import (
	"fmt"

	"github.com/alibaba/ioc-golang/autowire/singleton"
	"github.com/alibaba/ioc-golang/config"
	"github.com/go-redis/redis"

	normalRedis "github.com/alibaba/ioc-golang/extension/normal/redis"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:paramType=Param
// +ioc:autowire:constructFunc=Init
// +ioc:autowire:alias=RedisAlias

type Redis struct {
	NormalRedis    normalRedis.Redis `normal:"github.com/alibaba/ioc-golang/extension/normal/redis.Impl"`
	NormalDB1Redis normalRedis.Redis `normal:"github.com/alibaba/ioc-golang/extension/normal/redis.Impl,db1-redis"`
	NormalDB2Redis normalRedis.Redis `normal:"github.com/alibaba/ioc-golang/extension/normal/redis.Impl,db2-redis"`
	NormalDB3Redis normalRedis.Redis `normal:"github.com/alibaba/ioc-golang/extension/normal/redis.Impl,address=192.168.1.11:6379&db=3"`

	privateClient *redis.Client
}

type Param struct {
	RedisAddr string
}

func (p *Param) Init(a *Redis) (*Redis, error) {
	privateClient := redis.NewClient(&redis.Options{
		Addr: p.RedisAddr,
	})
	a.privateClient = privateClient
	return a, nil
}

func (a *Redis) Run() {
	if _, err := a.NormalRedis.Set("mykey", "db0", -1); err != nil {
		panic(err)
	}

	if _, err := a.NormalDB1Redis.Set("mykey", "db1", -1); err != nil {
		panic(err)
	}

	if _, err := a.NormalDB2Redis.Set("mykey", "db2", -1); err != nil {
		panic(err)
	}

	if _, err := a.NormalDB3Redis.Set("mykey", "db3", -1); err != nil {
		panic(err)
	}

	val1, err := a.NormalRedis.Get("mykey")
	if err != nil {
		panic(err)
	}
	fmt.Println("client0 get ", val1)

	val2, err := a.NormalDB1Redis.Get("mykey")
	if err != nil {
		panic(err)
	}
	fmt.Println("client1 get ", val2)

	val3, err := a.NormalDB2Redis.Get("mykey")
	if err != nil {
		panic(err)
	}
	fmt.Println("client2 get ", val3)

	val4, err := a.NormalDB3Redis.Get("mykey")
	if err != nil {
		panic(err)
	}
	fmt.Println("client3 get ", val4)

	singletonRedisClientAddress := "127.0.0.1:6379"
	normalRedisClientAddress := "127.0.0.1:6379"
	err = config.LoadConfigByPrefix("autowire.singleton.<RedisAlias>.param.RedisAddr", &singletonRedisClientAddress)
	if err != nil {
		panic(err)
	}
	err = config.LoadConfigByPrefix("autowire.normal.<github.com/alibaba/ioc-golang/extension/normal/redis.Impl>.db1-redis.param.address", &normalRedisClientAddress)
	if err != nil {
		panic(err)
	}
	fmt.Printf("autowire.singleton.<RedisAlias>.param.RedisAddr=%s\n", normalRedisClientAddress)
	fmt.Printf("autowire.normal.<github.com/alibaba/ioc-golang/extension/normal/redis.Impl>.db1-redis.param.address=%s\n", normalRedisClientAddress)

}

func Run() {
	paymentI, _ := singleton.GetImpl("RedisAlias")
	redisClient := paymentI.(*Redis)

	redisClient.Run()
}
