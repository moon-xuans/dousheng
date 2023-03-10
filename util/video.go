package util

import (
	"dousheng/cache"
	"dousheng/config"
	"dousheng/model"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"time"
)

func GetVideoFileUrl(fileName string) string {
	base := fmt.Sprintf("http://%s:%d/%s/%s/%s", config.Info.IP, config.Info.MPort, config.Info.Bucket, config.Info.VideoDir, fileName)
	return base
}

func GetImageFileUrl(fileName string) string {
	base := fmt.Sprintf("http://%s:%d/%s/%s/%s", config.Info.IP, config.Info.MPort, config.Info.Bucket, config.Info.ImageDir, fileName)
	return base
}

// NewFileName 根据userId + 用户发布的视频数量连接成独一无二的文件名
func NewFileName(userId int64) string {
	var count int64

	err := model.NewVideoDAO().QueryVideoCountByUserId(userId, &count)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%d-%d", userId, count)
}

// SaveImageFromVideo 将视频切一帧保存到本地
// isDebug用于控制是否打印出执行的ffmepg命令
func SaveImageFromVideo(videoPath, name string, isDebug bool) error {
	v2i := NewVideo2Image()
	if isDebug {
		v2i.Debug()
	}
	v2i.InputPath = videoPath
	v2i.OutputPath = filepath.Join(config.Info.StaticSourcePath, name+defaultImageSuffix)
	v2i.FrameCount = 1
	queryString, err := v2i.GetQueryString()
	if err != nil {
		return err
	}
	return v2i.ExecCommand(queryString)
}

// FillVideoListFields 填充每个视频的作者信息(因为作者与视频的一对多关系，数据库中存入的是作者的id)
// 当userId > 0，我们判断当前为登陆状态，其余情况为登陆状态，则不需要填充IsFavorite字段
func FillVideoListFields(userId int64, videos *[]*model.Video) (*time.Time, error) {
	size := len(*videos)
	if videos == nil || size == 0 {
		return nil, errors.New("util.FillVideoListFields videos 为空")
	}
	dao := model.NewUserInfoDAO()
	p := cache.NewProxyIndexMap()

	latestTime := (*videos)[size-1].CreatedAt // 获取最近的投稿时间
	// 添加作者信息，以及is_follow状态
	for i := 0; i < size; i++ {
		var userInfo model.UserInfo
		err := dao.QueryUserInfoById((*videos)[i].UserInfoId, &userInfo)
		if err != nil {
			continue
		}
		userInfo.IsFollow = p.GetUserRelation(userId, userInfo.Id) //根据cache更新是否被关注
		(*videos)[i].Author = userInfo
		// 填充有登陆信息的点赞状态
		if userId > 0 {
			(*videos)[i].IsFavorite = p.GetVideoFavorState(userId, (*videos)[i].Id)
		}
	}
	return &latestTime, nil
}
