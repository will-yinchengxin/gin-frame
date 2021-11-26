package redis

import (
	"frame/global"
	"github.com/gomodule/redigo/redis"
	"errors"
	"time"
)

type RedisClient struct {
	db *redis.Pool
}

func GetRedisDB(key string) (redisClient *RedisClient, rdb *redis.Pool, err error) {
	rdbPool, ok := global.Redis[key]
	if !ok { //redis的 DB不存在 重新获取
		global.Logger.Fatalf("GetRedisDB fail, time %d", time.Now().Local().UnixNano())
		return &RedisClient{}, nil, errors.New("get redis conn fail")
	}

	// 这里将封装的 redis 及原生的连接都返回, 方便选择使用任何一类redis命令
	return &RedisClient{
		db: rdbPool,
	}, rdbPool, nil
}

func (r *RedisClient) Get(key string) (string, error) {
	conn := r.db.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

func (r *RedisClient) Set(key string, value string) (int64, error) {
	conn := r.db.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("SET", key, value))
}

func (r *RedisClient) SETEX(key string, sec int, value string) (int64, error) {
	conn := r.db.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("SETEX", key, sec, value))
}

func (r *RedisClient) Expire(key string, sec int64) (int64, error) {
	conn := r.db.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("EXPIRE", key, sec))
}

func (r *RedisClient) HGETALL(key string) (map[string]string, error) {
	conn := r.db.Get()
	defer conn.Close()
	return redis.StringMap(conn.Do("HGETALL", key))
}

func (r *RedisClient) HGET(key string, field string) (string, error) {
	conn := r.db.Get()
	defer conn.Close()
	return redis.String(conn.Do("HGET", key, field))
}

func (r *RedisClient) HSET(key string, field string, value string) (int64, error) {
	conn := r.db.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("HSET", key, field, value))
}
