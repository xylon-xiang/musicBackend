package controller

import (
	"MusicBackend/config"
	"MusicBackend/module"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type NeteaseCloudMusicSearchResult struct {
	NeedLogin bool   `json:"needLogin"`
	Result    Result `json:"result"`
}

type Result struct {
	Songs []Song `json:"songs"`
}

type Song struct {
	Name    string   `json:"name"`
	Id      int64    `json:"id"`
	Artists []Artist `json:"ar"`
}

type Artist struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type NeteaseCloudMusicUri struct {
	Data []NeteaseData `json:"data"`
	//Code string        `json:"code"`
}

type NeteaseData struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}

///*** below is the QQ dom ***///
type QQMusicSearchResult struct {
	Data QQSearchData `json:"data"`
}

type QQSearchData struct {
	List []QQDataObj `json:"list"`
}

type QQDataObj struct {
	SongName string `json:"songname"`
	SongMid  string `json:"songmid"`
	Singer   Singer `json:"singer"`
}

type Singer struct {
	Mid  string `json:"mid"`
	Name string `json:"name"`
}

// ++++++ //
type QQMusicUrl struct {
	Data QQData `json:"data"`
}

type QQData struct {
	MusicUrl string `json:"musicUrl"`
}

func getNeteaseCloudMusicUri(id int64) (string, error) {

	musicUrl, err := url.Parse("http://localhost:3000/song/url")
	if err != nil {
		return "", err
	}

	params := url.Values{}
	strId := strconv.FormatInt(id, 10)
	params.Set("id", strId)

	musicUrl.RawQuery = params.Encode()

	resp, err := http.Get(musicUrl.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var neteaseCloudMusicUri NeteaseCloudMusicUri
	err = json.Unmarshal(body, &neteaseCloudMusicUri)
	if err != nil {
		return "", err
	}

	return neteaseCloudMusicUri.Data[0].Url, nil
}

func SearchNeteaseCloudMusic(keywords string, limit int, searchType int) (*[]module.MusicInfo, error) {

	var searchResult NeteaseCloudMusicSearchResult

	//musicUrl, err := url.Parse(config.Config.MusicConfig.NeteaseCloudMusicConfig.BaseUrl +
	//	config.Config.MusicConfig.NeteaseCloudMusicConfig.SearchPath)
	musicUrl, err := url.Parse("http://localhost:3000/cloudsearch")
	if err != nil {
		return nil, err
	}

	params := url.Values{}

	params.Set("keywords", keywords)
	//params.Set(config.Config.MusicConfig.NeteaseCloudMusicConfig.SearchQuery.Limit, string(limit))
	//params.Set(config.Config.MusicConfig.NeteaseCloudMusicConfig.SearchQuery.Type, string(searchType))

	musicUrl.RawQuery = params.Encode()

	urlPath := musicUrl.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &searchResult)
	if err != nil {
		return nil, err
	}

	musicInfo := make([]module.MusicInfo, len(searchResult.Result.Songs))

	for key, song := range searchResult.Result.Songs {
		songUri, err := getNeteaseCloudMusicUri(song.Id)
		if err != nil {
			return nil, err
		}

		songId := strconv.FormatInt(song.Id, 10)
		arId := strconv.FormatInt(song.Artists[0].Id, 10)

		musicInfo[key].ArtistId = arId
		musicInfo[key].ArtistName = song.Artists[0].Name
		musicInfo[key].SongId = songId
		musicInfo[key].SongName = song.Name
		musicInfo[key].SongUrl = songUri
	}

	return &musicInfo, nil

}

func getQQMusicUri(songMid string) (songUri string, err error) {

	rqsUrl, err := url.Parse("https://api.zsfmyz.top/music/song")
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("songmid", songMid)
	params.Set("guid", config.Config.MusicConfig.QQMusicConfig.Guid)

	rqsUrl.RawQuery = params.Encode()

	resp, err := http.Get(rqsUrl.String())
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var musicUrl QQMusicUrl
	err = json.Unmarshal(body, &musicUrl)
	if err != nil {
		return "", err
	}

	return musicUrl.Data.MusicUrl, nil
}

func SearchQQCloudMusic(keywords string, page string) (*[]module.MusicInfo, error) {

	searchUrl, err := url.Parse("https://api.zsfmyz.top/music/list")
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Set("w", keywords)
	params.Set("p", page)
	searchUrl.RawQuery = params.Encode()

	resp, err := http.Get(searchUrl.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var qqMusicSearchResult QQMusicSearchResult
	err = json.Unmarshal(body, &qqMusicSearchResult)
	if err != nil {
		return nil, err
	}

	musicInfo := make([]module.MusicInfo, len(qqMusicSearchResult.Data.List))
	for key, value := range qqMusicSearchResult.Data.List {
		songUrl, err := getQQMusicUri(value.SongMid)
		if err != nil {
			return nil, err
		}

		musicInfo[key].SongUrl = songUrl
		musicInfo[key].SongName = value.SongName
		musicInfo[key].SongId = value.SongMid
		musicInfo[key].ArtistName = value.Singer.Name
		musicInfo[key].ArtistId = value.Singer.Mid
	}

	return &musicInfo, nil
}
