package global

import (
	"frame/pkg/setting"
	"github.com/gomodule/redigo/redis"
	"time"
)

func initRedisPool(conn *setting.RedisSetting) (*redis.Pool, error) {
	redisPool := &redis.Pool{
		MaxIdle:     conn.MaxIdleNum,
		MaxActive:   conn.MaxActive,
		IdleTimeout: time.Duration(conn.MaxIdleTimeout) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", conn.Host,
				redis.DialPassword(conn.Password), redis.DialDatabase(conn.Database),
				redis.DialConnectTimeout(time.Duration(conn.ConnectTimeout)*time.Second),
				redis.DialReadTimeout(time.Duration(conn.ReadTimeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(conn.ReadTimeout)*time.Second),
			)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	return redisPool, nil
}
