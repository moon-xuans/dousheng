package cache

import (
	"context"
	"dousheng/config"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// 用户id -> 被点赞的视频id集合 -> 是否含有该视频id

var ctx = context.Background()
var rdb *redis.Client

const (
	favor    = "favor"
	relation = "relation"
)

func init() {
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Info.RDB.IP, config.Info.RDB.Port),
			Password: "",
			DB:       config.Info.RDB.Database,
		})
}

var (
	proxyIndexOperation ProxyIndexMap
)

type ProxyIndexMap struct {
}

func NewProxyIndexMap() *ProxyIndexMap {
	return &proxyIndexOperation
}

// GetVideoFavorState 得到点赞状态
func (i *ProxyIndexMap) GetVideoFavorState(userId, videoId int64) bool {
	key := fmt.Sprintf("%s:%d", favor, userId)
	ret := rdb.SIsMember(ctx, key, videoId)
	return ret.Val()
}

// GetUserRelation 得到关注状态
func (i *ProxyIndexMap) GetUserRelation(userId, followId int64) bool {
	key := fmt.Sprintf("%s:%d", relation, userId)
	ret := rdb.SIsMember(ctx, key, followId)
	return ret.Val()
}
