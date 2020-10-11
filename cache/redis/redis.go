package redis

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

// Redis Redis
type Redis struct {
	*redis.Pool
}

// Config Config
type Config struct {
}

// New New
func New() *Redis {
	return nil
}

// Do Do
func (r *Redis) Do(ctx context.Context, commandName string, args ...interface{}) (reply interface{}, err error) {
	conn, err := r.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	reply, err = conn.Do(commandName, args...)
	conn.Close()
	return reply, err
}
