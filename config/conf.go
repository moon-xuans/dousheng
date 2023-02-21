package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/BurntSushi/toml"
)

type Mysql struct {
	Host      string
	Port      int
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool `toml:"parse_time"`
	Loc       string
}

type Redis struct {
	IP       string
	Port     int
	Database int
}

type Server struct {
	IP   string
	Port int
}

type Path struct {
	FfmpegPath       string `toml:"ffmpeg_path"`
	StaticSourcePath string `toml:"static_source_path"`
}

type Minio struct {
	Bucket          string
	VideoDir        string `toml:"video_dir"`
	ImageDir        string `toml:"image_dir"`
	MPort           int    `toml:"minio_port"`
	AccessKeyId     string `toml:"access_key_id"`
	SecretAccessKey string `toml:"secret_access_key"`
}

type Config struct {
	DB     Mysql `toml:"mysql"`
	RDB    Redis `toml:"redis"`
	Server `toml:"server"`
	Path   `toml:"path"`
	Minio  `toml:"minio"`
}

var Info Config

const ConfFile = "./config/config.toml"

func init() {
	if _, err := toml.DecodeFile(ConfFile, &Info); err != nil {
		panic(err)
	}
	// 去除左右的空格
	strings.Trim(Info.Server.IP, " ")
	strings.Trim(Info.RDB.IP, " ")
	strings.Trim(Info.DB.Host, " ")
}

// Dsn 填充得到数据库连接字符串
func Dsn() string {
	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?"+
		""+
		"charset=%s&parseTime=%v&loc=%s",
		Info.DB.Username, Info.DB.Password,
		Info.DB.Host, Info.DB.Port, Info.DB.Database,
		Info.DB.Charset, Info.DB.ParseTime, Info.DB.Loc)
	log.Println(arg)
	return arg
}
