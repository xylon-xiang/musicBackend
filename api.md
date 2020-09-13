# API Docs

There is just a function that this backend service provides.
It is a music API and return a module containing all the information about the music.

To improve the request speed, the project will memory all the request into a database to build a cache.

`GET \music`

the database module: 

MusicInfo {

    songUrl:    string
    songName:   string
    songId:     int
    artistName: string
    artistId:   int
    
    
}

