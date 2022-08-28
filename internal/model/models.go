package model

type User struct {
	Id    int64  `json:"id,omitempty"`
	Code  int64  `json:"code,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Artist struct {
	Id   int64  `json:"id,omitempty"`
	Code int64  `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type Song struct {
	Id       int64  `json:"id,omitempty"`
	Code     int64  `json:"code,omitempty"`
	Name     string `json:"name,omitempty"`
	Artist   Artist `json:"artist,omitempty"`
	ArtistId int64  `json:"artistId,omitempty"`
}
