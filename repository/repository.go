package repository

import (
	"tinder/domain/model"

	"github.com/google/uuid"
)

type repository struct {
	userList []*model.User
}

func (repo *repository) AddUser(user *model.User) string {
	userID := uuid.New().String()
	user.ID = userID
	user.MatchUserList = make([]string, 0)
	repo.userList = append(repo.userList, user)
	return userID
}

func (repo *repository) GetUser(userID string) *model.User {
	var user *model.User
	for _, u := range repo.userList {
		if u.ID == userID {
			user = u
			break
		}
	}
	return user

}

func (repo *repository) MatchUser(user *model.User) []string {
	for _, u := range repo.userList {
		if user.DateNum <= 0 {
			break
		}
		switch false {
		case user.ID != u.ID:
		case u.DateNum > 0:
		case user.Gender != u.Gender:
		case (user.Gender == model.Men && user.Height > u.Height) || (user.Gender == model.Woman && user.Height < u.Height):
		default:
			user.DateNum--
			user.MatchUserList = append(user.MatchUserList, u.ID)
			u.DateNum--
			u.MatchUserList = append(u.MatchUserList, user.ID)
		}
	}
	return user.MatchUserList
}

func (repo *repository) QuerySinglePeople(num int) []*model.User {
	result := make([]*model.User, 0)
	for _, user := range repo.userList {
		if num <= 0 {
			break
		}
		if user.DateNum > 0 {
			result = append(result, user)
			num--
		}
	}
	return result
}

func (repo *repository) RemoveUser(userToRemove *model.User) {
	for _, matchUserID := range userToRemove.MatchUserList {
		matchUser := repo.GetUser(matchUserID)
		updatedMatchUserList := make([]string, 0)
		for _, id := range matchUser.MatchUserList {
			if id == userToRemove.ID {
				updatedMatchUserList = append(updatedMatchUserList, id)
			}
		}
		matchUser.MatchUserList = updatedMatchUserList
		matchUser.DateNum++
	}

	updatedList := make([]*model.User, 0)
	for _, user := range repo.userList {
		if user.ID == userToRemove.ID {
			continue
		}

		updatedList = append(updatedList, user)
	}
	repo.userList = updatedList
}

// NewRepository
func NewRepository() RepositoryInterface {
	userList := make([]*model.User, 0)
	return &repository{
		userList: userList,
	}
}
