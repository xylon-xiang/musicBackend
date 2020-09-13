package module

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type MusicInfo struct {
	SongUrl    string `bson:"songUrl"`
	SongName   string `bson:"songName"`
	SongId     string `bson:"songId"`
	ArtistName string `bson:"artistName"`
	ArtistId   string `bson:"artistId"`
}

func InsertMusicInfo(musicInfo ...MusicInfo) error {

	var (
		musicInfos []bson.M
		temp       bson.M
	)

	for _, info := range musicInfo {

		byteStream, err := bson.Marshal(bson.M{})
		if err != nil {
			return err
		}

		// turn the info struct into bson.M by reflect while []byte is a intermediate
		t := reflect.TypeOf(info)
		v := reflect.ValueOf(info)
		for i := 0; i <= v.NumField(); i++ {

			byteStream, err = bson.MarshalAppend(byteStream,
				bson.M{t.Field(i).Name: v.Field(i)})
			if err != nil {
				return err
			}
		}

		err = bson.Unmarshal(byteStream, &temp)
		if err != nil {
			return err
		}

		musicInfos = append(musicInfos, temp)

	}

	return nil
}
