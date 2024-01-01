package repository

import "tinder/domain/model"

// RepositoryInterface
type RepositoryInterface interface {
	AddUser(*model.User) string
	GetUser(string) *model.User
	MatchUser(*model.User) []string
	RemoveUser(*model.User)
  QuerySinglePeople(int) []*model.User
}
