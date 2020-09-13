package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ConfigObj struct {
	MusicConfig    MusicConfig    `json:"music_config"`
	DatabaseConfig DataBaseConfig `json:"database_config"`
}

type DataBaseConfig struct {
	MongoConfig MongoConfig `json:"mongo_config"`
}

type MongoConfig struct {
	DBAddress           string `json:"db_address"`
	DBName              string `json:"db_name"`
	TimeOut             int    `json:"time_out"`
	MusicInfoCollection string `json:"music_info_collection"`
}

type MusicConfig struct {
	NeteaseCloudMusicConfig NeteaseCloudMusicConfig `json:"netease_cloud_music_config"`
	QQMusicConfig           QQMusicConfig           `json:"qq_music_config"`
}

type NeteaseCloudMusicConfig struct {
	BaseUrl         string                       `json:"base_url"`
	SearchPath      string                       `json:"search_path"`
	GetMusicUrlPath string                       `json:"get_music_url_path"`
	SearchQuery     NeteaseCloudMusicSearchQuery `json:"search_query"`
}

type NeteaseCloudMusicSearchQuery struct {
	Keywords string `json:"keywords"`
	Limit    string `json:"limit"`
	Type     string `json:"type"`
	Id       string `json:"id"`
}

type QQMusicConfig struct {
	BaseUrl         string `json:"base_url"`
	SearchPath      string `json:"search_path"`
	GetMusicUrlPath string `json:"get_music_url_path"`
	Guid            string `json:"guid"`
}

var Config ConfigObj

func init() {

	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	byteStream, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteStream, &Config)
	if err != nil {
		log.Fatal(err)
	}

}
