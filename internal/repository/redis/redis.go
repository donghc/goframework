package redis

import (
	"github.com/go-redis/redis/v7"
	"goframework/configs"
	"goframework/pkg/errors"
	"goframework/pkg/timeutil"
	"goframework/pkg/trace"
	"strings"
	"time"
)

type Option func(*option)

type Trace = trace.T

type option struct {
	Trace *trace.Trace
	Redis *trace.Redis
}

func newOption() *option {
	return &option{}
}

var _ Repo = (*cacheRepo)(nil)

type Repo interface {
	i()
	Set(key, value string, ttl time.Duration, options ...Option) error
	Get(key string, options ...Option) (string, error)
	TTL(key string) (time.Duration, error)
	Expire(key string, ttl time.Duration) bool
	ExpireAt(key string, ttl time.Time) bool
	Del(key string, options ...Option) bool
	Exists(keys ...string) bool
	Incr(key string, options ...Option) int64
	Close() error
	Version() string
}

type cacheRepo struct {
	client *redis.Client
}

func New() (Repo, error) {
	client, err := redisConnect()
	if err != nil {
		return nil, err
	}
	return &cacheRepo{client: client}, nil
}

func redisConnect() (*redis.Client, error) {
	cfg := configs.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Pass,
		DB:           cfg.Db,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis err")
	}

	return client, nil
}

func (c *cacheRepo) i() {}

func (c *cacheRepo) Set(key, value string, ttl time.Duration, options ...Option) error {
	ts := time.Now()
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Redis.Timestamp = timeutil.CSTLayoutString()
			opt.Redis.Handle = "set"
			opt.Redis.Key = key
			opt.Redis.Value = value
			opt.Redis.TTL = ttl.Minutes()
			opt.Redis.CostSeconds = time.Since(ts).Seconds()
			opt.Trace.AppendRedis(opt.Redis)
		}
	}()

	for _, f := range options {
		f(opt)
	}

	if err := c.client.Set(key, value, ttl).Err(); err != nil {
		return errors.Wrapf(err, "redis set key : %s err", err)
	}
	return nil
}

func (c *cacheRepo) Get(key string, options ...Option) (string, error) {
	ts := time.Now()
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Redis.Timestamp = timeutil.CSTLayoutString()
			opt.Redis.Handle = "get"
			opt.Redis.Key = key
			opt.Redis.CostSeconds = time.Since(ts).Seconds()
			opt.Trace.AppendRedis(opt.Redis)
		}
	}()

	for _, f := range options {
		f(opt)
	}

	value, err := c.client.Get(key).Result()
	if err != nil {
		return "", errors.Wrapf(err, "redis get key: %s err", key)
	}

	return value, nil
}

func (c *cacheRepo) Del(key string, options ...Option) bool {
	ts := time.Now()
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Redis.Timestamp = timeutil.CSTLayoutString()
			opt.Redis.Handle = "del"
			opt.Redis.Key = key
			opt.Redis.CostSeconds = time.Since(ts).Seconds()
			opt.Trace.AppendRedis(opt.Redis)
		}
	}()

	for _, f := range options {
		f(opt)
	}

	if key == "" {
		return true
	}

	value, _ := c.client.Del(key).Result()
	return value > 0
}

func (c *cacheRepo) TTL(key string) (time.Duration, error) {
	ttl, err := c.client.TTL(key).Result()
	if err != nil {
		return -1, errors.Wrapf(err, "redis get key: %s err", key)
	}

	return ttl, nil
}

func (c *cacheRepo) Expire(key string, ttl time.Duration) bool {
	ok, _ := c.client.Expire(key, ttl).Result()
	return ok
}

func (c *cacheRepo) ExpireAt(key string, ttl time.Time) bool {
	ok, _ := c.client.ExpireAt(key, ttl).Result()
	return ok
}

func (c *cacheRepo) Exists(keys ...string) bool {
	if len(keys) == 0 {
		return true
	}
	value, _ := c.client.Exists(keys...).Result()
	return value > 0
}

func (c *cacheRepo) Incr(key string, options ...Option) int64 {
	ts := time.Now()
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Redis.Timestamp = timeutil.CSTLayoutString()
			opt.Redis.Handle = "incr"
			opt.Redis.Key = key
			opt.Redis.CostSeconds = time.Since(ts).Seconds()
			opt.Trace.AppendRedis(opt.Redis)
		}
	}()

	for _, f := range options {
		f(opt)
	}
	value, _ := c.client.Incr(key).Result()
	return value
}

func (c *cacheRepo) Close() error {
	return c.client.Close()
}

func (c *cacheRepo) Version() string {
	server := c.client.Info("server").Val()
	spl1 := strings.Split(server, "# Server")
	spl2 := strings.Split(spl1[1], "redis_version:")
	spl3 := strings.Split(spl2[1], "redis_git_sha1:")
	return spl3[0]
}

// WithTrace 设置trace信息
func WithTrace(t Trace) Option {
	return func(o *option) {
		if t!=nil{
			o.Trace=t.(*trace.Trace)
			o.Redis=new(trace.Redis)
		}
	}
}