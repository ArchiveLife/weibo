package api

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type WeiboAPI struct {
	cache *cache.Cache
}

func NewWeiboAPI() *WeiboAPI {
	return &WeiboAPI{cache: cache.New(5*time.Minute, 1*time.Hour)}
}
