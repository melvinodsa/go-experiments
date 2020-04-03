package cache

type CacheLike interface {
	Get(key string) (interface{}, bool)
	Put(key string, value interface{})
}
