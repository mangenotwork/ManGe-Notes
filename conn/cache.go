package conn

import (
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

type Caches struct {
	C *cache.Cache
}

//	AllCaches　所有缓存
var (
	caches = new(Caches)
)

func CachesInit() {
	caches.C = cache.New(60*24*7*time.Minute, 10*time.Minute)
	log.Println("初始化Caches")
}

func CachesSet(key string, value interface{}) {
	caches.C.Set(key, value, cache.DefaultExpiration)
	return
}

func CachesSetAlways(key string, value interface{}) {
	caches.C.Set(key, value, -1)
	return
}

func CachesGet(key string) (value interface{}, isOk bool) {
	value, isOk = caches.C.Get(key)
	return
}

func CachesGetStr(key string) (value string) {
	v, isOk := caches.C.Get(key)
	value = ""
	if isOk {
		value = v.(string)
	}
	return
}
