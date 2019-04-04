package cache

import (
	"github.com/go-redis/redis"
)

//Client for redis
var Client *redis.Client

//will init the database connection
func init() {
	/*
	 * We will init the connection with the redis database
	 */
	Client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

//RedisCache will act as a redis cache process
func RedisCache(req Request) Request {
	switch req.Type {
	case READ:
		//Read operation
		req.Payload, _ = Client.Get(req.Key).Result()
	case WRITE:
		//Write operation
		//convert the payload to json string
		Client.Set(req.Key, req.Payload, 0)
	case DELETE:
		//Delete operation
		Client.Del(req.Key)
	}
	return req
}
