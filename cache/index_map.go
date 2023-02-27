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

// UpdateVideoFavorState 更新点赞状态，state:true为点赞，false为取消点赞
func (i *ProxyIndexMap) UpdateVideoFavorState(userId, videoId int64, state bool) {
	key := fmt.Sprintf("%s:%d", favor, userId)
	if state {
		rdb.SAdd(ctx, key, videoId)
		return
	}
	rdb.SRem(ctx, key, videoId)
}

// GetUserRelation 得到关注状态
func (i *ProxyIndexMap) GetUserRelation(userId, followId int64) bool {
	key := fmt.Sprintf("%s:%d", relation, userId)
	ret := rdb.SIsMember(ctx, key, followId)
	return ret.Val()
}

// UpdateUserRelation 更新点赞状态， state：true为点关注，false为取消关注
func (i *ProxyIndexMap) UpdateUserRelation(userId, followId int64, state bool) {
	key := fmt.Sprintf("%s:%d", relation, userId)
	if state {
		rdb.SAdd(ctx, key, followId)
		return
	}
	rdb.SRem(ctx, key, followId)
}
