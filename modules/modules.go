package modules

import (
	"context"
)

func GenerateContext() context.Context {
	return context.Background()
}

type RedisPayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreateAPI interface {
	Create(c context.Context, payload RedisPayload) (interface{}, error)
}

type DeleteAPI interface {
	Delete(c context.Context, key string) (interface{}, error)
}

type UpdateAPI interface {
	Update(c context.Context, payload RedisPayload) (interface{}, error)
}

type ReadAPI interface {
	Read(c context.Context, key string) (interface{}, error)
}

