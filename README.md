### Setup
- Imports
```go
import (
  "github.com/justsimplify/redis-client/module"
  "github.com/justsimplify/redis-client/modules/redis"
)
```
  
### Usage
- Create a redis client
```go
rc := redis.Client{
    Host:     <redis-host>,
    Port:     <redis-port>,
    Password: <password>,
}
```

- READ operation
```go
rc.Read(ctx context.Context, redisKey string)
```

- DELETE operation
```go
rc.Delete(ctx context.Context, redisKey string)
```

- CREATE operation
```go
payload := modules.RedisPayload{
    Key:   "redis-key",
    Value: "redis-value",
}
rc.Create(ctx context.Context, payload modules.RedisPayload)
```

- UPDATE operation
```go
payload := modules.RedisPayload{
    Key:   "redis-key",
    Value: "redis-value",
}
rc.Update(ctx context.Context, payload modules.RedisPayload)
```