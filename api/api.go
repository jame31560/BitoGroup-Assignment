package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
	"tinder/domain/model"
	"tinder/repository"

	"github.com/gin-gonic/gin"
)

type api struct {
	repo repository.RepositoryInterface
}

func NewApi(repo repository.RepositoryInterface) ApiInterface {
	return &api{
		repo: repo,
	}
}

// @BasePath /api

// @Summary	Add a person and match
// @Schemes
// @Description	Add a new user to the matching system and find any possible matches for the new user.
// @Param	data body AddSinglePersonAndMatchReq true "Add user"
// @Tags user
// @Accept json
// @Produce json
// @Success	201	{object} AddSinglePersonAndMatchRes
// @Failure default {object} Err
// @Router /add_user [post]
func (a *api) AddSinglePersonAndMatch(c *gin.Context) {
	reqData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	req := &AddSinglePersonAndMatchReq{}

	err = json.Unmarshal(reqData, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":   "error",
			"data":      "Input fromat Error",
			"timestamp": time.Now().Unix(),
		})
		return
	}

	user := &model.User{
		Name:    req.Name,
		Height:  req.Height,
		Gender:  model.GenderEnum(req.Gender),
		DateNum: req.DateNum,
	}

	userID := a.repo.AddUser(user)

	matchUserList := a.repo.MatchUser(user)

	c.JSON(http.StatusCreated, AddSinglePersonAndMatchRes{
		UserID:        userID,
		MatchUserList: matchUserList,
	})
}

// @Summary	Remove a user.
// @Schemes
// @Description	Remove a user from the matching system so that the user cannot be matched anymore.
// @Param	data body RemoveSinglePersonReq true "Remove user"
// @Tags user
// @Accept json
// @Produce json
// @Success	200	{object} RemoveSinglePersonRes
// @Failure default {object} Err
// @Router /remove_user [delete]
func (a *api) RemoveSinglePerson(c *gin.Context) {

}

// @Summary	Query Single People.
// @Schemes
// @Param	num query integer true "number to query" Quantity
// @Description	Find the most N possible matched single people, where N is a request parameter.
// @Tags user
// @Accept json
// @Produce json
// @Success	200	{object} QuerySinglePeopleRes
// @Failure default {object} Err
// @Router /query_single_user [get]
func (a *api) QuerySinglePeople(c *gin.Context) {
  numString := c.Request.URL.Query().Get("num")
  num, err := strconv.Atoi(numString)
  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":   "error",
			"data":      "Input fromat Error",
			"timestamp": time.Now().Unix(),
		})
		return
	}

  result:= QuerySinglePeopleRes{
    UserList: make([]*UserRes, 0),
  }

  userList := a.repo.QuerySinglePeople(num)

  for _, user := range userList {
    result.UserList = append(result.UserList, &UserRes{
      ID: user.ID,
      Name: user.Name,
      Height: user.Height,
      Gender: int8(user.Gender),
      DateNum: user.DateNum,
    })
  }

  c.JSON(http.StatusOK, result)
}
