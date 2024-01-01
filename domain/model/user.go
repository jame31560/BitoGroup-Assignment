package model

type User struct {
	ID            string
	Name          string
	Height        float32
	Gender        GenderEnum
	DateNum       int64
	MatchUserList []string
}

type GenderEnum int8

const (
	None GenderEnum = iota
	Men
	Woman
)
