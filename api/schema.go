package api

import "time"

type AddSinglePersonAndMatchReq struct {
	Name    string  `json:"name"`
	Height  float32 `json:"height"`
	Gender  int8    `json:"gender"`
	DateNum int64   `json:"date_num"`
}

type AddSinglePersonAndMatchRes struct {
	UserID        string   `json:"user_id"`
	MatchUserList []string `json:"match_user_list"`
}

type UserRes struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Height  float32 `json:"height"`
	Gender  int8    `json:"gender"`
	DateNum int64   `json:"date_num"`
}

type QuerySinglePeopleRes struct {
	UserList []*UserRes `json:"user_list"`
}

type RemoveSinglePersonReq struct {
	UserID string `json:"user_id"`
}

type RemoveSinglePersonRes struct {
	UserID string `json:"user_id"`
}

type Err struct {
	Message   string    `json:"message"`
	Data      string    `json:"data"`
	Timestamp time.Time `json:"time"`
}
