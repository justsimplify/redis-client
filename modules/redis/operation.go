package redis

import (
	"context"
	"github.com/justsimplify/redis-client/modules"
)

func (r Client) Create(c context.Context, payload modules.RedisPayload) (interface{}, error) {
	rc := getClient(r)
	return rc.Set(c, payload.Key, payload.Value, 0).Result()
}

func (r Client) Delete(c context.Context, key string) (interface{}, error) {
	rc := getClient(r)
	return rc.Del(c, key).Result()
}

func (r Client) Read(c context.Context, key string) (interface{}, error) {
	rc := getClient(r)
	return rc.Get(c, key).Result()
}

func (r Client) Update(c context.Context, payload modules.RedisPayload) (interface{}, error) {
	_, err := r.Read(c, payload.Key)
	if err != nil {
		return nil, err
	}
	rc := getClient(r)
	return rc.Set(c, payload.Key, payload.Value, 0).Result()
}
