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
	ParseTime bool `toml:parse_time`
	Loc       string
}

type Server struct {
	IP   string
	Port int
}

type Path struct {
	FfmpegPath       string `toml"ffmpeg_path"`
	StaticSourcePath string `toml:static_source_path`
}

type Config struct {
	DB     Mysql `toml:"mysql"'`
	Server `toml:"server"`
	Path   `toml:"path"`
}

var Info Config

func init() {
	if _, err := toml.DecodeFile("E:\\code\\GoCode\\src\\dousheng\\config\\config.toml", &Info); err != nil {
		panic(err)
	}
	// 去除左右的空格
	strings.Trim(Info.Server.IP, " ")
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
